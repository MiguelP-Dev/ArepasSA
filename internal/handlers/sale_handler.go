package handlers

import (
	"net/http"
	"strconv"
	"time"

	"ArepasSA/internal/models"
	"ArepasSA/internal/services"

	"github.com/gin-gonic/gin"
)

type SaleHandler struct {
	service *services.SaleService
}

func NewSaleHandler(service *services.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

func (h *SaleHandler) CreateSale(c *gin.Context) {
	var sale models.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sale.Date = time.Now() // Establecer la fecha actual

	if err := h.service.CreateSale(&sale); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sale)
}

func (h *SaleHandler) GetSale(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
		return
	}

	sale, err := h.service.GetSale(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}

	c.JSON(http.StatusOK, sale)
}

func (h *SaleHandler) GetAllSales(c *gin.Context) {
	startDate, _ := time.Parse("2006-01-02", c.Query("start_date"))
	endDate, _ := time.Parse("2006-01-02", c.Query("end_date"))

	sales, err := h.service.GetAllSales(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sales)
}

func (h *SaleHandler) AddComment(c *gin.Context) {
	saleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de venta inválido"})
		return
	}

	var comment models.SaleComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Obtener el autor (usuario) del contexto de autenticación
	user, _ := c.Get("user")
	comment.Author = user.(string)

	if err := h.service.AddComment(uint(saleID), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *SaleHandler) GetComments(c *gin.Context) {
	saleID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de venta inválido"})
		return
	}

	includeInternal := c.Query("internal") == "true"
	comments, err := h.service.GetComments(uint(saleID), includeInternal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
