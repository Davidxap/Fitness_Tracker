package routes

import (
	"fitness-tracker/backend/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Ruta raíz de prueba
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Fitness‑Tracker API 🚀"))
	}).Methods("GET")

	// Usuarios
	r.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// Sesiones de entrenamiento
	r.HandleFunc("/api/sessions", controllers.GetWorkoutSessions).Methods("GET")
	r.HandleFunc("/api/sessions", controllers.CreateWorkoutSession).Methods("POST")
	r.HandleFunc("/api/sessions/{id}", controllers.GetWorkoutSessionByID).Methods("GET")
	r.HandleFunc("/api/sessions/{id}", controllers.UpdateWorkoutSession).Methods("PUT")
	r.HandleFunc("/api/sessions/{id}", controllers.DeleteWorkoutSession).Methods("DELETE")

	// Ejercicios
	r.HandleFunc("/api/exercises", controllers.GetExercises).Methods("GET")
	r.HandleFunc("/api/exercises", controllers.CreateExercise).Methods("POST")
	r.HandleFunc("/api/exercises/{id}", controllers.GetExerciseByID).Methods("GET")
	r.HandleFunc("/api/exercises/{id}", controllers.UpdateExercise).Methods("PUT")
	r.HandleFunc("/api/exercises/{id}", controllers.DeleteExercise).Methods("DELETE")

	// Detalles de sesión-ejercicio
	r.HandleFunc("/api/session-exercises", controllers.GetSessionExercises).Methods("GET")
	r.HandleFunc("/api/session-exercises", controllers.CreateSessionExercise).Methods("POST")
	r.HandleFunc("/api/session-exercises/{id}", controllers.GetSessionExerciseByID).Methods("GET")
	r.HandleFunc("/api/session-exercises/{id}", controllers.UpdateSessionExercise).Methods("PUT")
	r.HandleFunc("/api/session-exercises/{id}", controllers.DeleteSessionExercise).Methods("DELETE")

	return r
}
