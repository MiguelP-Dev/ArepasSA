package handlers

import (
	"ArepasSA/internal/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetDailySalesReport(c *gin.Context) {
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
		return
	}

	total, count, err := h.service.GetDailySalesReport(date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"date":         date.Format("2006-01-02"),
		"total_amount": total,
		"sale_count":   count,
	})
}

func (h *ReportHandler) GetPeakHoursReport(c *gin.Context) {
	hours, err := h.service.GetPeakHoursReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"peak_hours": hours})
}

func (h *ReportHandler) GetTopProductsReport(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	products, err := h.service.GetTopProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"top_products": products})
}

func (h *ReportHandler) GetLeastSoldProductsReport(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	products, err := h.service.GetLeastSoldProducts(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"least_sold_products": products})
}

func (h *ReportHandler) GetPriceRangeReport(c *gin.Context) {
	ranges, err := h.service.GetSalesByPriceRange()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price_ranges": ranges})
}
