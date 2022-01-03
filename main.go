package main

import (
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/murugan-dev/propertymgmt/handler"
	logger2 "github.com/sirupsen/logrus"
)

var db *sqlx.DB
var log *logger2.Logger

func main() {
	app := iris.New()
	app.Use(recover.New(), logger.New(logger.DefaultConfig()))

	log = logger2.New()

	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=dev password=dev dbname=dev sslmode=disable")
	if err != nil {
		panic(err)
	}

	h := handler.NewHandler(db, log)
	party := app.Party("/api/property")
	{
		party.Post("/sell/v1", h.SellHouse)
		party.Post("/rent/v1", h.RentHouse)
		party.Get("/find/v1/{country}/{locality}/{type}", h.FindHouse)
	}

	app.Listen(":8080")
}
