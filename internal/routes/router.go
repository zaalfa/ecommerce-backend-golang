package routes

import (
	"ecommerce-backend-golang/internal/controllers"
	"ecommerce-backend-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	authController := controllers.AuthController{}
	userController := controllers.UserController{}
	productController := controllers.ProductController{}

	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/users/me", userController.Me)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Public
	r.GET("/products", productController.GetAll)
	r.GET("/products/:id", productController.GetByID)

	// Admin
	admin := r.Group("/admin")
	admin.Use(middleware.AuthRequired(), middleware.AdminOnly())
	{
		admin.POST("/products", productController.Create)
	}
	return r
}