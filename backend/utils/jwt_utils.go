package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Genera un token JWT con userID en las claims
func GenerateToken(userID int) (string, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	// Definimos claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // expira en 24h
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// Verifica el token y devuelve las claims
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar algoritmo
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
}
