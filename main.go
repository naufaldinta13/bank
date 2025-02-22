package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naufaldinta13/bank/config"
	"github.com/naufaldinta13/bank/handler"
	"github.com/naufaldinta13/bank/utils"
)

func initPostgresqlConnection() {
	c := &config.PostgresConfig{
		Server:   os.Getenv("POSTGRES_SERVER"),
		Username: os.Getenv("POSTGRES_USERNAME"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
	}

	if e := config.NewConnection(c); e != nil {
		slog.Error(fmt.Sprintf("Postgresql %s", e.Error()))
	}
}

func initRestServer() {

	e := echo.New()
	e.Validator = utils.NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.RegisterHandler(e)

	e.Logger.Fatal(e.Start(os.Getenv("REST_SERVER")))
}

func main() {
	godotenv.Load()

	initPostgresqlConnection()
	initRestServer()

}
