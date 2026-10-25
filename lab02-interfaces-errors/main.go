package main

import (
	"errors"
	"fmt"
)

func main() {
	repository := NewMemoryUserRepository()
	service := NewUserService(repository)

	user := User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	fmt.Println("1. Registering a valid user:")
	err := service.Register(user)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User registered")
	}

	fmt.Println("\n2. Getting the registered user:")
	foundUser, err := service.GetUser(user.ID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", foundUser)
	}

	fmt.Println("\n3. Registering the same ID again:")
	err = service.Register(user)
	if err == nil {
		fmt.Println("Unexpected success")
	} else if errors.Is(err, ErrUserAlreadyExists) {
		fmt.Println("Expected error:", err)
	} else {
		fmt.Println("Unexpected error:", err)
	}

	fmt.Println("\n4. Getting a missing user:")
	missingUser, err := service.GetUser(99)
	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("Expected error:", err)
	} else if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Println("User:", missingUser)
	}

	fmt.Println("\n5. Registering an invalid user:")
	invalidUser := User{ID: 2, Name: "", Email: "wrong-email"}
	err = service.Register(invalidUser)
	if err == nil {
		fmt.Println("Unexpected success")
	} else if errors.Is(err, ErrInvalidUser) {
		fmt.Println("Expected error:", err)
	} else {
		fmt.Println("Unexpected error:", err)
	}
}
