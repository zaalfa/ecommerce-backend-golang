package models

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `gorm:"not null" json:"user_id"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TotalPrice int         `gorm:"not null" json:"total_price"`
	Status     string      `gorm:"type:varchar(20);default:'pending'" json:"status"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

type OrderItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `gorm:"not null" json:"order_id"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     int       `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}