package models

import (
	"gorm.io/gorm"
)

// Пользователи системы
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"type:varchar(20);default:'user'" json:"role"`
	Phone    string `gorm:"size:20" json:"phone"`
	Address  string `gorm:"size:255" json:"address"`
}

type UserSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
