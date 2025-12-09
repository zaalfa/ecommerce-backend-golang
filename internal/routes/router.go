package routes

import (
	"ecommerce-backend-golang/internal/repositories"
	"ecommerce-backend-golang/internal/services"
	"ecommerce-backend-golang/internal/controllers"
	"ecommerce-backend-golang/internal/middleware"
	"ecommerce-backend-golang/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := config.DB 
	// repositories
	productRepo := repositories.NewProductRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// services
	productService := services.NewProductService(productRepo)
	authService := services.NewAuthService(userRepo)

	// controllers
	productController := controllers.NewProductController(productService)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController()

	// routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	user := r.Group("/users")
	user.Use(middleware.AuthRequired())
	{
		user.GET("/me", userController.Me)
	}

	r.GET("/products", productController.GetAll)
	r.GET("/products/:id", productController.GetByID)

	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		admin.POST("/products", productController.Create)
	}

	return r
}


