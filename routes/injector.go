//go:build wireinject
// +build wireinject

package routes

import (
	"golang_framework_echo/controller"
	"golang_framework_echo/repository"
	"golang_framework_echo/service"

	"github.com/google/wire"
)

var PegawaiSet = wire.NewSet(
	repository.NewPegawaiRepository,
	service.NewPegawaiService,
	controller.NewPegawaiController,
)
var UserSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	controller.NewUserController,
)

func InitializePegawaiController() controller.PegawaiController {
	wire.Build(
		NewValidator,
		PegawaiSet,
	)
	return nil
}

func InitializeUserController() controller.UserController {
	wire.Build(
		NewValidator,
		UserSet,
	)
	return nil
}
