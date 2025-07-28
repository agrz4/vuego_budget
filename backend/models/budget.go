package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	UserID      uint    `gorm:"not null" json:"user_id"`
	Category    string  `gorm:"not null" json:"category" binding:"required"`
	Amount      float64 `gorm:"not null" json:"amount" binding:"required,gt=0"`
	Description string  `json:"description"`
}
