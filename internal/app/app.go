package app

import (
	"io"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/controller/handler"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/labstack/echo/v4"

	logger "github.com/FackOff25/GoToTeamGradGoLibs/logger"
)

func Run(configFilePath string) {
	config, err := config.GetConfig(configFilePath)

	if err != nil {
		log.Fatal(err)
	}

	configOutput := config.LogOutput
	if err := os.MkdirAll(filepath.Dir(configOutput), 0770); err != nil {
		panic(err)
	}
	logFile, err := os.OpenFile(configOutput, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Error opening log file: %s", err)
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			panic(err)
		}
	}()

	logOutput := io.MultiWriter(os.Stdout, logFile)

	logger.InitEx(logger.Options{
		Name:      config.LogAppName,
		LogLevel:  log.Level(config.LogLevel),
		LogFormat: config.LogFormat,
		Out:       logOutput,
	})

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
