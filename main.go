package main

import (
	"golang_framework_echo/db"
	"golang_framework_echo/routes"
)

func main() {
	db.InitDB()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":9000"))
}
