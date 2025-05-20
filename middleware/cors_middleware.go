package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSHandler envuelve tu router para permitir dominios
func CORSHandler(h http.Handler) http.Handler {
	// Permitir todo (*) para pruebas; en producción restringe orígenes
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(h)
}
