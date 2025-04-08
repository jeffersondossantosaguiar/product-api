package repositories

import (
	"github.com/jeffersondossantosaguiar/product-api/config"
	"github.com/jeffersondossantosaguiar/product-api/models"
)

func GetAllProducts() []models.Product {
	var products []models.Product
	config.DB.Find(&products)
	return products
}

func GetAllProductByID(id uint) (models.Product, error) {
	var product models.Product
	result := config.DB.First(&product, id)
	return product, result.Error
}

func CreateProduct(product models.Product) {
	config.DB.Create(&product)
}

func UpdateProduct(product models.Product) {
	config.DB.Save(product)
}

func DeleteProduct(id uint) {
	config.DB.Delete(&models.Product{}, id)
}
