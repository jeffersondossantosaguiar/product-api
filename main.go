package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersondossantosaguiar/product-api/config"
	"github.com/jeffersondossantosaguiar/product-api/models"
	"github.com/jeffersondossantosaguiar/product-api/routes"
)

func main() {
	app := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	routes.SetupRoutes(app)

	app.Run(":8080")
}
