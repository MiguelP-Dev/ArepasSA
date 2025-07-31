package repositories

import (
    "ArepasSA/internal/models"
    "time"
    
    "gorm.io/gorm"
)

type AlertRepository struct {
    *BaseRepository
}

func NewAlertRepository(db *gorm.DB) *AlertRepository {
    return &AlertRepository{NewBaseRepository(db)}
}

func (r *AlertRepository) Create(alert *models.Alert) error {
    return r.db.Create(alert).Error
}

func (r *AlertRepository) Resolve(id uint) error {
    now := time.Now()
    return r.db.Model(&models.Alert{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "is_resolved": true,
            "resolved_at": now,
        }).Error
}

func (r *AlertRepository) GetActiveAlerts() ([]models.Alert, error) {
    var alerts []models.Alert
    err := r.db.Where("is_resolved = ?", false).Find(&alerts).Error
    return alerts, err
}

func (r *AlertRepository) GetResolvedAlerts() ([]models.Alert, error) {
    var alerts []models.Alert
    err := r.db.Where("is_resolved = ?", true).Find(&alerts).Error
    return alerts, err
}