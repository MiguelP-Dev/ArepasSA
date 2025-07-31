package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"time"
)

type ReportService struct {
	repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) GetDailySalesReport(date time.Time) (float64, int, error) {
	return s.repo.GetDailySales(date)
}

func (s *ReportService) GetPeakHoursReport() (map[int]int, error) {
	return s.repo.GetSalesByHour()
}

func (s *ReportService) GetTopProducts(limit int) ([]models.ProductSales, error) {
	return s.repo.GetTopProducts(limit)
}

func (s *ReportService) GetLeastSoldProducts(limit int) ([]models.ProductSales, error) {
	return s.repo.GetLeastSoldProducts(limit)
}

func (s *ReportService) GetSalesByPriceRange() ([]models.PriceRangeSales, error) {
	return s.repo.GetSalesByPriceRange()
}
