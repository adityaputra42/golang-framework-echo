package controller

import (
	"github.com/labstack/echo/v4"
)

type PegawaiController interface {
	Create(c *echo.Context)
	Update(c *echo.Context)
	Delete(c *echo.Context)
	FindById(c *echo.Context)
	FindAll(c *echo.Context)
}
