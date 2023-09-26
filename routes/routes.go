package routes

import (
	"fmt"
	"golang_framework_echo/controller"
	"golang_framework_echo/repository"
	"golang_framework_echo/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println(string(reqBody))
		fmt.Println(string(resBody))
	}))
	validate := validator.New()
	repositoryPegawai := repository.NewPegawaiRepository()
	servicePegawai := service.NewPegawaiService(repositoryPegawai, validate)
	controllerPegawai := controller.NewPegawaiController(servicePegawai)

	e.GET("/api/pegawai", controllerPegawai.FindAll)
	e.GET("/api/pegawai/:pegawaiId", controllerPegawai.FindById)
	e.POST("/api/pegawai", controllerPegawai.Create)
	e.PUT("/api/pegawai/:pegawaiId", controllerPegawai.Update)
	e.DELETE("/api/pegawai/:pegawaiId", controllerPegawai.Delete)

	repositoryUser := repository.NewUserRepository()
	serviceUser := service.NewUserService(repositoryUser, validate)
	controllerUser := controller.NewUserController(serviceUser)

	e.GET("/api/user/:userId", controllerUser.FetchUSer)
	e.POST("/api/login", controllerUser.Login)
	e.POST("/api/user", controllerUser.Create)
	e.PUT("/api/user/:userIf", controllerUser.Update)
	e.DELETE("/api/user/:userId", controllerUser.Delete)

	return e
}
