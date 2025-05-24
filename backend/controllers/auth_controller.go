package controllers

import (
	"encoding/json"
	"net/http"

	"fitness-tracker/backend/database"
	"fitness-tracker/backend/models"
	"fitness-tracker/backend/utils"
)

// LoginRequest estructura para recibir email+password
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse envía token + datos de usuario
type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// LoginHandler valida credenciales y emite JWT + User
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Buscamos usuario por email
	var u models.User
	err := database.DB.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email=$1",
		req.Email,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Comparamos password (plano)
	if req.Password != u.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generamos JWT
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// No devolvemos la contraseña
	u.Password = ""

	// Armamos respuesta
	resp := LoginResponse{
		Token: token,
		User:  u,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
