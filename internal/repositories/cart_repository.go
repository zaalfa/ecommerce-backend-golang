package repositories

import (
	"ecommerce-backend-golang/internal/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) FindByUserID(userID uint) (*models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

func (r *CartRepository) Create(cart *models.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) FindCartItem(cartID, productID uint) (*models.CartItem, error) {
	var item models.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", cartID, productID).First(&item).Error
	return &item, err
}

func (r *CartRepository) CreateItem(item *models.CartItem) error {
	return r.db.Create(item).Error
}

func (r *CartRepository) UpdateItem(item *models.CartItem) error {
	return r.db.Save(item).Error
}

func (r *CartRepository) DeleteItem(itemID uint) error {
	return r.db.Delete(&models.CartItem{}, itemID).Error
}

func (r *CartRepository) ClearCart(cartID uint) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}