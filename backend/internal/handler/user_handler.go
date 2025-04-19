package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ahmadnafi30/bobobed/backend/entity"
	"github.com/ahmadnafi30/bobobed/backend/internal/service"
	"github.com/ahmadnafi30/bobobed/backend/model"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req model.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user := &entity.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err := h.Service.Register(user, req.ConfirmPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "register success"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "login success"})
}
