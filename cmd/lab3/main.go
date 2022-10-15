package main

import (
	"context"
	"lab3/internal/pkg/app"
	"log"
	"os"
)

// @title DOCS
// @description My first docs
// @version 1.0
// @contact.name Dasha
// @license.name license1
// @host 127.0.0.1
// @schemes https http
// @BasePath /

func main() {
	log.Println("Application start")
	ctx := context.Background()
	a, err := app.New(ctx)
	if err != nil {
		log.Println("Application failed")
		os.Exit(2)
	}
	a.StartServer()
	log.Println("Application terminate")
}
