//go:build wireinject
// +build wireinject

package rent

import "github.com/google/wire"

func InitializeRentController() *RentController {
	wire.Build(rentSet)
	return &RentController{}
}
