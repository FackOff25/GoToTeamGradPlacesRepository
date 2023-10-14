package app

import (
	"log"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller/handler"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/utils/configReader"
	"github.com/labstack/echo/v4"
)

var serverAddress = "0.0.0.0"

func Run(configFilePath string) {
	config, err := configReader.NewConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	serverAddress += ":" + config.ServerPort

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
