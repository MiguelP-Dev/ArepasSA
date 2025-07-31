package models

import "gorm.io/gorm"

type Combo struct {
	gorm.Model
	Name        string `gorm:"not null" validate:"required,min=3"`
	Description string
	Price       float64 `gorm:"not null" validate:"required,gt=0"`
	IsActive    bool    `gorm:"default:true"`
	Items       []ComboItem
}

type ComboItem struct {
	gorm.Model
	ComboID   uint    `gorm:"not null" validate:"required"`
	ProductID uint    `gorm:"not null" validate:"required"`
	Quantity  float64 `gorm:"not null" validate:"gt=0"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
