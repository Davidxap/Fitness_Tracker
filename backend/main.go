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
	// 1) Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, system environment variables will be used")
	}

	// 2) DB Connection
	database.InitDB()

	// 3) Register routes
	router := routes.RegisterRoutes()

	// 4) Wrap with CORS
	handler := middleware.CORSHandler(router)

	// 5) Start server
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}