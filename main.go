package main

import (
	"clinic/config"
	"clinic/routes"
)

func main() {
	config.InitDB()
	e:= routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}