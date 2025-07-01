package models

import (
	"gorm.io/gorm"
)

// Магазин(приют если включен)
type Store struct {
	gorm.Model
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Address     string    `gorm:"size:255;not null" json:"address"`
	Phone       string    `gorm:"size:20;not null" json:"phone"`
	Email       string    `gorm:"size:100" json:"email"`
	IsShelter   bool      `gorm:"default:false" json:"is_shelter"`
	OwnerID     uint      `json:"owner_id"`
	Owner       User      `gorm:"foreignKey:OwnerID" json:"owner"`
	Animals     []Animal  `gorm:"foreignKey:StoreID" json:"animals"`
	Products    []Product `gorm:"foreignKey:StoreID" json:"products"`
}
