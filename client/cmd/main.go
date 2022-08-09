package main

import (
	"log"

	"github.com/alexkopcak/gophkeeper/client/internal/app"
)

func main() {
	a := app.NewApp()
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
