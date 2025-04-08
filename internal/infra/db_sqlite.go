package infra

import (
	"log"

	"github.com/jeffersondossantosaguiar/product-api/internal/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
