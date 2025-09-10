package service

import (
	"MyApi/internal/model"
	"MyApi/internal/repository"
)

func CreateUser(name, email string) error {
	user := &model.User{Name: name, Email: email}
	return repository.CreateUser(user)
}

func ListUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
