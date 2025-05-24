package controllers

import (
	"encoding/json"
	"net/http"

	"fitness-tracker/backend/database"
	"fitness-tracker/backend/models"
	"fitness-tracker/backend/utils"
)

// LoginRequest recibe email + password
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse retorna token + user data
type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// LoginHandler valida credenciales y emite JWT + datos del usuario
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Seleccionamos todos los campos necesarios
	var u models.User
	err := database.DB.QueryRow(
		`SELECT id, name, email, password, age, weight, created_at
         FROM users WHERE email=$1`,
		req.Email,
	).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.Age,
		&u.Weight,
		&u.CreatedAt,
	)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Comparamos password en claro
	if req.Password != u.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generamos el token
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	u.Password = "" // no enviamos password en la respuesta

	resp := LoginResponse{
		Token: token,
		User:  u,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
