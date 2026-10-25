package main

import (
	"errors"
	"testing"
)

func TestUserServiceRegister(t *testing.T) {
	existingUser := User{ID: 1, Name: "Existing", Email: "existing@example.com"}

	tests := []struct {
		name         string
		existingUser *User
		user         User
		wantErr      error
	}{
		{
			name:    "valid user",
			user:    User{ID: 1, Name: "Alice", Email: "alice@example.com"},
			wantErr: nil,
		},
		{
			name:    "ID is not positive",
			user:    User{ID: 0, Name: "Alice", Email: "alice@example.com"},
			wantErr: ErrInvalidUser,
		},
		{
			name:    "name is empty",
			user:    User{ID: 1, Name: "  ", Email: "alice@example.com"},
			wantErr: ErrInvalidUser,
		},
		{
			name:    "email is invalid",
			user:    User{ID: 1, Name: "Alice", Email: "wrong-email"},
			wantErr: ErrInvalidUser,
		},
		{
			name:         "user already exists",
			existingUser: &existingUser,
			user:         User{ID: 1, Name: "Alice", Email: "alice@example.com"},
			wantErr:      ErrUserAlreadyExists,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repository := NewMemoryUserRepository()
			service := NewUserService(repository)

			if test.existingUser != nil {
				if err := repository.Save(*test.existingUser); err != nil {
					t.Fatalf("prepare repository: %v", err)
				}
			}

			err := service.Register(test.user)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("Register() error = %v, want %v", err, test.wantErr)
			}

			if test.wantErr == nil {
				savedUser, err := service.GetUser(test.user.ID)
				if err != nil {
					t.Fatalf("GetUser() after Register(): %v", err)
				}
				if savedUser != test.user {
					t.Fatalf("saved user = %v, want %v", savedUser, test.user)
				}
			}
		})
	}
}

func TestUserServiceGetUser(t *testing.T) {
	storedUser := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	tests := []struct {
		name    string
		id      int
		want    User
		wantErr error
	}{
		{
			name:    "existing user",
			id:      1,
			want:    storedUser,
			wantErr: nil,
		},
		{
			name:    "missing user",
			id:      99,
			want:    User{},
			wantErr: ErrUserNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repository := NewMemoryUserRepository()
			if err := repository.Save(storedUser); err != nil {
				t.Fatalf("prepare repository: %v", err)
			}

			service := NewUserService(repository)
			got, err := service.GetUser(test.id)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("GetUser() error = %v, want %v", err, test.wantErr)
			}
			if got != test.want {
				t.Fatalf("GetUser() = %v, want %v", got, test.want)
			}
		})
	}
}
