package models

import "time"

type Exercise struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	MuscleGroup string    `json:"muscle_group"`
	CreatedAt   time.Time `json:"created_at"`
}
