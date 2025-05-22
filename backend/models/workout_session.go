package models

import "time"

type WorkoutSession struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Date            string    `json:"date"`
	DurationMinutes int       `json:"duration_minutes"`
	CreatedAt       time.Time `json:"created_at"`
}
