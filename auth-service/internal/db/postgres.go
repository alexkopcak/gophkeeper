package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/alexkopcak/gophkeeper/auth-service/internal/models"
	"github.com/alexkopcak/gophkeeper/auth-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

var _ (Storage) = (*Postgres)(nil)

var (
	ErrorUserNotFound     = errors.New("user not found")
	ErrorUserAlreadyExist = errors.New("user already exist")
)

func NewPostgres(url string) Storage {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	return &Postgres{db}
}

func (p *Postgres) GetUser(value *models.User) (*models.User, error) {
	user := models.User{
		Name: value.Name,
	}

	tx := p.DB.Where(&models.User{Name: value.Name})

	fmt.Println(tx.Error)
	fmt.Println(tx.RowsAffected)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return &user, ErrorUserNotFound
	}

	txf := tx.First(&user)
	if txf.Error != nil {
		return &user, ErrorUserNotFound
	}

	return &user, nil
}

func (p *Postgres) AddUser(value *models.User) (*models.User, error) {
	user, err := p.GetUser(value)

	if err == nil {
		return user, ErrorUserAlreadyExist
	}

	user.Password, err = utils.HashPassword(value.Password)

	if err != nil {
		return nil, err
	}

	tx := p.DB.Create(&user)

	return user, tx.Error
}
