package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUser       = errors.New("invalid user")
)

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Save(user User) error
	FindByID(id int) (User, error)
}

type MemoryUserRepository struct {
	users map[int]User
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users: make(map[int]User),
	}
}

func (repository *MemoryUserRepository) Save(user User) error {
	if _, exists := repository.users[user.ID]; exists {
		return fmt.Errorf("save user %d: %w", user.ID, ErrUserAlreadyExists)
	}

	repository.users[user.ID] = user
	return nil
}

func (repository *MemoryUserRepository) FindByID(id int) (User, error) {
	user, exists := repository.users[id]
	if !exists {
		return User{}, fmt.Errorf("find user %d: %w", id, ErrUserNotFound)
	}

	return user, nil
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (service *UserService) Register(user User) error {
	if user.ID <= 0 {
		return fmt.Errorf("%w: ID must be positive", ErrInvalidUser)
	}
	if strings.TrimSpace(user.Name) == "" {
		return fmt.Errorf("%w: name must not be empty", ErrInvalidUser)
	}
	if !strings.Contains(user.Email, "@") {
		return fmt.Errorf("%w: email must contain @", ErrInvalidUser)
	}

	if err := service.repository.Save(user); err != nil {
		return fmt.Errorf("register user: %w", err)
	}

	return nil
}

func (service *UserService) GetUser(id int) (User, error) {
	user, err := service.repository.FindByID(id)
	if err != nil {
		return User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
