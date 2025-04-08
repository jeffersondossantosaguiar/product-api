package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersondossantosaguiar/product-api/internal/handler"
	"github.com/jeffersondossantosaguiar/product-api/internal/infra"
	"github.com/jeffersondossantosaguiar/product-api/internal/routes"
	"github.com/jeffersondossantosaguiar/product-api/internal/usecase"
)

func main() {
	db := infra.InitDB()

	productRepo := infra.NewProductGormRepository(db)
	productUseCase := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUseCase)

	app := gin.Default()
	routes.RegisterRoutes(app, productHandler)

	app.Run(":8080")

}
