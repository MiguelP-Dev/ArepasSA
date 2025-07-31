package handlers

import (
	"ArepasSA/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ... (métodos existentes)

func (h *ClientHandler) AddComment(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	var comment models.ClientComment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddComment(uint(clientID), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, comment)
}

func (h *ClientHandler) GetPreferences(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de cliente inválido"})
		return
	}

	preferences, err := h.service.GetPreferences(uint(clientID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"preferences": preferences})
}
