package repository

import (
	"errors"

	"github.com/ahmadnafi30/bobobed/backend/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type InMemoryUserRepo struct {
	users map[string]*entity.User
	id    int64
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]*entity.User),
		id:    1,
	}
}

func (r *InMemoryUserRepo) CreateUser(user *entity.User) error {
	if _, exists := r.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	user.ID = r.id
	r.id++
	r.users[user.Email] = user
	return nil
}

func (r *InMemoryUserRepo) FindByEmail(email string) (*entity.User, error) {
	user, exists := r.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
