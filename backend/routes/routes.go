package routes

import (
	"net/http"

	"fitness-tracker/backend/controllers"
	"fitness-tracker/backend/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Health-check público
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fitness‑Tracker API 🚀"))
	}).Methods("GET")

	// ----- Autenticación público -----
	r.HandleFunc("/api/login", controllers.LoginHandler).Methods("POST")
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST") // registro

	// ----- Rutas protegidas: requiere JWT -----
	auth := r.NewRoute().Subrouter()
	auth.Use(middleware.JWTAuth)

	// Usuarios (excepto Create)
	auth.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	auth.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	auth.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	auth.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Sesiones
	auth.HandleFunc("/api/sessions", controllers.GetWorkoutSessions).Methods("GET")
	auth.HandleFunc("/api/sessions", controllers.CreateWorkoutSession).Methods("POST")
	auth.HandleFunc("/api/sessions/{id}", controllers.GetWorkoutSessionByID).Methods("GET")
	auth.HandleFunc("/api/sessions/{id}", controllers.UpdateWorkoutSession).Methods("PUT")
	auth.HandleFunc("/api/sessions/{id}", controllers.DeleteWorkoutSession).Methods("DELETE")

	// Ejercicios
	auth.HandleFunc("/api/exercises", controllers.GetExercises).Methods("GET")
	auth.HandleFunc("/api/exercises", controllers.CreateExercise).Methods("POST")
	auth.HandleFunc("/api/exercises/{id}", controllers.GetExerciseByID).Methods("GET")
	auth.HandleFunc("/api/exercises/{id}", controllers.UpdateExercise).Methods("PUT")
	auth.HandleFunc("/api/exercises/{id}", controllers.DeleteExercise).Methods("DELETE")

	// Session-exercises
	auth.HandleFunc("/api/session-exercises", controllers.GetSessionExercises).Methods("GET")
	auth.HandleFunc("/api/session-exercises", controllers.CreateSessionExercise).Methods("POST")
	auth.HandleFunc("/api/session-exercises/{id}", controllers.GetSessionExerciseByID).Methods("GET")
	auth.HandleFunc("/api/session-exercises/{id}", controllers.UpdateSessionExercise).Methods("PUT")
	auth.HandleFunc("/api/session-exercises/{id}", controllers.DeleteSessionExercise).Methods("DELETE")

	return r
}
