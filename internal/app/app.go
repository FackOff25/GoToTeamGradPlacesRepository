package app

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Run() {

	e := echo.New()

	serverAddress := "0.0.0.0:8000"

	if err := e.Start(serverAddress); err != nil {
		log.Fatalf("error while starting server", err)
	}
}

func configureServer(e *echo.Echo) error {

	//e.GET("/api/v1/session/create", placesController.GetPlaces)

	return nil
}
