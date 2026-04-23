package domain

import "awesomeProject1/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]models.User, error)
	DeleteUser(user *models.User) error
	UpdateUser(user *models.User) error
}
