package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB loads .env and opens the global connection
func InitDB() {
	// 1. Load variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %v", err) // Original was in Spanish (log.Fatalf("Error al cargar .env: %v", err))
	}
	// 2. Build the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// 3. Open connection
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err) // Original was in Spanish (log.Fatalf("No se pudo conectar a DB: %v", err))
	}
	// 4. Verify with Ping
	if err = DB.Ping(); err != nil {
		log.Fatalf("DB Ping failed: %v", err) // Original was in Spanish (log.Fatalf("Ping a DB fallido: %v", err))
	}
	fmt.Println("✅ Connected to PostgreSQL") // Original was in Spanish (fmt.Println("✅ Conectado a PostgreSQL"))
}