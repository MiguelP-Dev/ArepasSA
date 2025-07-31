package models

import (
    "time"
    "gorm.io/gorm"
)

// Alert representa una alerta de stock mínimo
type Alert struct {
    gorm.Model
    ProductID   uint      `gorm:"not null" validate:"required"`
    ProductName string    `gorm:"not null"`
    Current     float64   `gorm:"not null"`
    Minimum     float64   `gorm:"not null"`
    TriggeredAt time.Time `gorm:"not null"`
    ResolvedAt  *time.Time
    IsResolved  bool `gorm:"default:false"`
}