package services

import (
	"ecommerce-backend-golang/internal/models"
	"ecommerce-backend-golang/internal/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: repo,
	}
}

func (s *ProductService) Create(product *models.Product) error {
	return s.productRepo.Create(product)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.productRepo.FindAll()
}

func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}
