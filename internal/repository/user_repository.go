package repository

import (
	"context"
	"errors"

	"go-cloud-api/internal/model"

	"github.com/jackc/pgx/v5"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
	Create(ctx context.Context, u model.User) (model.User, error)
}

type userRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email
		FROM users
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (model.User, error) {
	var u model.User
	err := r.db.QueryRow(ctx, `
		SELECT id, name, email
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Name, &u.Email)

	if err == pgx.ErrNoRows {
		return model.User{}, ErrUserNotFound
	}
	return u, err
}

func (r *userRepository) Create(ctx context.Context, u model.User) (model.User, error) {
	err := r.db.QueryRow(ctx, `
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id
	`, u.Name, u.Email).Scan(&u.ID)

	return u, err
}
