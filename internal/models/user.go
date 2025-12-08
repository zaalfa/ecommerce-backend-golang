package models

import "time"

type User struct {
	ID       uint		`gorm:"primary_key" json:"id"`
	Name     string		`json:"name"`
	Email    string		`gorm:"unique" json:"email"`
	Password string		`json:"-"`
	Role      string    `gorm:"type:varchar(20);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}