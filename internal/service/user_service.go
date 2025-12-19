package service

import (
	"errors"

	"go-cloud-api/internal/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id int) (model.User, error)
	Create(name, email string) (model.User, error)
}

type userService struct {
	users []model.User
}

func NewUserService() UserService {
	return &userService{
		users: []model.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}
func (s *userService) GetAll() ([]model.User, error) {
	return s.users, nil
}
func (s *userService) GetByID(id int) (model.User, error) {
	for _, u := range s.users {
		if u.ID == id {
			return u, nil
		}
	}
	return model.User{}, ErrUserNotFound
}
func (s *userService) Create(name, email string) (model.User, error) {
	u := model.User{
		ID:    len(s.users) + 1,
		Name:  name,
		Email: email,
	}
	s.users = append(s.users, u)
	return u, nil
}
func NewUserServiceWithSeed(seed []model.User) UserService {
	usersCopy := make([]model.User, len(seed))
	copy(usersCopy, seed)

	return &userService{
		users: usersCopy,
	}
}
