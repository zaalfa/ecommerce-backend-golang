package services

import (
	"errors"

	"ecommerce-backend-golang/internal/models"
	"ecommerce-backend-golang/internal/repositories"
	"gorm.io/gorm"
)

type CartService struct {
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewCartService(cartRepo *repositories.CartRepository, productRepo *repositories.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *CartService) GetOrCreateCart(userID uint) (*models.Cart, error) {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newCart := &models.Cart{UserID: userID}
			if err := s.cartRepo.Create(newCart); err != nil {
				return nil, err
			}
			return newCart, nil
		}
		return nil, err
	}
	return cart, nil
}

func (s *CartService) AddItem(userID, productID uint, quantity int) error {
	product, err := s.productRepo.FindByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	if product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	cart, err := s.GetOrCreateCart(userID)
	if err != nil {
		return err
	}

	existingItem, err := s.cartRepo.FindCartItem(cart.ID, productID)
	if err == nil {
		existingItem.Quantity += quantity
		if existingItem.Quantity > product.Stock {
			return errors.New("insufficient stock")
		}
		return s.cartRepo.UpdateItem(existingItem)
	}

	item := &models.CartItem{
		CartID:    cart.ID,
		ProductID: productID,
		Quantity:  quantity,
	}
	return s.cartRepo.CreateItem(item)
}

func (s *CartService) UpdateItemQuantity(userID, itemID uint, quantity int) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("cart not found")
	}

	var item *models.CartItem
	for i := range cart.Items {
		if cart.Items[i].ID == itemID {
			item = &cart.Items[i]
			break
		}
	}

	if item == nil {
		return errors.New("item not found in cart")
	}

	product, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	if quantity > product.Stock {
		return errors.New("insufficient stock")
	}

	item.Quantity = quantity
	return s.cartRepo.UpdateItem(item)
}

func (s *CartService) RemoveItem(userID, itemID uint) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return errors.New("cart not found")
	}

	found := false
	for _, item := range cart.Items {
		if item.ID == itemID {
			found = true
			break
		}
	}

	if !found {
		return errors.New("item not found in cart")
	}

	return s.cartRepo.DeleteItem(itemID)
}

func (s *CartService) GetCart(userID uint) (*models.Cart, error) {
	return s.cartRepo.FindByUserID(userID)
}

func (s *CartService) ClearCart(userID uint) error {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return err
	}
	return s.cartRepo.ClearCart(cart.ID)
}