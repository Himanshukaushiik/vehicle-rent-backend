//go:build wireinject
// +build wireinject

package vehicles

import "github.com/google/wire"

func InitializeVehicleController() *VehicleController {
	wire.Build(VehicleSet)
	return &VehicleController{}
}
