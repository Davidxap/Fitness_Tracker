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

// GetExercises lista todos los ejercicios del catálogo.
func GetExercises(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		"SELECT id, name, description, muscle_group, created_at FROM exercises",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var e models.Exercise
		if err := rows.Scan(
			&e.ID, &e.Name, &e.Description, &e.MuscleGroup, &e.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		exercises = append(exercises, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

// GetExerciseByID devuelve un ejercicio según {id}.
func GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var e models.Exercise
	err = database.DB.QueryRow(
		"SELECT id, name, description, muscle_group, created_at FROM exercises WHERE id=$1",
		id,
	).Scan(&e.ID, &e.Name, &e.Description, &e.MuscleGroup, &e.CreatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Ejercicio no encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

// CreateExercise inserta un nuevo ejercicio.
func CreateExercise(w http.ResponseWriter, r *http.Request) {
	var e models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := database.DB.QueryRow(
		"INSERT INTO exercises(name,description,muscle_group) VALUES($1,$2,$3) RETURNING id,created_at",
		e.Name, e.Description, e.MuscleGroup,
	).Scan(&e.ID, &e.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

// UpdateExercise modifica un ejercicio existente.
func UpdateExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var e models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec(
		"UPDATE exercises SET name=$1, description=$2, muscle_group=$3 WHERE id=$4",
		e.Name, e.Description, e.MuscleGroup, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Ejercicio no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteExercise elimina un ejercicio según {id}.
func DeleteExercise(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("DELETE FROM exercises WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "Ejercicio no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
