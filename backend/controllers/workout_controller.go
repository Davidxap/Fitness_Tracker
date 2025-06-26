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

// GetWorkoutSessions lists all workout sessions.
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

// GetWorkoutSessionByID retrieves a session by its {id}.
func GetWorkoutSessionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest) //  
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

// CreateWorkoutSession inserts a new session into the database.
func CreateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	var s models.WorkoutSession
	// Decode JSON from body into the struct
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)  
		return
	}

	// Insert with all fields name, date, duration, observations, user_id
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

	// Return the created session (includes ID and created_at)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// UpdateWorkoutSession updates the fields of an existing session.
func UpdateWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)  
		return
	}

	var s models.WorkoutSession
	// Decode JSON with fields to update
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)  
		return
	}

	// Execute UPDATE including name and observations
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

	// Check if at least one row was updated
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Session not found", http.StatusNotFound)  
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}

// DeleteWorkoutSession deletes a session by its {id}.
func DeleteWorkoutSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)  
		return
	}

	// Delete the session
	res, err := database.DB.Exec("DELETE FROM workout_sessions WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if it exists
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Session not found", http.StatusNotFound)  
		return
	}

	w.WriteHeader(http.StatusNoContent)
}