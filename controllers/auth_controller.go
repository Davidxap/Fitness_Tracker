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

// LoginHandler valida credenciales y emite JWT
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Buscamos usuario por email
	var u models.User
	err := database.DB.QueryRow(
		"SELECT id, password FROM users WHERE email=$1", req.Email,
	).Scan(&u.ID, &u.Password)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Comparamos password (texto plano)
	if req.Password != u.Password {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Generamos JWT
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}

	// Devolvemos el token en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
