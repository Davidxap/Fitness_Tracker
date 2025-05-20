package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"fitness-tracker/backend/database"
	"fitness-tracker/backend/routes"

	"fitness-tracker/backend/middleware"

	"github.com/joho/godotenv"
)

func main() {
	// 1) Carga .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró .env, se usarán variables de entorno del sistema")
	}

	// 2) Conexión BD
	database.InitDB()

	// 3) Registrar rutas
	router := routes.RegisterRoutes()

	// 4) Envolver con CORS
	handler := middleware.CORSHandler(router)

	// 5) Arrancar servidor
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("🚀 Servidor escuchando en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
