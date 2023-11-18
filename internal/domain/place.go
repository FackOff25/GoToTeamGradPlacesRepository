package domain

import (
	"github.com/google/uuid"
)

const (
	ReactionLike    = "like"
	ReactionVisited = "visited"
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

type PlaceInfo struct {
	Id          string      `json:"id,omitempty"`
	Cover       string      `json:"cover,omitempty"`
	Rating      float64     `json:"rating,omitempty"`
	RatingCount int64       `json:"rating_count,omitempty"`
	Name        string      `json:"name,omitempty"`
	Location    ApiLocation `json:"location,omitempty"`
	PlaceId     string      `json:"place_id,omitempty"`
	//ApiRatingCount int64         `json:"user_ratings_total,omitempty"`
	Description  string   `json:"description,omitempty"`
	Photos       []string `json:"photos,omitempty"`
	Address      string   `json:"address,omitempty"`
	WorkingHours []string `json:"workingHours,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Reactions    []string `json:"reactions,omitempty"`
}

type Place struct {
	Id        uuid.UUID
	PlaceId   string
	Rating    float32
	RateCount int
}
