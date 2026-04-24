package user

import (
	"awesomeProject1/config"
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type userRepo struct{}

func NewUserRepository() domain.UserRepository {
	return &userRepo{}
}

// Create
func (r *userRepo) CreateUser(u *models.User) error {
	return config.DB.Create(u).Error
}

// Get all
func (r *userRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User

	err := config.DB.Find(&users).Error
	return users, err
}

// Delete
func (r *userRepo) DeleteUser(u *models.User) error {
	return config.DB.Delete(u).Error
}

// update
func (r *userRepo) UpdateUser(u *models.User) error {
	return config.DB.Save(u).Error
}
