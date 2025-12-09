package controllers

import (
	"net/http"
	"strconv"

	"ecommerce-backend-golang/internal/models"
	"ecommerce-backend-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func (c *ProductController) Create(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.productService.Create(&product); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func (c *ProductController) GetAll(ctx *gin.Context) {
	products, _ := c.productService.GetAll()
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := c.productService.GetByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
