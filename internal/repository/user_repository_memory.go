package repository

import (
	"context"

	"go-cloud-api/internal/model"
)

type InMemoryUserRepository struct {
	users []model.User
}

func NewInMemoryUserRepository(seed []model.User) *InMemoryUserRepository {
	usersCopy := make([]model.User, len(seed))
	copy(usersCopy, seed)
	return &InMemoryUserRepository{users: usersCopy}
}

func (r *InMemoryUserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	out := make([]model.User, len(r.users))
	copy(out, r.users)
	return out, nil
}

func (r *InMemoryUserRepository) GetByID(ctx context.Context, id int) (model.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, ErrUserNotFound
}

func (r *InMemoryUserRepository) Create(ctx context.Context, u model.User) (model.User, error) {
	u.ID = len(r.users) + 1
	r.users = append(r.users, u)
	return u, nil
}
