//go:build wireinject
// +build wireinject

package main

import (
	rent "awesomeProject1/internal/rent"
	user "awesomeProject1/internal/user"
	vehicles "awesomeProject1/internal/vehicles"

	"github.com/google/wire"
)

func InitializeVehicleController() *vehicles.VehicleController {
	wire.Build(
		vehicles.NewVehicleRepository,
		vehicles.NewVehicleService,
		vehicles.NewVehicleController,
	)
	return &vehicles.VehicleController{}
}

func InitializeRentController() *rent.RentController {
	wire.Build(
		rent.NewRentRepository,
		rent.NewRentService,
		rent.NewRentController,
	)
	return &rent.RentController{}
}

func InitializeUserController() *user.UserController {
	wire.Build(
		user.NewUserRepository,
		user.NewUserService,
		user.NewUserController,
	)
	return &user.UserController{}
}
