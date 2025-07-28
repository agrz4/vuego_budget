package controllers

import (
	"budget_app/database"
	"budget_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBudget(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	budget.UserID = userID

	if result := database.DB.Create(&budget); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat anggaran", "details": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Anggaran berhasil dibuat", "budget": budget})
}

func GetBudgets(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var budgets []models.Budget
	if result := database.DB.Where("user_id = ?", userID).Find(&budgets); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil anggaran", "details": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budgets": budgets})
}

func GetBudgetByID(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var budget models.Budget
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&budget); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggaran tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"budget": budget})
}

func UpdateBudget(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var budget models.Budget
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&budget); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggaran tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	budget.UserID = userID // Pastikan UserID tidak berubah

	if result := database.DB.Save(&budget); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui anggaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Anggaran berhasil diperbarui", "budget": budget})
}

func DeleteBudget(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var budget models.Budget
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&budget); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anggaran tidak ditemukan"})
		return
	}

	if result := database.DB.Delete(&budget); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus anggaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Anggaran berhasil dihapus"})
}

func GetDashboardSummary(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Total Budget
	var totalBudget float64
	database.DB.Model(&models.Budget{}).Where("user_id = ?", userID).Select("sum(amount)").Scan(&totalBudget)

	// Total Expense
	var totalExpense float64
	database.DB.Model(&models.Expense{}).Where("user_id = ?", userID).Select("sum(amount)").Scan(&totalExpense)

	// Expenses by Category for chart
	var expensesByCategory []struct {
		Category string  `json:"category"`
		Total    float64 `json:"total"`
	}
	database.DB.Model(&models.Expense{}).Where("user_id = ?", userID).
		Select("category, sum(amount) as total").
		Group("category").
		Scan(&expensesByCategory)

	c.JSON(http.StatusOK, gin.H{
		"total_budget":         totalBudget,
		"total_expense":        totalExpense,
		"remaining_budget":     totalBudget - totalExpense,
		"expenses_by_category": expensesByCategory,
	})
}
