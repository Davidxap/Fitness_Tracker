package middleware

import (
	"net/http"
	"strings"

	"fitness-tracker/backend/utils"
)

// Middleware for protected routes
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorization Header: Bearer <token>
		header := r.Header.Get("Authorization")
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Token required", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		// You can extract claims if you need them:
		// claims := token.Claims.(jwt.MapClaims)
		next.ServeHTTP(w, r)
	})
}