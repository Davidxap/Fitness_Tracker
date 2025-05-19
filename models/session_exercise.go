package models

import "time"

type SessionExercise struct {
	ID         int       `json:"id"`
	SessionID  int       `json:"session_id"`
	ExerciseID int       `json:"exercise_id"`
	Sets       int       `json:"sets"`
	Reps       int       `json:"reps"`
	Weight     float64   `json:"weight"`
	CreatedAt  time.Time `json:"created_at"`
}
