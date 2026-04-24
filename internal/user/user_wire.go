//go:build wireinject
// +build wireinject

package user

import "github.com/google/wire"

func InitializeUserController() *UserController {
	wire.Build(UserSet)
	return &UserController{}
}
