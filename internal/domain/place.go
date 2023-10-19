package domain

import (
	"github.com/google/uuid"
)

type ApiLocation struct {
	Lat float64 `json:"lat,omitempty"`
	Lng float64 `json:"lng,omitempty"`
}

type ApiPlace struct {
	Id             uuid.UUID   `json:"id,omitempty"`
	Cover          string      `json:"cover,omitempty"`
	Rating         float32     `json:"rating,omitempty"`
	RatingCount    int         `json:"rating_count,omitempty"`
	Name           string      `json:"name,omitempty"`
	Location       ApiLocation `json:"location,omitempty"`
	PlaceId        string      `json:"place_id,omitempty"`
	ApiRatingCount int         `json:"user_ratings_total,omitempty"`
	// Types          []string  `json:"types,omitempty"`
}

type Place struct {
	Id        uuid.UUID
	PlaceId   string
	Rating    float32
	RateCount int
}
