package controller

import (
	"bytes"
	"encoding/json"
	"html"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
	log "github.com/sirupsen/logrus"

	"github.com/FackOff25/GoToTeamGradGoLibs/categories"
	"github.com/FackOff25/GoToTeamGradGoLibs/googleApi"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const uuidHeader = "X-Uuid"

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

func (pc *PlacesController) formPlaceInfo(result googleApi.Place, placeId, userId string) (domain.PlaceInfo, error) {
	location := domain.ApiLocation{
		Lat: result.Geometry.Location.Lat,
		Lng: result.Geometry.Location.Lng,
	}

	var photos []string
	for _, photoStruct := range result.Photos {
		reference := photoStruct.Reference
		url := pc.Config.PlacesApiHost + "place/photo?maxwidth=" + strconv.FormatInt(photoStruct.Width, 10) + "&photo_reference=" + reference
		photos = append(photos, url) //now the links return 403
	}

	var tags []string
	tagsMap := categories.GetCategoriesMap()
	for _, v := range result.Types {
		tag, ok := tagsMap[v]
		if ok {
			tags = append(tags, tag)
		}
	}

	var uuid string
	var reactions []string
	uuid, reactions, err := pc.PlacesUsecase.GetUserReaction(userId, placeId)
	if err != nil && err != pgx.ErrNoRows {
		log.Error("getting reaction from db error: ", err)
	}

	return domain.PlaceInfo{
		Id:          uuid,
		Name:        result.Name,
		Rating:      result.Rating,
		RatingCount: result.RatingCount,
		Location:    location,
		//ApiRatingCount: result.RatingCount,
		Description:  result.Summary.Overview,
		Address:      result.Address,
		WorkingHours: result.OpeningHours.WeekTimetable,
		Photos:       photos,
		Tags:         tags,
		Reactions:    reactions,
	}, nil
}

func (pc *PlacesController) CreatePlaceInfoHandler(c echo.Context) error {
	defer c.Request().Body.Close()

	id := html.EscapeString(c.Request().Header.Get("X-UUID"))

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
		log.Errorf("Error in GetInfoOnPlace: %s", err)
		if err.Error() == googleApi.STATUS_NOT_FOUND || err.Error() == googleApi.STATUS_INVALID_REQUEST {
			return echo.ErrNotFound
		}

		return echo.NewHTTPError(500, "INTERNAL")
	}

	place, err := pc.formPlaceInfo(placeInterface, placeId, id)
	if err != nil {
		log.Errorf("Error getting info from googleApi: %s", err)
		return echo.ErrNotFound
	}

	place.PlaceId = placeId
	resBodyBytes := new(bytes.Buffer)
	encoder := json.NewEncoder(resBodyBytes)
	encoder.SetEscapeHTML(false)
	encoder.Encode(place)

	return c.JSONBlob(http.StatusOK, resBodyBytes.Bytes())
}
