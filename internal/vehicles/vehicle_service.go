package vehicles

import (
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type VehicleService struct {
	repo domain.VehicleRepository
}

func NewVehicleService(r domain.VehicleRepository) *VehicleService {
	return &VehicleService{repo: r}
}

func (s *VehicleService) CreateVehicle(v *models.Vehicle) error {
	return s.repo.CreateVehicle(v)
}
func (s *VehicleService) GetVehicle() ([]models.Vehicle, error) {
	return s.repo.GetVehicles()
}

func (s *VehicleService) GetVehicleById(id string) (models.Vehicle, error) {
	return s.repo.GetVehicleByID(id)
}

func (s *VehicleService) UpdateVehicle(v *models.Vehicle) error {
	return s.repo.UpdateVehicle(v)
}
func (s *VehicleService) DeleteVehicle(v *models.Vehicle) error {
	return s.repo.DeleteVehicle(v)
}
