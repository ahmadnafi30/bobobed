package repository

import (
	"database/sql"
	"errors"

	"github.com/ahmadnafi30/bobobed/backend/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

// --- In-Memory Implementation ---
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

// --- PostgreSQL Implementation ---
type PostgresUserRepo struct {
	db *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) CreateUser(user *entity.User) error {
	// Cek apakah email sudah ada
	var existingUser entity.User
	err := r.db.QueryRow("SELECT id FROM users WHERE email = $1", user.Email).Scan(&existingUser.ID)
	if err == nil {
		// Jika hasil query tidak mengembalikan error, berarti email sudah ada
		return errors.New("email already exists")
	}
	if err != sql.ErrNoRows {
		// Jika ada error lain, kembalikan error tersebut
		return err
	}

	// Insert user baru
	query := `INSERT INTO users (first_name, last_name, email, password, phone)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err = r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepo) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, first_name, last_name, email, password, phone FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
