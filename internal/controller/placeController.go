package controller

import (
	"net/http"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PlacesController struct {
	PlacesUsecase usecase.UseCase
}

func (pc *PlacesController) CreatePlacesListHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	location := ""

	uuid, _ := uuid.NewUUID()

	places, _ := pc.PlacesUsecase.GetNearbyPlaces(uuid, location)

	return c.JSON(http.StatusOK, places)
}
