package routes

import (
	"fitness-tracker/backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Usuarios
	r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	// (añade PUT, DELETE)

	// Sesiones de entrenamiento
	r.HandleFunc("/api/sessions", controllers.GetWorkoutSessions).Methods("GET")
	r.HandleFunc("/api/sessions", controllers.CreateWorkoutSession).Methods("POST")
	// …

	// Ejercicios
	r.HandleFunc("/api/exercises", controllers.GetExercises).Methods("GET")
	r.HandleFunc("/api/exercises", controllers.CreateExercise).Methods("POST")
	// …

	// Detalles de sesión-ejercicio
	r.HandleFunc("/api/session-exercises", controllers.GetSessionExercises).Methods("GET")
	r.HandleFunc("/api/session-exercises", controllers.CreateSessionExercise).Methods("POST")
	// …

	// Ruta raíz de prueba
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fitness‑Tracker API 🚀"))
	}).Methods("GET")

	return r
}
