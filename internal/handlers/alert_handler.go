package handlers

import (
	"ArepasSA/internal/services"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	alertService *services.AlertService
}

func NewAlertHandler(alertService *services.AlertService) *AlertHandler {
	return &AlertHandler{alertService: alertService}
}

// Inicia el monitoreo periódico de alertas
func (h *AlertHandler) StartAlertMonitor() {
	ticker := time.NewTicker(5 * time.Minute) // Verificar cada 5 minutos
	go func() {
		for range ticker.C {
			if _, err := h.alertService.CheckStockAlerts(); err != nil {
				log.Printf("Error checking stock alerts: %v", err)
			}
		}
	}()
}

// Endpoint para obtener alertas activas
func (h *AlertHandler) GetActiveAlerts(c *gin.Context) {
	alerts, err := h.alertService.GetActiveAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"alerts": alerts})
}

// Endpoint para resolver una alerta
func (h *AlertHandler) ResolveAlert(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Alert ID is required"})
		return
	}

	// Convertir ID a uint
	var alertID uint
	if _, err := fmt.Sscanf(id, "%d", &alertID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	if err := h.alertService.ResolveAlert(alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert resolved successfully"})
}

// Endpoint para obtener alertas resueltas
func (h *AlertHandler) GetResolvedAlerts(c *gin.Context) {
	alerts, err := h.alertService.GetResolvedAlerts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"alerts": alerts})
}
