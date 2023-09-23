package routes

import (
	"golang_framework_echo/controller"
	"golang_framework_echo/repository"
	"golang_framework_echo/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	validate := validator.New()
	repository := repository.NewPegawaiRepository()
	service := service.NewPegawaiService(repository, validate)
	controller := controller.NewPegawaiController(service)

	e.GET("/api/pegawai", controller.FindAll)
	e.GET("/api/pegawai/:pegawaiId", controller.FindById)
	e.POST("/api/pegawai", controller.Create)
	e.PUT("/api/pegawai/:pegawaiId", controller.Update)
	e.DELETE("/api/pegawai/:pegawaiId", controller.Delete)

	return e
}
