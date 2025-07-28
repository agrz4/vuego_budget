package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"user_id"`
	BudgetID    uint      `json:"budget_id"` // Opsional, bisa dikaitkan dengan Budget
	Category    string    `gorm:"not null" json:"category" binding:"required"`
	Amount      float64   `gorm:"not null" json:"amount" binding:"required,gt=0"`
	Description string    `json:"description"`
	Date        time.Time `gorm:"not null" json:"date" binding:"required"`
}
