package main

import (
	"log"
	"net/http"

	"booklibrary/commons/utils"
	"booklibrary/controller"
	"booklibrary/postgres"
	"booklibrary/router"
)

var (
	conn   *utils.Postgres
	models *postgres.Models
	app    *controller.Application
)

func init() {
	var err error
	// Load Env

	// Connect to DB
	conn, err = utils.NewPostgres()
	if err != nil {
		panic(err)
	}

	// load Models
	models = postgres.NewModels(conn)

	// initialize app object
	app = controller.New(conn, models)
}

func main() {
	router.InitializeRoutes(app)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
