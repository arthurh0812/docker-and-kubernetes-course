package main

import (
	"log"
	"net/http"

	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"
)

func main() {
	app := framework.New()

	app.Register(http.MethodGet, "/", func(ctx *context.Context) {
		d := context.NewDirector(ctx)

		d.WriteString("Hello there! This will become an API!")
	})

	log.Fatal(app.ListenAndServe(":443"))
}
