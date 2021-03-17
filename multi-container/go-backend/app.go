package main

import (
	"github.com/arthurh0812/framework"
	"github.com/arthurh0812/framework/context"
	"github.com/arthurh0812/framework/lib"
	"log"
)

func main() {
	app := framework.New()

	err := app.Build()
	if err != nil {
		log.Fatalf("failed to build the routes: %v", err)
	}

	app.Router.Use(func(ctx *context.Context) {
		ctx.Header().Set(lib.HeaderKeyAccessControlAllowOrigin, "*")
		ctx.Header().Set(lib.HeaderKeyAccessControlAllowMethods, "GET, POST, DELETE, OPTIONS")
		ctx.Header().Set(lib.HeaderKeyAccessControlAllowHeaders, "Content-Type")
	})

	err = app.ListenAndServe(":80")
	if err != nil {
		log.Fatalf("failed to listen to and serve requests: %v", err)
	}
}