package main

import (
	"log"
	"os"

	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"
)

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" { // default port
		port = "8080"
	}
}

func main() {
	app := framework.New()

	err := app.Build()
	if err != nil {
		log.Fatal(err)
	}

	app.Register("GET", "/", func(ctx *context.Context) {
		ctx.Write([]byte("Hi there! My name is Max"))
	})

	err = app.ListenAndServe(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
