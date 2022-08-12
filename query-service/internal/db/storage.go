package db

import "github.com/alexkopcak/gophkeeper/query-service/internal/models"

type Storage interface {
	GetRecord(value *models.Record) (*[]models.Record, error)
}
