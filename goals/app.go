package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"

	"github.com/arthurh0812/goals/db"
	"github.com/arthurh0812/goals/graphql"

	"github.com/go-pg/pg"
)

var port string

func init() {
	port = os.Getenv("PORT")
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	info := &db.ConnectionInfo{
		User:     dbUser,
		Password: dbPass,
		Database: dbName,
		Host:     dbHost,
		Port:     dbPort,
	}

	database, err := db.Connect(info)
	if err != nil {
		log.Fatalf("connecting to database server: fail: %v", err)
	}
	log.Printf("successfully connected to database server at %s:%s", dbHost, dbPort)

	err = db.CreateAccountsTable(database)
	if err != nil {
		log.Fatalf("creeating accounts table: fail: %v", err)
	}
	log.Printf("successfully created accounts table in database server %s", database.String())

	err = saveTestAccount(database, 5324)
	if err != nil {
		log.Fatalf("saving account record: %v", err)
	}
	log.Printf("successfully saved account record in database server %s", database.String())

	app := framework.New()

	app.Register("POST", "/goals", func(ctx *context.Context) {
		dir := context.NewDirector(ctx)
		dec := json.NewDecoder(ctx.Body())
		input := &graphql.Input{}
		err := dec.Decode(input)
		if err != nil {
			_ = dir.Error(http.StatusInternalServerError, "failed to parse request body")
		}

	})

	err = app.Build()
	if err != nil {
		log.Fatal(err)
	}

	err = app.ListenAndServe(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}

func saveTestAccount(database *pg.DB, ownerID int64) error {
	newAccount := &db.Account{
		Name: "my-special-fortnite-account",
		Owner: ownerID,
		Token: "random-token",
	}
	return newAccount.Save(database)
}