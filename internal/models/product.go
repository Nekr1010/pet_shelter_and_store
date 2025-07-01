package models

import (
	"gorm.io/gorm"
)

// Товары
type Product struct {
	gorm.Model
	Name        string        `gorm:"size:100;not null" json:"name"`
	Description string        `gorm:"type:text" json:"description"`
	Price       float64       `gorm:"type:decimal(10,2);not null" json:"price"`
	Quantity    int           `gorm:"default:0" json:"quantity"`
	StoreID     uint          `json:"store_id"`
	Store       Store         `gorm:"foreignKey:StoreID" json:"store"`
	ImageURL    string        `gorm:"size:255" json:"image_url"`
	Status      ProductStatus `gorm:"type:varchar(20);default:'available'" json:"status"`
}

// Статусы товаров
type ProductStatus string

const (
	ProductAvailable ProductStatus = "available"
	ProductReserved  ProductStatus = "reserved"
	ProductSoldOut   ProductStatus = "sold_out"
)
