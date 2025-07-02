package models

import (
	"gorm.io/gorm"
)

// Товары
type Product struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Quantity    int     `gorm:"default:0" json:"quantity"`
	StoreID     uint    `json:"store_id"`
	Store       Store   `gorm:"foreignKey:StoreID" json:"store"`
}
