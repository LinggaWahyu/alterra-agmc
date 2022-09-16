package main

import (
	"alterra-agmc-day-4/config"
	"alterra-agmc-day-4/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
