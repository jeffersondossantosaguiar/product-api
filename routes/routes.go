package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersondossantosaguiar/product-api/handlers"
)

func SetupRoutes(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	{
		v1.GET("/products", handlers.GetProducts)
		v1.GET("/products/:id", handlers.GetProduct)
		v1.POST("/products", handlers.CreateProduct)
		v1.PUT("/products/:id", handlers.UpdateProduct)
		v1.DELETE("/products/:id", handlers.DeleteProduct)
	}
}
