package vehicles

import (
	"awesomeProject1/config"
	"awesomeProject1/internal/domain"
	"awesomeProject1/models"
)

type vehicleRepo struct{}

func NewVehicleRepository() domain.VehicleRepository {
	return &vehicleRepo{}
}

// Create
func (r *vehicleRepo) CreateVehicle(v *models.Vehicle) error {
	return config.DB.Create(v).Error
}

//GetALL

func (r *vehicleRepo) GetVehicles() ([]models.Vehicle, error) {
	var vehicles []models.Vehicle
	err := config.DB.Find(&vehicles).Error
	return vehicles, err
}

//GETBYID

func (r *vehicleRepo) GetVehicleByID(id string) (models.Vehicle, error) {
	var vehicle models.Vehicle
	err := config.DB.First(&vehicle, id).Error
	return vehicle, err
}

//Update

func (r *vehicleRepo) UpdateVehicle(v *models.Vehicle) error {
	return config.DB.Save(v).Error
}

//Delete

func (r *vehicleRepo) DeleteVehicle(v *models.Vehicle) error {
	return config.DB.Delete(v).Error
}
