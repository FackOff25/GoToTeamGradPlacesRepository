package controller

import (
	"net/http"

	"log"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/FackOff25/GoToTeamGradGoLibs/googleApi"
)

type GoogleParams struct {
	Name                string
	Rating              string
	RatingCount         string
	Description         string
	Address             string
	Geometry            string
	Summary             string
	Location            string
	WorkingHours        string
	WorkingHoursWeekday string
	Photos              string
	Lat                 string
	Lng                 string
	Photo               string
	Types               string
}

func getGoogleParams() GoogleParams {
	return GoogleParams{
		Name:                "name",
		Rating:              "rating",
		RatingCount:         "user_ratings_total",
		Summary:             "editorial_summary",
		Description:         "overview",
		Address:             "formatted_address",
		Geometry:            "geometry",
		Location:            "location",
		WorkingHours:        "opening_hours",
		WorkingHoursWeekday: "weekday_text",
		Photos:              "photos",
		Lat:                 "lat",
		Lng:                 "lng",
		Photo:               "photo",
		Types:               "types",
	}
}

func getTags() map[string]string {
	return map[string]string{
		"tourist_attraction": "Для туристов",
		"museum":             "Музей",
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

func (pc *PlacesController) formPlaceInfo(result googleApi.Place) (domain.PlaceInfo, error) {
	uuid, _ := uuid.NewUUID() //TODO: replace with actual uuid

	location := domain.ApiLocation{
		Lat: result.Geometry.Location.Lat,
		Lng: result.Geometry.Location.Lng,
	}

	var photos []string
	for _, photoStruct := range result.Photos {
		reference := photoStruct.Reference
		photos = append(photos, pc.Config.PlacesApiHost + "place/photo?photo_reference=" + reference)  //now the links return 403
	}

	return domain.PlaceInfo{
		Id:             uuid,
		Name:           result.Name,
		Rating:         result.Rating,
		RatingCount:    0,
		Location:       location,
		ApiRatingCount: result.RatingCount,
		Description:    result.Summary.Overview,
		Address:        result.Address,
		WorkingHours:   result.OpeningHours.WeekTimetable,
		Photos:         photos,
		Tags:           result.Types,
	}, nil
}
func (pc *PlacesController) CreatePlaceInfoHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	if !c.QueryParams().Has("place_id") {
		return echo.ErrBadRequest
	}

	placeId := c.QueryParam("place_id")

	googleParams := getGoogleParams()

	fields := []string{
		googleParams.Name,
		googleParams.Rating,
		googleParams.RatingCount,
		googleParams.Summary,
		googleParams.Address,
		googleParams.Geometry,
		googleParams.WorkingHours,
		googleParams.Photo,
		googleParams.Types,
	}

	placeInterface, err := pc.PlacesUsecase.GetInfoOnPlace(pc.Config, placeId, fields)

	if err != nil {
		log.Print(err)
	}

	place, err := pc.formPlaceInfo(placeInterface)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, place)
}
