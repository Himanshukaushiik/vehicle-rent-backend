package domain

import "awesomeProject1/models"

type RentRepository interface {
	CreateRent(r *models.Rent) error
	GetAllRent() ([]models.Rent, error)
}
