package rent

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type rentService struct {
	repo domain.RentRepository
}

func NewRentService(r domain.RentRepository) domain.RentService {
	return &rentService{repo: r}
}

func (s *rentService) CreateRent(r *models.Rent) error {
	return s.repo.CreateRent(r)
}

func (s *rentService) GetAllRent() ([]models.Rent, error) {
	return s.repo.GetAllRent()
}
