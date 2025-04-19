package service

import (
	"errors"

	"github.com/ahmadnafi30/bobobed/backend/entity"
	"github.com/ahmadnafi30/bobobed/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Register(user *entity.User, confirmPassword string) error {
	if user.Password != confirmPassword {
		return errors.New("passwords do not match")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)
	return s.Repo.CreateUser(user)
}

func (s *UserService) Login(email, password string) error {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return errors.New("invalid password")
	}

	return nil
}
