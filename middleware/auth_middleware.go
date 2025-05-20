package middleware

import (
	"net/http"
	"strings"

	"fitness-tracker/backend/utils"
)

// Middleware para rutas protegidas
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cabecera Authorization: Bearer <token>
		header := r.Header.Get("Authorization")
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Token requerido", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		token, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}
		// Puedes extraer claims si las necesitas:
		// claims := token.Claims.(jwt.MapClaims)
		next.ServeHTTP(w, r)
	})
}
