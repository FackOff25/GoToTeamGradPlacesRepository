package app

import (
	"log"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller/handler"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/labstack/echo/v4"
)

func Run(configFilePath string) {
	config, err := config.GetConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	serverAddress := config.ServerAddress + ":" + config.ServerPort

	e := echo.New()

	if err := configureServer(config, e); err != nil {
		log.Fatalf("error while configuring server: %s", err)
	}

	if err := e.Start(serverAddress); err != nil {
		log.Fatalf("error while starting server: %s", err)
	}
}

func configureServer(cfg config.Config, e *echo.Echo) error {

	placesUsecase := usecase.UseCase{}
	placesController := controller.PlacesController{PlacesUsecase: placesUsecase, Config: cfg}

	e.GET("/api/v1/places/list", placesController.CreatePlacesListHandler)

	e.GET("/api/v1/places/info", placesController.CreatePlaceInfoHandler)

	e.GET("/api/v1/dummy", handler.CreateNotImplementedResponse)

	return nil
}
