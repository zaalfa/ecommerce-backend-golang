package controllers

import (
	"net/http"

	"ecommerce-backend-golang/internal/repositories"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo repositories.UserRepository
}

func (c *UserController) Me(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	user, err := c.userRepo.FindByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}
