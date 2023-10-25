package controller

import (
	"errors"
	"net/http"

	"log"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/usecase"
	"github.com/FackOff25/GoToTeamGradPlacesRepository/pkg/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

func (pc *PlacesController) formPlaceInfo(rawInfo interface{}) (domain.PlaceInfo, error) {
	googleParams := getGoogleParams()

	_, ok := rawInfo.(map[string]interface{})["result"]
	if !ok {
		return domain.PlaceInfo{}, errors.New("Bad answer")
	}

	infoMap := rawInfo.(map[string]interface{})["result"].(map[string]interface{})

	//log.Printf("%#v", infoMap[googleParams.WorkingHours])

	var name string
	nameI, ok := infoMap[googleParams.Name]
	if ok {
		name = nameI.(string)
	}

	var rating float64
	ratingI, ok := infoMap[googleParams.Rating]
	if ok {
		rating = ratingI.(float64)
	}

	var ratingCount int
	ratingCountI, ok := infoMap[googleParams.RatingCount]
	if ok {
		ratingCount = int(ratingCountI.(float64))
	}

	location := domain.ApiLocation{}
	geometry, geomOk := infoMap[googleParams.Geometry]
	if geomOk {
		geomMap := geometry.(map[string]interface{})
		locationMap := geomMap[googleParams.Location].(map[string]interface{})
		location = domain.ApiLocation{
			Lat: locationMap[googleParams.Lat].(float64),
			Lng: locationMap[googleParams.Lng].(float64),
		}
	}

	var address string
	addressI, ok := infoMap[googleParams.Address]
	if ok {
		address = addressI.(string)
	}

	var description string
	editorial_summaryI, ok := infoMap[googleParams.Summary]
	if ok {
		editorial_summary := editorial_summaryI.(map[string]interface{})
		description = editorial_summary[googleParams.Description].(string)

	}

	var workingHours []string
	openingHoursI, ok := infoMap[googleParams.WorkingHours]
	if ok {
		openingHours := openingHoursI.(map[string]interface{})
		weekdayTextI, ok := openingHours[googleParams.WorkingHoursWeekday]
		if ok {
			weekdayTextSliceI := weekdayTextI.([]interface{})
			workingHours = make([]string, len(weekdayTextSliceI))
			for i, v := range weekdayTextSliceI {
				workingHours[i] = v.(string)
			}
		}
	}

	var photos []string
	photosI, ok := infoMap[googleParams.Photos]
	if ok {
		photosSliceI := photosI.([]interface{})
		photos = make([]string, len(photosSliceI))
		for i, photoI := range photosSliceI {
			photo := photoI.(map[string]interface{})
			photos[i] = pc.Config.PlacesApiHost + "place/photo?photo_reference=" + photo["photo_reference"].(string) //now the links return 403
		}
	}

	var types []string
	typesI, ok := infoMap[googleParams.Types]
	if ok {
		typesSliceI := typesI.([]interface{})
		log.Printf("%#v", typesSliceI)
		for _, v := range typesSliceI {
			tags := getTags()
			tag, ok := tags[v.(string)]
			if ok {
				types = append(types, tag)
			}
		}
	}

	uuid, _ := uuid.NewUUID() //TODO: replace with actual uuid

	return domain.PlaceInfo{
		Id:             uuid,
		Name:           name,
		Rating:         rating,
		RatingCount:    0,
		Location:       location,
		ApiRatingCount: ratingCount,
		Description:    description,
		Address:        address,
		WorkingHours:   workingHours,
		Photos:         photos,
		Tags:           types,
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
