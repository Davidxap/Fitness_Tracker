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

// GetUsers responde con la lista de todos los usuarios.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Ejecutamos la consulta
	rows, err := database.DB.Query(
		"SELECT id, name, email, password, created_at FROM users",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	// Iteramos cada fila y la "scaneamos" en un User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(
			&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	// Cabecera JSON + envío
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID devuelve un único usuario según su {id}.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extraemos el parámetro "id" de la URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var u models.User
	// Consulta parametrizada para evitar inyección SQL
	err = database.DB.QueryRow(
		"SELECT id, name, email, password, created_at FROM users WHERE id=$1",
		id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)

	if err == sql.ErrNoRows {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// CreateUser crea un nuevo usuario leyendo el cuerpo JSON.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	// Decodificamos JSON a struct
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Insert y devolver id + timestamp
	err := database.DB.QueryRow(
		"INSERT INTO users(name,email,password) VALUES($1,$2,$3) RETURNING id,created_at",
		u.Name, u.Email, u.Password,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// UpdateUser modifica campos de un usuario existente.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Ejecutamos UPDATE
	res, err := database.DB.Exec(
		"UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4",
		u.Name, u.Email, u.Password, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Verificar que algo cambió
	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 sin cuerpo
}

// DeleteUser elimina un usuario por su {id}.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
