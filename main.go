package main

import (
	"graphql/blog-go-graphql/infrastructure/server"
	"log"
	"os"
)

func main() {
	app := server.InitServer()
	port := os.Getenv("PORT")

	err := app.Listen(":" + port)
	if err != nil {
		log.Panic(err)
	}
}
