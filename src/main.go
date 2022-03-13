package main

import (
	"github.com/echo_CRUD/src/database"
	"github.com/echo_CRUD/src/routers"
)

func main() {
	e := routers.Init()
	err := database.ConnectDB()
	if err != nil {
		return 
	}
	e.Logger.Fatal(e.Start(":5400"))
}
