package models

import "time"

type WorkoutSession struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Name            string    `json:"name"`
	Date            string    `json:"date"`
	DurationMinutes int       `json:"duration_minutes"`
	Observations    string    `json:"observations,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}
