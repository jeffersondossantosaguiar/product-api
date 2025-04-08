package infra

import (
	"github.com/jeffersondossantosaguiar/product-api/internal/domain"
	"github.com/jeffersondossantosaguiar/product-api/internal/repository"
	"gorm.io/gorm"
)

type ProductGormRepository struct {
	db *gorm.DB
}

func NewProductGormRepository(db *gorm.DB) repository.ProductRepository {
	return &ProductGormRepository{db: db}
}

func (r *ProductGormRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductGormRepository) Delete(id uint) error {
	return r.db.Delete(id).Error
}

func (r *ProductGormRepository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductGormRepository) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductGormRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}
