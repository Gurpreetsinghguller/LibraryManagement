package main

import (
	"library/controller"
	"library/db/dbconn"
	"library/postgres"
	"library/router"

	"github.com/labstack/echo/v4"
)

var (
	conn   *dbconn.Postgres
	models *postgres.Models
	app    *controller.Application
)

func init() {
	var err error
	// Load Env file

	// Connect to DB
	conn, err = dbconn.NewPostgres()
	if err != nil {
		panic(err)
	}

	// load Models
	models = postgres.NewModels(conn)

	// initialize app object
	app = controller.New(conn, models)
}

func main() {
	e := echo.New()
	// initialize routes
	router.InitializeRoutes(app, e)

	// run server
	e.Logger.Fatal(e.Start(":3000"))
}
