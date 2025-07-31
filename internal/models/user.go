package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null" validate:"required"`
	Email    string `gorm:"unique;not null" validate:"required,email"`
	Password string `gorm:"not null" validate:"required,min=6"`
	Phone    string `validate:"omitempty,numeric"`
	IsActive bool   `gorm:"default:true"`
}
