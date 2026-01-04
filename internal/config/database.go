package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ecommerce-backend-golang/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Baca dari environment variable
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "rahasia")
	dbname := getEnv("DB_NAME", "ecommerce_db")
	port := getEnv("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db
	fmt.Println("Database connected successfully")
	
	// Auto migrate models - TAMBAH cart & order models!
	DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.OrderItem{},
	)
	fmt.Println("Database migrated successfully")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}