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

// InitDB carga .env y abre la conexión global
func InitDB() {
	// 1. Carga variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar .env: %v", err)
	}
	// 2. Construye la cadena de conexión
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// 3. Abre conexión
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a DB: %v", err)
	}
	// 4. Verifica con Ping
	if err = DB.Ping(); err != nil {
		log.Fatalf("Ping a DB fallido: %v", err)
	}
	fmt.Println("✅ Conectado a PostgreSQL")
}
