package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generates a JWT token with userID in the claims
func GenerateToken(userID int) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	// Define claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // expires in 24h
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Validates the token and returns the claims
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
}