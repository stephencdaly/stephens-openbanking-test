package api

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stephencdaly/stephens-openbanking-test/database"
)

type Config struct {
	DB *database.DB
}

func Start(config Config) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8080"))
}
