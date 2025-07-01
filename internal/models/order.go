package models

import (
	"gorm.io/gorm"
)

// Заказы
type Order struct {
	gorm.Model
	IsActive  bool    `gorm:"default:true" json:"is_active"` // true - активный, false - неактивный
	UserID    uint    `json:"user_id"`
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	StoreID   uint    `json:"store_id"`
	Store     Store   `gorm:"foreignKey:StoreID" json:"store"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int     `gorm:"default:1" json:"quantity"`
}
