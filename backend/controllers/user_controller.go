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

// GetUsers lists all users (password omitted)
func GetUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(
		"SELECT id, name, email, age, weight, created_at FROM users",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Weight, &u.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID retrieves user data
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)  
		return
	}

	var u models.User
	err = database.DB.QueryRow(
		"SELECT id, name, email, age, weight, created_at FROM users WHERE id=$1",
		id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.Weight, &u.CreatedAt)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)  
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// CreateUser creates a user with age and weight
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)  
		return
	}

	err := database.DB.QueryRow(
		`INSERT INTO users(name,email,password,age,weight)
		 VALUES($1,$2,$3,$4,$5) RETURNING id,created_at`,
		u.Name, u.Email, u.Password, u.Age, u.Weight,
	).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// UpdateUser updates name, email, password, age, weight
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)  
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)  
		return
	}

	res, err := database.DB.Exec(
		`UPDATE users
		 SET name=$1, email=$2, password=$3, age=$4, weight=$5
		 WHERE id=$6`,
		u.Name, u.Email, u.Password, u.Age, u.Weight, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "User not found", http.StatusNotFound)  
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser deletes user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)  
		return
	}

	res, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if cnt, _ := res.RowsAffected(); cnt == 0 {
		http.Error(w, "User not found", http.StatusNotFound)  
		return
	}

	w.WriteHeader(http.StatusNoContent)
}