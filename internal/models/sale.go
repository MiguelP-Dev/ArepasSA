package models

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	Date        time.Time `gorm:"not null"`
	TotalAmount float64   `gorm:"not null"`
	Items       []SaleItem
	ClientID    *uint         // Opcional
	Client      *Client       `gorm:"foreignKey:ClientID"`
	Comments    []SaleComment `gorm:"foreignKey:SaleID"`
}
type SaleItem struct {
	gorm.Model
	SaleID     uint     `gorm:"not null" validate:"required"`
	ProductID  *uint    // nil si es un combo
	ComboID    *uint    // nil si es un producto individual
	Quantity   float64  `gorm:"not null" validate:"gt=0"`
	UnitPrice  float64  `gorm:"not null" validate:"gt=0"`
	TotalPrice float64  `gorm:"not null" validate:"gt=0"`
	IsPartial  bool     `gorm:"default:false"`
	Product    *Product `gorm:"foreignKey:ProductID"`
	Combo      *Combo   `gorm:"foreignKey:ComboID"`
}

type ProductSales struct {
	ProductID     uint    `json:"product_id"`
	TotalQuantity float64 `json:"total_quantity"`
	TotalSales    float64 `json:"total_sales"`
	ProductName   string  `json:"product_name" gorm:"-"`
}

type PriceRangeSales struct {
	PriceRange string  `json:"price_range"`
	Count      int     `json:"count"`
	Total      float64 `json:"total"`
}

type SaleComment struct {
	gorm.Model
	SaleID     uint      `gorm:"not null" validate:"required"`
	Author     string    `gorm:"not null"` // usuario que hizo el comentario
	Content    string    `gorm:"not null" validate:"required,min=5"`
	CommentAt  time.Time `gorm:"not null"`
	IsInternal bool      `gorm:"default:false"` // PAra comentaros visibles solo al personal
}
