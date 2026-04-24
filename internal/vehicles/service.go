package vehicles

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type vehicleService struct {
	repo domain.VehicleRepository
}

func NewVehicleService(r domain.VehicleRepository) domain.VehicleService {
	return &vehicleService{repo: r}
}

func (s *vehicleService) CreateVehicle(v *models.Vehicle) error {
	return s.repo.CreateVehicle(v)
}
func (s *vehicleService) GetVehicles() ([]models.Vehicle, error) {
	return s.repo.GetVehicles()
}

func (s *vehicleService) GetVehicleByID(id string) (models.Vehicle, error) {
	return s.repo.GetVehicleByID(id)
}

func (s *vehicleService) DeleteVehicle(v *models.Vehicle) error {
	return s.repo.DeleteVehicle(v)
}
