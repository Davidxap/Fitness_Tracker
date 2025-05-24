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
		`SELECT id, user_id, name, date, duration_minutes, observations, created_at 
		 FROM workout_sessions`,
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
			&s.ID,
			&s.UserID,
			&s.Name,
			&s.Date,
			&s.DurationMinutes,
			&s.Observations,
			&s.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sessions = append(sessions, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetWorkoutSessionByID obtiene una sesión según su {id}.
func GetWorkoutSessionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	var s models.WorkoutSession
	err = database.DB.QueryRow(
		`SELECT id, user_id, name, date, duration_minutes, observations, created_at 
		 FROM workout_sessions WHERE id=$1`, id,
	).Scan(
		&s.ID,
		&s.UserID,
		&s.Name,
		&s.Date,
		&s.DurationMinutes,
		&s.Observations,
		&s.CreatedAt,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

// CreateWorkoutSession inserta una nueva sesión en la base de datos.
func CreateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	var s models.WorkoutSession
	// Decodifica JSON del body en el struct
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Insert con todos los campos name, date, duration, observations, user_id
	err := database.DB.QueryRow(
		`INSERT INTO workout_sessions 
		 (user_id, name, date, duration_minutes, observations)
		 VALUES($1,$2,$3,$4,$5)
		 RETURNING id, created_at`,
		s.UserID,
		s.Name,
		s.Date,
		s.DurationMinutes,
		s.Observations,
	).Scan(&s.ID, &s.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolvemos la sesión creada (incluye ID y created_at)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateWorkoutSession actualiza los campos de una sesión existente.
func UpdateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	var s models.WorkoutSession
	// Decodifica JSON con campos a actualizar
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Ejecuta UPDATE incluyendo name y observations
	res, err := database.DB.Exec(
		`UPDATE workout_sessions
		 SET user_id=$1, name=$2, date=$3, duration_minutes=$4, observations=$5
		 WHERE id=$6`,
		s.UserID,
		s.Name,
		s.Date,
		s.DurationMinutes,
		s.Observations,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verifica que se actualizó al menos una fila
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// DeleteWorkoutSession elimina una sesión por su {id}.
func DeleteWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	// Elimina la sesión
	res, err := database.DB.Exec("DELETE FROM workout_sessions WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verifica que exista
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Session not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
