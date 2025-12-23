package service

import (
	"context"
	"testing"

	"go-cloud-api/internal/model"
	"go-cloud-api/internal/repository"
)

func TestUserService_GetAll(t *testing.T) {
	repo := repository.NewInMemoryUserRepository([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
		{ID: 2, Name: "Bob", Email: "bob@test.com"},
	})
	svc := NewUserService(repo)

	users, err := svc.GetAll(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestUserService_GetByID_Found(t *testing.T) {
	repo := repository.NewInMemoryUserRepository([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
	})
	svc := NewUserService(repo)

	u, err := svc.GetByID(context.Background(), 1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.ID != 1 {
		t.Fatalf("expected ID=1, got %d", u.ID)
	}
}

func TestUserService_GetByID_NotFound(t *testing.T) {
	repo := repository.NewInMemoryUserRepository([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
	})
	svc := NewUserService(repo)

	_, err := svc.GetByID(context.Background(), 999)
	if err != ErrUserNotFound {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestUserService_Create(t *testing.T) {
	repo := repository.NewInMemoryUserRepository([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
	})
	svc := NewUserService(repo)

	u, err := svc.Create(context.Background(), "Charlie", "charlie@test.com")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.ID != 2 {
		t.Fatalf("expected new ID=2, got %d", u.ID)
	}

	all, _ := svc.GetAll(context.Background())
	if len(all) != 2 {
		t.Fatalf("expected 2 users after create, got %d", len(all))
	}
}
