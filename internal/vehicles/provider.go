package vehicles

import "github.com/google/wire"

var VehicleSet = wire.NewSet(
	NewVehicleRepository,
	NewVehicleService,
	NewVehicleController)
