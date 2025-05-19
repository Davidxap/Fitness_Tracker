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

// GetSessionExercises lista todas las asociaciones sesión↔ejercicio.
func GetSessionExercises(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		`SELECT id, session_id, exercise_id, sets, reps, weight, created_at
    FROM session_exercises`,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.SessionExercise
	for rows.Next() {
		var se models.SessionExercise
		if err := rows.Scan(
			&se.ID, &se.SessionID, &se.ExerciseID,
			&se.Sets, &se.Reps, &se.Weight, &se.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, se)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetSessionExerciseByID devuelve un registro según {id}.
func GetSessionExerciseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var se models.SessionExercise
	err = database.DB.QueryRow(
		`SELECT id, session_id, exercise_id, sets, reps, weight, created_at
        FROM session_exercises WHERE id=$1`, id,
	).Scan(
		&se.ID, &se.SessionID, &se.ExerciseID,
		&se.Sets, &se.Reps, &se.Weight, &se.CreatedAt,
	)
	if err == sql.ErrNoRows {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(se)
}

// CreateSessionExercise inserta una nueva asociación.
func CreateSessionExercise(w http.ResponseWriter, r *http.Request) {
	var se models.SessionExercise
	if err := json.NewDecoder(r.Body).Decode(&se); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow(
		`INSERT INTO session_exercises(session_id,exercise_id,sets,reps,weight)
         VALUES($1,$2,$3,$4,$5) RETURNING id,created_at`,
		se.SessionID, se.ExerciseID, se.Sets, se.Reps, se.Weight,
	).Scan(&se.ID, &se.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(se)
}

// UpdateSessionExercise modifica un registro existente.
func UpdateSessionExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var se models.SessionExercise
	if err := json.NewDecoder(r.Body).Decode(&se); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec(
		`UPDATE session_exercises
        SET session_id=$1, exercise_id=$2, sets=$3, reps=$4, weight=$5
        WHERE id=$6`,
		se.SessionID, se.ExerciseID, se.Sets, se.Reps, se.Weight, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteSessionExercise elimina por {id}.
func DeleteSessionExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("DELETE FROM session_exercises WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Registro no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
