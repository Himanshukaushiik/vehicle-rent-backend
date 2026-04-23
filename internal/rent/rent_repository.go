package rent

import (
	"awesomeProject1/config"
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type rentRepo struct{}

func NewRentRepository() domain.RentRepository {
	return &rentRepo{}
}

func (r *rentRepo) CreateRent(rent *models.Rent) error {
	return config.DB.Create(rent).Error

}

func (r *rentRepo) GetAllRent() ([]models.Rent, error) {
	var rents []models.Rent
	err := config.DB.Find(&rents).Error
	return rents, err
}
