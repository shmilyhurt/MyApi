package repository

import (
	"MyApi/internal/model"
	"MyApi/pkg/database"
)

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

func GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := database.DB.Find(&users).Error
	return users, err
}
