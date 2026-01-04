package routes

import (
	"ecommerce-backend-golang/internal/config"
	"ecommerce-backend-golang/internal/controllers"
	"ecommerce-backend-golang/internal/middleware"
	"ecommerce-backend-golang/internal/repositories"
	"ecommerce-backend-golang/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db := config.DB

	// Repositories
	productRepo := repositories.NewProductRepository(db)
	userRepo := repositories.NewUserRepository(db)
	cartRepo := repositories.NewCartRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Services
	productService := services.NewProductService(productRepo)
	authService := services.NewAuthService(userRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	orderService := services.NewOrderService(orderRepo, cartRepo, productRepo, db)

	// Controllers
	productController := controllers.NewProductController(productService)
	authController := controllers.NewAuthController(authService)
	userController := controllers.NewUserController()
	cartController := controllers.NewCartController(cartService)
	orderController := controllers.NewOrderController(orderService)

	// Public routes
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	r.GET("/products", productController.GetAll)
	r.GET("/products/:id", productController.GetByID)

	// Protected routes (user)
	user := r.Group("/users")
	user.Use(middleware.AuthRequired())
	{
		user.GET("/me", userController.Me)
	}

	// Cart routes (user)
	cart := r.Group("/cart")
	cart.Use(middleware.AuthRequired())
	{
		cart.GET("", cartController.GetCart)
		cart.POST("/items", cartController.AddItem)
		cart.PUT("/items/:id", cartController.UpdateItem)
		cart.DELETE("/items/:id", cartController.RemoveItem)
		cart.DELETE("", cartController.ClearCart)
	}

	// Order routes (user)
	orders := r.Group("/orders")
	orders.Use(middleware.AuthRequired())
	{
		orders.POST("", orderController.CreateOrder)
		orders.GET("", orderController.GetMyOrders)
		orders.GET("/:id", orderController.GetOrderByID)
	}

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		admin.POST("/products", productController.Create)
		admin.PUT("/orders/:id/status", orderController.UpdateOrderStatus)
	}

	return r
}