package controllers

import (
	"net/http"

	"ecommerce-backend-golang/internal/config"
	"ecommerce-backend-golang/internal/models"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) Me(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
