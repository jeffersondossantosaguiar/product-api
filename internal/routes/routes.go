package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffersondossantosaguiar/product-api/internal/handler"
)

func RegisterRoutes(app *gin.Engine, productHandler *handler.ProductHandler) {
	products := app.Group("/v1/products")
	{
		products.GET("/:id", productHandler.GetByID)
		products.GET("", productHandler.GetAll)
		products.POST("", productHandler.Create)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}
}
