package main

import (
	"graphql/blog-go-graphql/infrastructure/server"
	"log"
)

func main() {
	app := server.InitServer()

	err := app.Listen(":8080")
	if err != nil {
		log.Panic(err)
	}
}
