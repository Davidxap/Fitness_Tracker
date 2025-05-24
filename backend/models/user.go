package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	Age       int       `json:"age"`
	Weight    float64   `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}
