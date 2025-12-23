package service

import (
	"context"
	"errors"

	"go-cloud-api/internal/model"
	repository "go-cloud-api/internal/repository"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserService interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id int) (model.User, error)
	Create(ctx context.Context, name, email string) (model.User, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}
func (s *userService) GetByID(ctx context.Context, id int) (model.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return u, nil
}
func (s *userService) Create(ctx context.Context, name, email string) (model.User, error) {
	u := model.User{
		Name:  name,
		Email: email,
	}
	return s.repo.Create(ctx, u)
}
