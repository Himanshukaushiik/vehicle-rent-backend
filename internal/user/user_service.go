package user

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(r domain.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(u *models.User) error {
	return s.repo.CreateUser(u)
}

// GET ALL
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

// DELETE
func (s *UserService) DeleteUser(u *models.User) error {
	return s.repo.DeleteUser(u)
}

// UPDATE
func (s *UserService) UpdateUser(u *models.User) error {
	return s.repo.UpdateUser(u)
}
