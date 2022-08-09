package main

import (
	"log"

	"github.com/alexkopcak/gophkeeper/command-producer-service/internal/app"
)

func main() {
	if err := app.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
