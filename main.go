package main

import (
	"log"
	"net/http"

	"library/commons/dbconn"
	"library/controller"
	"library/postgres"
	"library/router"
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
	// initialize routes
	router.InitializeRoutes(app)

	// run server
	log.Fatal(http.ListenAndServe(":3000", nil))
}
