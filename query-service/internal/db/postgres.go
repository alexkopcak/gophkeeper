package db

import (
	"log"

	"github.com/alexkopcak/gophkeeper/query-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

var _ Storage = (*Postgres)(nil)

func NewPostgres(url string) Storage {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.Record{}); err != nil {
		log.Fatal(err)
	}

	return &Postgres{db}
}

func (p *Postgres) GetRecord(value *models.Record) (*[]models.Record, error) {
	res := p.DB.Where(&value)

	if res.Error != nil {
		return nil, res.Error
	}

	count := res.RowsAffected
	rows, err := res.Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]models.Record, 0, count)

	var record *models.Record
	for rows.Next() {
		err = p.DB.ScanRows(rows, record)

		if err != nil {
			return nil, err
		}

		items = append(items, models.Record{
			Id:          record.Id,
			UserId:      record.UserId,
			MessageType: record.MessageType,
			Data:        record.Data,
			Meta:        record.Meta,
		})
	}

	return &items, nil
}
