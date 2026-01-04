package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-backend-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{orderService: service}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	order, err := c.orderService.CreateFromCart(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}

func (c *OrderController) GetMyOrders(ctx *gin.Context) {
	userID := ctx.MustGet("user_id").(uint)

	orders, err := c.orderService.GetUserOrders(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch orders"})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("id"))

	order, err := c.orderService.GetOrderByID(uint(orderID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) UpdateOrderStatus(ctx *gin.Context) {
	orderID, _ := strconv.Atoi(ctx.Param("id"))

	var req struct {
		Status string `json:"status" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.orderService.UpdateOrderStatus(uint(orderID), req.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "order status updated"})
}