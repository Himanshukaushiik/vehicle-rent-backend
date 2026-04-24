package rent

import "github.com/google/wire"

var rentSet = wire.NewSet(
	NewRentRepository,
	NewRentService,
	NewRentController)
