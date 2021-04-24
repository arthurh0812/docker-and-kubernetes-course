package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"log"
	"net/http"
)

const port = "80"
const hashEnding = "_hash"

func generateHash(entered string) (hashed string) {
	return entered + hashEnding
}

func isValidHash(entered, hashed string) bool {
	return hashed == entered + hashEnding
}

func main() {
	app := iris.Default()

	app.Use(func(ctx *context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST,GET,OPTION")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		ctx.Next()
	})

	app.Get("/verify-token/{token:string}", func(ctx *context.Context) {
		token := ctx.Params().GetString("token")

		if token == "abc" {
			ctx.StopWithJSON(http.StatusOK, APIResponse{
				Status: http.StatusOK,
				Message: "Valid token.",
				Data: iris.Map{
					"uid": "u1",
				},
			})
			return
		}
		ctx.StopWithJSON(http.StatusUnauthorized, APIResponse{
			Status: http.StatusUnauthorized,
			Message: "Invalid token.",
			Data: nil,
		})
	})

	app.Get("/token/{hashedPassword:string}/{enteredPassword:string}", func(ctx *context.Context) {
		hashedPassword := ctx.Params().GetString("hashedPassword")
		enteredPassword := ctx.Params().GetString("enteredPassword")

		if isValidHash(enteredPassword, hashedPassword) {
			token := "abc"
			ctx.StopWithJSON(http.StatusOK, APIResponse{
				Status: http.StatusOK,
				Message: "Token was created.",
				Data: iris.Map{
					"token": token,
				},
			})
			return
		}
		ctx.StopWithJSON(http.StatusUnauthorized, APIResponse{
			Status: http.StatusUnauthorized,
			Message: "Passwords do not match.",
			Data: nil,
		})
	})

	app.Get("/hashed-password/{password:string}", func(ctx *context.Context) {
		enteredPassword := ctx.Params().GetString("password")
		hashedPassword := generateHash(enteredPassword)
		ctx.StopWithJSON(http.StatusOK, APIResponse{
			Status: http.StatusOK,
			Message: "Password was hashed.",
			Data: iris.Map{
				"hashedPassword": hashedPassword,
			},
		})
	})

	err := app.Run(iris.Addr(":"+port))
	if err != nil {
		log.Fatalf("failed to listen to port %s: %v", port, err)
	}
}
