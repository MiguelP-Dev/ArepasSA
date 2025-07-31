package handlers

import (
	"net/http"
	"strconv"

	"ArepasSA/internal/models"
	"ArepasSA/internal/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	activeOnly := c.DefaultQuery("active", "true") == "true"
	products, err := h.service.GetAllProducts(activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := h.service.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ProductHandler) SoftDeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := h.service.SoftDeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deactivated successfully"})
}
