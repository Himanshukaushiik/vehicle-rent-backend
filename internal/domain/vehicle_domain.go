package domain

import "awesomeProject1/models"

type VehicleRepository interface {
	CreateVehicle(v *models.Vehicle) error
	GetVehicles() ([]models.Vehicle, error)
	GetVehicleByID(id string) (models.Vehicle, error)
	UpdateVehicle(v *models.Vehicle) error
	DeleteVehicle(v *models.Vehicle) error
}

type VehicleService interface {
	CreateVehicle(vehicle *models.Vehicle) error
	GetVehicles() ([]models.Vehicle, error)
	GetVehicleByID(id string) (models.Vehicle, error)
	DeleteVehicle(vehicle *models.Vehicle) error
}

//Interface = system ka contract
