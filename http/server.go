package http

import (
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stephencdaly/stephens-openbanking-test/database"
	"github.com/stephencdaly/stephens-openbanking-test/internal/truelayer"
	"github.com/stephencdaly/stephens-openbanking-test/http/api"
)

type Config struct {
	DB *database.DB
}

func Start(config Config) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	truelayerToken := truelayer.GeneratePaymentToken()
	log.Printf("Got token, expires in %d", truelayerToken.ExpiresIn)

	e.POST("/v1/api/payments", api.CreatePaymentHandler(config.DB))

	e.Logger.Fatal(e.Start(":8080"))
}
