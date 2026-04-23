package rent

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type RentService struct {
	repo domain.RentRepository
}

func NewRentService(r domain.RentRepository) *RentService {
	return &RentService{repo: r}
}
func (s *RentService) CreateRent(r *models.Rent) error {
	return s.repo.CreateRent(r)
}

func (s *RentService) GetAllRent() ([]models.Rent, error) {
	return s.repo.GetAllRent()
}
