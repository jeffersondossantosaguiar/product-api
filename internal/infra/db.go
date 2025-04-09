package infra

import (
	"log"

	"github.com/jeffersondossantosaguiar/product-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=db user=test password=test dbname=product port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&domain.Product{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
