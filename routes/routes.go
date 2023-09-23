package routes

import (
	"golang_framework_echo/controller"
	"golang_framework_echo/repository"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	repository := repository.NewPegawaiRepository()
	controller := controller.NewPegawaiController(repository)

	e.GET("/api/pegawai", controller.FindAll)
	e.GET("/api/pegawai/:pegawaiId", controller.FindById)
	e.POST("/api/pegawai", controller.Create)
	e.PUT("/api/pegawai/:pegawaiId", controller.Update)
	e.DELETE("/api/pegawai/:pegawaiId", controller.Delete)

	return e
}
