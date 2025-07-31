package models

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string `gorm:"not null" validate:"required,min=3"`
	Email    string `validate:"omitempty,email"`
	Phone    string `validate:"omitempty,numeric"`
	Address  string
	Comments []ClientComment `gorm:"foreignKey:ClientID"`
	IsActive bool            `gorm:"default:true"`
}

type ClientComment struct {
	gorm.Model
	ClientID     uint      `gorm:"not null" validate:"required"`
	Author       string    `gorm:"not null"` // Puede ser "system" o el usuario que hizo el comentario
	Content      string    `gorm:"not null" validate:"required,min=5"`
	CommentAt    time.Time `gorm:"not null"`
	IsPreference bool      `gorm:"default:false"` // Para marcar comentarios sobre preferencias
}
