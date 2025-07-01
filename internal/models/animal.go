package models

import (
	"gorm.io/gorm"
)

// Животное
type Animal struct {
	gorm.Model
	Name    string `gorm:"size:100;not null" json:"name"`
	Type    string `gorm:"size:50;not null" json:"type"`
	Breed   string `gorm:"size:100" json:"breed"`
	Age     uint   `json:"age"`
	Gender  string `gorm:"size:10" json:"gender"`
	StoreID *uint  `json:"store_id"`
	Store   Store  `gorm:"foreignKey:StoreID" json:"store"`
}
