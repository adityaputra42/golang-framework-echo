package main

import (
	"golang_framework_echo/db"
	"golang_framework_echo/routes"
)

func main() {
	db.InitDB()
	server := routes.Init()
	// e := InitializedServer()
	server.Logger.Fatal(server.Start(":9000"))
}
