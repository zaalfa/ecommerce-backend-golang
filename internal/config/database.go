package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ecommerce-backend-golang/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=rahasia dbname=ecommerce_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
	
	//auto migrate models
	DB.AutoMigrate(&models.User{}, &models.Product{})
	fmt.Println("Database migrated (User table created)")
}