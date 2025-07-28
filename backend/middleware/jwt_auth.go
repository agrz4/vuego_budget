package middleware

import (
	"budget_app/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token autentikasi diperlukan"})
			c.Abort()
			return
		}

		// Token expected format: "Bearer <token>"
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		if tokenString == authHeader { // No "Bearer " prefix found
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}

		userID, err := utils.ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid atau kadaluarsa"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
