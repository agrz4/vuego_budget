package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTSecret key for signing and verifying tokens
var JWTSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims defines the structure of the JWT claims
type Claims struct {
	UserID               uint `json:"user_id"`
	jwt.RegisteredClaims      // Mengganti jwt.StandardClaims dengan jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for a given user ID
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{ // Menggunakan RegisteredClaims
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Menggunakan NewNumericDate
			IssuedAt:  jwt.NewNumericDate(time.Now()),     // Menggunakan NewNumericDate
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", errors.New("Gagal menandatangani token")
	}
	return tokenString, nil
}

// ParseJWT parses and validates a JWT token string
func ParseJWT(tokenString string) (uint, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Memastikan metode penandatanganan sesuai
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Metode penandatanganan tidak valid")
		}
		return JWTSecret, nil
	})

	if err != nil {
		// Menangani berbagai jenis kesalahan validasi token
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return 0, errors.New("Tanda tangan token tidak valid")
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, errors.New("Token kadaluarsa")
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return 0, errors.New("Token belum berlaku")
		}
		return 0, errors.New("Token tidak valid: " + err.Error())
	}

	if !token.Valid {
		return 0, errors.New("Token tidak valid")
	}

	return claims.UserID, nil
}
