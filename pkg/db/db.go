package db

import (
	"ebooks/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// dbUrl := "postgres://postgres:postgres@localhost:5432/books"
	dsn := "host=localhost user=postgres password=example dbname=books"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Author{})

	return db
}
