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

	return r
}