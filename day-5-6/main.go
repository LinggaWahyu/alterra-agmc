package main

import (
	"alterra-agmc-day-5-6/database"
	"alterra-agmc-day-5-6/database/migration"
	"alterra-agmc-day-5-6/internal/factory"
	"alterra-agmc-day-5-6/internal/http"
	"alterra-agmc-day-5-6/internal/middleware"
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// load env configuration
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	database.CreateConnection()

	var m string // for check migration

	flag.StringVar(
		&m,
		"migrate",
		"run",
		`this argument for check if user want to migrate table, rollback table, or status migration

to use this flag:
	use -migrate=migrate for migrate table
	use -migrate=rollback for rollback table
	use -migrate=status for get status migration`,
	)
	flag.Parse()

	if m == "migrate" {
		migration.Migrate()
		return
	} else if m == "rollback" {
		migration.Rollback()
		return
	} else if m == "status" {
		migration.Status()
		return
	}

	f := factory.NewFactory()
	e := echo.New()
	middleware.Init(e)
	http.NewHttp(e, f)

	port := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
