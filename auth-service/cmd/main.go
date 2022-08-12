package main

import (
	"log"

	"github.com/alexkopcak/gophkeeper/auth-service/internal/app"
)

func main() {
	if err := app.NewApp().Run(); err != nil {
		log.Fatal(err)
	}
}
