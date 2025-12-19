package service

import (
	"go-cloud-api/internal/model"
	"testing"
)

func TestUserService_GetAll(t *testing.T) {
	svc := NewUserServiceWithSeed([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	})
	users, err := svc.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}
func TestUserService_GetByID_Found(t *testing.T) {
	svc := NewUserServiceWithSeed([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
	})
	u, err := svc.GetByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.ID != 1 {
		t.Fatalf("expected user ID 1, got %d", u.ID)
	}
}
func TestUserService_GetByID_NotFound(t *testing.T) {
	svc := NewUserServiceWithSeed([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
	})
	_, err := svc.GetByID(999)
	if err != ErrUserNotFound {
		t.Fatalf("expected error, got %v", err)
	}
}
func TestUserService_Create(t *testing.T) {
	svc := NewUserServiceWithSeed([]model.User{
		{ID: 1, Name: "Alice", Email: "alice@test.com"},
	})
	u, err := svc.Create("Charlie", "charlie@example.com")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if u.ID != 2 {
		t.Fatalf("expected user ID 2, got %d", u.ID)
	}
	all, _ := svc.GetAll()
	if len(all) != 2 {
		t.Fatalf("expected 2 users, got %d", len(all))
	}
}
