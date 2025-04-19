package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ahmadnafi30/bobobed/backend/internal/handler"
	"github.com/ahmadnafi30/bobobed/backend/internal/repository"
	"github.com/ahmadnafi30/bobobed/backend/internal/service"
	_ "github.com/lib/pq"
	"github.com/rs/cors" // Import package CORS
)

func main() {
	// Koneksi ke database PostgreSQL
	connStr := "host=localhost port=5432 user=postgres password=mypassword123 dbname=bobobed_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Verifikasi koneksi ke database
	if err := db.Ping(); err != nil {
		log.Fatal("Tidak bisa terhubung ke database: ", err)
	}

	// Membuat instance repository dan service
	userRepo := repository.NewPostgresUserRepo(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Mengatur route HTTP
	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Izinkan semua origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(http.DefaultServeMux)

	// Jalankan server
	log.Fatal(http.ListenAndServe(":8081", handler))
}
print