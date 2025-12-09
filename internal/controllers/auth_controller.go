package controllers

import (
	"net/http"

	"ecommerce-backend-golang/internal/services"
	"ecommerce-backend-golang/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.authService.Register(req.Name, req.Email, req.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "register success"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": "failed to generate token",
		})
		return
	}


	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
