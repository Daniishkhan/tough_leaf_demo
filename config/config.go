package config

import (
	"log"
	"toughleaf/internal/models"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	var err error
	log.Println("Connecting to database")
	Db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = Db.AutoMigrate(&models.User{}, &models.Bid{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}
	log.Println("Connected to database and migrated models")
}
