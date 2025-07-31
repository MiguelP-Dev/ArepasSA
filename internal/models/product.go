package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"not null" validate:"required,min=3"`
	Description string
	BuyPrice    float64 `gorm:"not null" validate:"required,gt=0"`
	SellPrice   float64 `gorm:"not null" validate:"required,gt=0"`
	Stock       float64 `gorm:"not null" validate:"gte=0"`
	MinStock    float64 `gorm:"not null" validate:"gte=0"`
	Barcode     string  `gorm:"unique"`
	SupplierID  uint
	Supplier    Supplier
	IsActive    bool `gorm:"default:true"`
	Alerts []Alert `gorm:"foreignKey:ProductID"`
}
