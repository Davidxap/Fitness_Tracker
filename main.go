package main

import (
	"fitness-tracker/backend/database"
	"fitness-tracker/backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// 1. Carga .env (requerido para SERVER_PORT)
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: no se encontró .env, se usarán variables de entorno del sistema")
	}
	// 2. Inicializa la DB
	database.InitDB()

	// 3. Registra rutas
	router := routes.RegisterRoutes()

	// 4. Arranca el servidor en el puerto de .env
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // fallback
	}
	fmt.Printf("🚀 Servidor escuchando en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
