package controller

import (
	"net/http"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
)

type GoogleParams struct{
	Name  string
	Rating  string
	RatingCount  string
	Description  string
	Address  string
	Geometry  string
	Location  string
	WorkingHours  string
	Photos  string
	Lat string
	Lng string
}


func getGoogleParams() GoogleParams {
	return GoogleParams{
		Name: "name",
		Rating: "rating",
		RatingCount: "user_ratings_total",
		Description: "editorial_summary",
		Address: "formatted_address",
		Geometry: "geometry",
		Location: "location",
		WorkingHours: "opening_hours",
		Photos: "photos",
		Lat: "lat",
		Lng: "lng",
	}
}

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


func formPlaceInfo(rawInfo interface{}) domain.PlaceInfo {
	googleParams := getGoogleParams()

	infoMap := rawInfo.(map[string]interface{})

	location := infoMap[googleParams.Geometry].(map[string]interface{})[googleParams.Location]

	return domain.PlaceInfo{
		Id: uuid.NewUUID(),
		Name: infoMap[googleParams.Name].(string),
		Rating: infoMap[googleParams.Rating].,
		//Cover: "string",
		RatingCount: 0,
		Location: domain.ApiLocation{
			Lat: location[googleParams.Lat],
			Lng: location[googleParams.Lng],
		},
		ApiRatingCount: infoMap[googleParams.RatingCount],
		Description: infoMap[googleParams.Description],
		Photos: infoMap[googleParams.Photos],
    	Address: infoMap[getGoogleParams.Address],
	}
}

func (pc *PlacesController) CreatePlaceInfoHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	placeId := c.QueryParam("place_id")

	googleParams := getGoogleParams()

	fields := []string{
		googleParams.Name,
		googleParams.Rating,
		googleParams.RatingCount,
		googleParams.Description,
		googleParams.Address,
		googleParams.Geometry,
		googleParams.WorkingHours,
		googleParams.Photos,
	}

	placeInterface, err := pc.PlacesUsecase.GetInfoOnPlace(pc.Config, placeId, fields)

	if err != nil {
		log.Fatal(err)
	}

	place := formPlaceInfo(placeInterface)

	return c.JSON(http.StatusOK, place)
}
