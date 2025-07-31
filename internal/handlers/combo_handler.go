package handlers

import (
	"net/http"
	"strconv"

	"ArepasSA/internal/models"
	"ArepasSA/internal/services"

	"github.com/gin-gonic/gin"
)

type ComboHandler struct {
	service *services.ComboService
}

func NewComboHandler(service *services.ComboService) *ComboHandler {
	return &ComboHandler{service: service}
}

func (h *ComboHandler) CreateCombo(c *gin.Context) {
	var combo models.Combo
	if err := c.ShouldBindJSON(&combo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateCombo(&combo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, combo)
}

func (h *ComboHandler) GetCombo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid combo ID"})
		return
	}

	combo, err := h.service.GetCombo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Combo not found"})
		return
	}

	c.JSON(http.StatusOK, combo)
}

func (h *ComboHandler) GetAllCombos(c *gin.Context) {
	activeOnly := c.DefaultQuery("active", "true") == "true"
	combos, err := h.service.GetAllCombos(activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, combos)
}

func (h *ComboHandler) UpdateCombo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid combo ID"})
		return
	}

	var combo models.Combo
	if err := c.ShouldBindJSON(&combo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	combo.ID = uint(id)
	if err := h.service.UpdateCombo(&combo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, combo)
}

func (h *ComboHandler) DeleteCombo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid combo ID"})
		return
	}

	if err := h.service.DeleteCombo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *ComboHandler) SellCombo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid combo ID"})
		return
	}

	var request struct {
		Quantity float64 `json:"quantity" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SellCombo(uint(id), request.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Combo sold successfully"})
}

func (h *ComboHandler) SellPartialCombo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid combo ID"})
		return
	}

	if err := h.service.SellPartialCombo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Partial combo sold successfully"})
}
