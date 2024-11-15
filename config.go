package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rismarahma13/bookstore/models"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Migrasi skema
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("failed to migrate database schema: ", err)
	}
}
