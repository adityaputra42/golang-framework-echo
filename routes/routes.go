package routes

import (
	"fmt"
	md "golang_framework_echo/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewValidator() *validator.Validate {
	return validator.New()
}
func Init() *echo.Echo {
	e := echo.New()
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println(string(reqBody))
		fmt.Println(string(resBody))
	}))
	controllerPegawai := InitializePegawaiController()
	controllerUser := InitializeUserController()

	e.GET("/api/pegawai", controllerPegawai.FindAll, md.IsAuthenticated)
	e.GET("/api/pegawai/:pegawaiId", controllerPegawai.FindById, md.IsAuthenticated)
	e.POST("/api/pegawai", controllerPegawai.Create, md.IsAuthenticated)
	e.PUT("/api/pegawai/:pegawaiId", controllerPegawai.Update, md.IsAuthenticated)
	e.DELETE("/api/pegawai/:pegawaiId", controllerPegawai.Delete, md.IsAuthenticated)

	e.GET("/api/user", controllerUser.FetchUSer, md.IsAuthenticated)
	e.POST("/api/login", controllerUser.Login)
	e.POST("/api/register", controllerUser.Create)
	e.PUT("/api/user", controllerUser.UpdatePassword, md.IsAuthenticated)
	e.DELETE("/api/user/:userId", controllerUser.Delete, md.IsAuthenticated)

	return e
}
