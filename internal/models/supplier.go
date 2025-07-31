package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
	Name     string `gorm:"not null" validate:"required"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,numeric"`
	Address  string
	IsActive bool `gorm:"default:true"`
}
