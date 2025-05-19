package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"fitness-tracker/backend/database"
	"fitness-tracker/backend/models"

	"github.com/gorilla/mux"
)

// GetWorkoutSessions lista todas las sesiones de entrenamiento.
func GetWorkoutSessions(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		"SELECT id, user_id, date, duration_minutes, created_at FROM workout_sessions",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sessions []models.WorkoutSession
	for rows.Next() {
		var s models.WorkoutSession
		if err := rows.Scan(
			&s.ID, &s.UserID, &s.Date, &s.DurationMinutes, &s.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessions = append(sessions, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetWorkoutSessionByID obtiene una sesión según {id}.
func GetWorkoutSessionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var s models.WorkoutSession
	err = database.DB.QueryRow(
		"SELECT id, user_id, date, duration_minutes, created_at FROM workout_sessions WHERE id=$1",
		id,
	).Scan(&s.ID, &s.UserID, &s.Date, &s.DurationMinutes, &s.CreatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Sesión no encontrada", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateWorkoutSession inserta nueva sesión.
func CreateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	var s models.WorkoutSession
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO workout_sessions(user_id,date,duration_minutes) VALUES($1,$2,$3) RETURNING id,created_at",
		s.UserID, s.Date, s.DurationMinutes,
	).Scan(&s.ID, &s.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateWorkoutSession edita los campos de una sesión existente.
func UpdateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var s models.WorkoutSession
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec(
		"UPDATE workout_sessions SET user_id=$1, date=$2, duration_minutes=$3 WHERE id=$4",
		s.UserID, s.Date, s.DurationMinutes, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Sesión no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteWorkoutSession elimina una sesión por {id}.
func DeleteWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("DELETE FROM workout_sessions WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Sesión no encontrada", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
