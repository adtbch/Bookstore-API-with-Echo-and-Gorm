// pkg/config/app.go
package config

import (
	"Weekly-Task3/pkg/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)


func InitDB() *gorm.DB {
	// Open SQLite database in pure Go mode
	db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Automigrate models (replace with your models)
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	return db
}
