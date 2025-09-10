package service

import (
	"MyApi/internal/model"
	"MyApi/internal/repository"
)

func CreateUser(name, pwd string) error {
	user := &model.User{Username: name, Password: pwd}
	return repository.CreateUser(user)
}

func ListUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}
