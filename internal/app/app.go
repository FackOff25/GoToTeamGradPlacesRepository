package app

import (
	"log"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller/handler"
	"github.com/labstack/echo/v4"
)

func Run() {

	e := echo.New()

	serverAddress := "0.0.0.0:8000"

	if err := configureServer(e); err != nil {
		log.Fatalf("error while configuring server", err)
	}

	if err := e.Start(serverAddress); err != nil {
		log.Fatalf("error while starting server", err)
	}
}

func configureServer(e *echo.Echo) error {

	e.GET("/api/v1/places/list", handler.CreateNotImplementedResponse)

	return nil
}
