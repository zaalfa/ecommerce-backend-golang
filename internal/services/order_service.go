package services

import (
	"errors"

	"ecommerce-backend-golang/internal/models"
	"ecommerce-backend-golang/internal/repositories"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo   *repositories.OrderRepository
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
	db          *gorm.DB
}

func NewOrderService(
	orderRepo *repositories.OrderRepository,
	cartRepo *repositories.CartRepository,
	productRepo *repositories.ProductRepository,
	db *gorm.DB,
) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		cartRepo:    cartRepo,
		productRepo: productRepo,
		db:          db,
	}
}

func (s *OrderService) CreateFromCart(userID uint) (*models.Order, error) {
	cart, err := s.cartRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var totalPrice int
	var orderItems []models.OrderItem

	for _, item := range cart.Items {
		product, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("product not found: " + product.Name)
		}

		if product.Stock < item.Quantity {
			tx.Rollback()
			return nil, errors.New("insufficient stock for: " + product.Name)
		}

		product.Stock -= item.Quantity
		if err := tx.Save(product).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		totalPrice += product.Price * item.Quantity
	}

	order := &models.Order{
		UserID:     userID,
		TotalPrice: totalPrice,
		Status:     "pending",
		Items:      orderItems,
	}

	if err := tx.Create(order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.cartRepo.ClearCart(cart.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	result, _ := s.orderRepo.FindByID(order.ID)
	return result, nil
}

func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.FindByUserID(userID)
}

func (s *OrderService) GetOrderByID(orderID uint) (*models.Order, error) {
	return s.orderRepo.FindByID(orderID)
}

func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	validStatuses := map[string]bool{
		"pending": true, "paid": true, "shipped": true, "delivered": true, "cancelled": true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return s.orderRepo.UpdateStatus(orderID, status)
}