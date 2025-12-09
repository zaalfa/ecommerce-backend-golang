package services

import (
	"errors"

	"ecommerce-backend-golang/internal/models"
	"ecommerce-backend-golang/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}
func (s *AuthService) Register(name, email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     "user",
	}

	return s.userRepo.Create(&user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("Invalid email or password")
	}

	return user, nil
}