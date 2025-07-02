package models

import (
	"gorm.io/gorm"
	"time"
)

// Животное
type Animal struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"animal_id" gorm:"primaryKey"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	Type       string    `gorm:"size:50;not null" json:"type"`
	Breed      string    `gorm:"size:100" json:"breed"`
	Age        uint      `json:"age"`
	Gender     string    `gorm:"size:10" json:"gender"`
	TakenAt    time.Time `json:"taken_at"`
	IsActive   bool      `json:"is_active"`   // находиться ли приюте или нет
	IsReserved bool      `json:"is_reserved"` //была ли созданна заявка на усыновления этого животного
	StoreID    *uint     `json:"store_id"`
	Store      Store     `gorm:"foreignKey:StoreID" json:"-"`
}
