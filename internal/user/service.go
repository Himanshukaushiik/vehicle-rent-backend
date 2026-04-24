package user

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) domain.UserService {
	return &userService{repo: r}
}

func (s *userService) CreateUser(u *models.User) error {
	return s.repo.CreateUser(u)
}

// GET ALL
func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// DELETE
func (s *userService) DeleteUser(u *models.User) error {
	return s.repo.DeleteUser(u)
}

// UPDATE
func (s *userService) UpdateUser(u *models.User) error {
	return s.repo.UpdateUser(u)
}
