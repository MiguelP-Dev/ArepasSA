package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	Name      string         `gorm:"not null" validate:"required,min=3"`
	Email     string         `validate:"omitempty,email"`
	Phone     string         `validate:"omitempty,numeric"`
	Address   string
	Comments  []ClientComment `gorm:"foreignKey:ClientID"`
	IsActive  bool           `gorm:"default:true"`
}

type ClientComment struct {
	gorm.Model
	ClientID    uint      `gorm:"not null" validate:"required"`
	Author      string    `gorm:"not null" validate:"required"` // "system" o username
	Content     string    `gorm:"not null" validate:"required,min=5,max=500"`
	CommentType string    `gorm:"not null" validate:"oneof=preference general note"` // tipos de comentario
	CommentDate time.Time `gorm:"not null"`
}
