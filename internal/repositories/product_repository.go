package repositories

import (
	"ecommerce-backend-golang/internal/config"
	"ecommerce-backend-golang/internal/models"
)

type ProductRepository struct{}

func (r *ProductRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	return &product, err
}
