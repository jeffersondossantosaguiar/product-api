package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jefferson/product-api/config"
	"github.com/jefferson/product-api/models"
	"github.com/jefferson/product-api/routes"
)

func main() {
	app := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	routes.SetupRoutes(app)

	app.Run(":8080")
}
