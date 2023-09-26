package controller

import "github.com/labstack/echo/v4"

type UserController interface {
	Create(c echo.Context) error
	UpdatePassword(c echo.Context) error
	Delete(c echo.Context) error
	Login(c echo.Context) error
	FetchUSer(c echo.Context) error
}
