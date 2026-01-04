package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-backend-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService *services.CartService
}

func NewCartController(service *services.CartService) *CartController {
	return &CartController{cartService: service}
}

func (c *CartController) GetCart(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	cart, err := c.cartService.GetCart(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "cart not found"})
		return
	}

	ctx.JSON(http.StatusOK, cart)
}

func (c *CartController) AddItem(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.cartService.AddItem(userID, req.ProductID, req.Quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "item added to cart"})
}

func (c *CartController) UpdateItem(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	itemID, _ := strconv.Atoi(ctx.Param("id"))

	var req struct {
		Quantity int `json:"quantity" binding:"required,min=1"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.cartService.UpdateItemQuantity(userID, uint(itemID), req.Quantity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "cart item updated"})
}

func (c *CartController) RemoveItem(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)
	itemID, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.cartService.RemoveItem(userID, uint(itemID)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "item removed from cart"})
}

func (c *CartController) ClearCart(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	if err := c.cartService.ClearCart(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear cart"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "cart cleared"})
}