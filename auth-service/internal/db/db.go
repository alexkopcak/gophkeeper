package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/alexkopcak/gophkeeper/auth-service/internal/models"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) *Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	return &Handler{db}
}
