package main

import (
	"net/http"

	"github.com/ahmadnafi30/bobobed/backend/internal/handler"
	"github.com/ahmadnafi30/bobobed/backend/internal/repository"
	"github.com/ahmadnafi30/bobobed/backend/internal/service"
)

func main() {
	userRepo := repository.NewInMemoryUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/register", userHandler.Register)
	http.HandleFunc("/login", userHandler.Login)

	http.ListenAndServe(":8080", nil)
}
