package main

import (
	"log"
	"os"

	"github.com/arthurh0812/cloudsql/db"
)

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
	log.Printf("successfully created accounts table in database %s", database.String())
}
