package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/timut2/avito_test_task/internal/handlers"
	"github.com/timut2/avito_test_task/internal/repository"
	"github.com/timut2/avito_test_task/internal/service"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=123123")

	if err != nil {
		log.Fatalf("Failed to connect to database: %v ", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	// Пример выполнения запроса
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatalf("Failed to query database version: %v", err)
	}

	fmt.Println(version)

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)
	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth", authHandler.Login)
	fmt.Println("connected to db")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
