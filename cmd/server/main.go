package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersondossantosaguiar/product-api/internal/domain"
	"github.com/jeffersondossantosaguiar/product-api/internal/handler"
	"github.com/jeffersondossantosaguiar/product-api/internal/infra"
	"github.com/jeffersondossantosaguiar/product-api/internal/usecase"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	db.AutoMigrate(&domain.Product{})

	productRepo := infra.NewProductGormRepository(db)
	productUseCase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUseCase)

	app := gin.Default()
	productHandler.RegisterRoutes(app)

	app.Run(":8080")

}
