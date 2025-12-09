package repositories

import (
	"ecommerce-backend-golang/internal/config"
	"ecommerce-backend-golang/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}