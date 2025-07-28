package controllers

import (
	"budget_app/database"
	"budget_app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateExpense(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var expense models.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense.UserID = userID

	if result := database.DB.Create(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pengeluaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pengeluaran berhasil dibuat", "expense": expense})
}

func GetExpenses(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var expenses []models.Expense
	query := database.DB.Where("user_id = ?", userID)

	// Filter by category
	if category := c.Query("category"); category != "" {
		query = query.Where("category ILIKE ?", "%"+category+"%") // ILIKE for case-insensitive search in PostgreSQL
	}

	// Filter by date range
	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err == nil {
			query = query.Where("date >= ?", startDate)
		}
	}
	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err == nil {
			// Add 23 hours, 59 minutes, 59 seconds to include the whole end day
			endDate = endDate.Add(24*time.Hour - time.Second)
			query = query.Where("date <= ?", endDate)
		}
	}

	// Search by description
	if search := c.Query("search"); search != "" {
		query = query.Where("description ILIKE ? OR category ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Order by date descending by default
	query = query.Order("date desc")

	if result := query.Find(&expenses); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil pengeluaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expenses": expenses})
}

func GetExpenseByID(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var expense models.Expense
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengeluaran tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"expense": expense})
}

func UpdateExpense(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var expense models.Expense
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengeluaran tidak ditemukan"})
		return
	}

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense.UserID = userID // Pastikan UserID tidak berubah

	if result := database.DB.Save(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pengeluaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengeluaran berhasil diperbarui", "expense": expense})
}

func DeleteExpense(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	id := c.Param("id")
	var expense models.Expense
	if result := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&expense); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengeluaran tidak ditemukan"})
		return
	}

	if result := database.DB.Delete(&expense); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pengeluaran", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengeluaran berhasil dihapus"})
}
