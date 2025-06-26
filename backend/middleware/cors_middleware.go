package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// CORSHandler wraps your router to allow domains
func CORSHandler(h http.Handler) http.Handler {
	// Allow all (*) for testing; restrict origins in production
	return handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
	)(h)
}