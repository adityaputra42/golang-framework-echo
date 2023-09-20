package main

import (
	"golang_framework_echo/app"

	"github.com/labstack/echo/v4"
)

func main() {
	r := echo.New()
	// response method string
	app.NewRouter(r)
	r.Start(":9000")
}
