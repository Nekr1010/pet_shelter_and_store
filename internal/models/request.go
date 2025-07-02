package models

import (
	"gorm.io/gorm"
)

// Заявки
type Request struct {
	gorm.Model
	IsAdoption bool   `gorm:"default:true" json:"is_adoption"` // true - усыновление, false - сдача
	IsAccepted bool   `gorm:"default:true" json:"is_active"`   // true - активная, false - неактивная
	UserID     uint   `json:"user_id"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	StoreID    *uint  `json:"store_id"`
	Store      Store  `gorm:"foreignKey:StoreID" json:"store"`
	AnimalID   uint   `json:"animal_id"`
	Animal     Animal `gorm:"foreignKey:AnimalID" json:"animal"`
}
