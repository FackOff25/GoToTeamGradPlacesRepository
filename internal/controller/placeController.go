package controller

import (
	"net/http"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PlacesController struct {
	Config        config.Config
	PlacesUsecase usecase.UseCase
}

func (pc *PlacesController) CreatePlacesListHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	location := ""

	uuid, _ := uuid.NewUUID()

	places, _ := pc.PlacesUsecase.GetNearbyPlaces(uuid, location)

	return c.JSON(http.StatusOK, places)
}

func (pc *PlacesController) CreatePlaceInfoHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	placeId := c.QueryParam("place_id")

	fields := []string{
		"name",
		"rating",
		"user_ratings_total",
	}

	places, err := pc.PlacesUsecase.GetInfoOnPlace(pc.Config, placeId, fields)

	if err != nil {
		log.Error(err)
	}

	return c.JSON(http.StatusOK, places)
}
