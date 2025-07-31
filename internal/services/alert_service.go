package services

import (
	"ArepasSA/internal/models"
	"ArepasSA/internal/repositories"
	"log"
	"time"
)

type AlertService struct {
	alertRepo   *repositories.AlertRepository
	productRepo *repositories.ProductRepository
}

func NewAlertService(
	alertRepo *repositories.AlertRepository,
	productRepo *repositories.ProductRepository,
) *AlertService {
	return &AlertService{
		alertRepo:   alertRepo,
		productRepo: productRepo,
	}
}

func (s *AlertService) CheckStockAlerts() ([]models.Alert, error) {
	products, err := s.productRepo.GetLowStockProducts(0)
	if err != nil {
		return nil, err
	}

	var newAlerts []models.Alert
	for _, product := range products {
		// Verificar si ya existe una alerta activa para este producto
		activeAlerts, err := s.alertRepo.GetActiveAlerts()
		if err != nil {
			log.Printf("Error checking active alerts: %v", err)
			continue
		}

		alertExists := false
		for _, alert := range activeAlerts {
			if alert.ProductID == product.ID {
				alertExists = true
				break
			}
		}

		if !alertExists {
			newAlert := models.Alert{
				ProductID:   product.ID,
				ProductName: product.Name,
				Current:     product.Stock,
				Minimum:     product.MinStock,
				TriggeredAt: time.Now(),
			}
			if err := s.alertRepo.Create(&newAlert); err != nil {
				log.Printf("Failed to create alert: %v", err)
			} else {
				newAlerts = append(newAlerts, newAlert)
				log.Printf("New stock alert created: Product %s (ID: %d) - Stock: %.2f < Min: %.2f",
					product.Name, product.ID, product.Stock, product.MinStock)
			}
		}
	}

	return newAlerts, nil
}

func (s *AlertService) ResolveAlert(id uint) error {
	return s.alertRepo.Resolve(id)
}

func (s *AlertService) GetActiveAlerts() ([]models.Alert, error) {
	return s.alertRepo.GetActiveAlerts()
}

func (s *AlertService) GetResolvedAlerts() ([]models.Alert, error) {
	return s.alertRepo.GetResolvedAlerts()
}
