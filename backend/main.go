package main

import (
	"budget_app/controllers"
	"budget_app/database"
	"budget_app/middleware"
	"budget_app/models"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{}, &models.Budget{}, &models.Expense{})

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	// Konfigurasi CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Izinkan frontend Vue.js
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	router.Use(cors.New(config))

	router.POST("api/register", controllers.Register)
	router.POST("api/login", controllers.Login)

	authorized := router.Group("/api")
	authorized.Use(middleware.JWTMiddleware())
	{
		// Rute Budget
		authorized.POST("/budgets", controllers.CreateBudget)
		authorized.GET("/budgets", controllers.GetBudgets)
		authorized.GET("/budgets/:id", controllers.GetBudgetByID)
		authorized.PUT("/budgets/:id", controllers.UpdateBudget)
		authorized.DELETE("/budgets/:id", controllers.DeleteBudget)

		// Rute Expense
		authorized.POST("/expenses", controllers.CreateExpense)
		authorized.GET("/expenses", controllers.GetExpenses) // Termasuk filter dan pencarian
		authorized.GET("/expenses/:id", controllers.GetExpenseByID)
		authorized.PUT("/expenses/:id", controllers.UpdateExpense)
		authorized.DELETE("/expenses/:id", controllers.DeleteExpense)

		// Rute Dashboard (ringkasan)
		authorized.GET("/dashboard", controllers.GetDashboardSummary)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di http://localhost:%s", port)
	router.Run(":" + port)
}
