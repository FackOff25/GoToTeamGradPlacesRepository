package domain

import (
	"github.com/google/uuid"
)

type OpeningHours struct {
	OpenNow bool `json:"open_now,omitempty"`
}

type ApiPlace struct {
	OpeningHours     OpeningHours `json:"current_opening_hours,omitempty"`
	FormattedAddress string       `json:"formatted_address,omitempty"`
	Name             string       `json:"name,omitempty"`
	PlaceId          string       `json:"place_id,omitempty"`
	Rating           float32      `json:"rating,omitempty"`
	RateCount        int          `json:"user_ratings_total,omitempty"`
	Types            []string     `json:"types,omitempty"`
}

type Place struct {
	Id        uuid.UUID
	ApiId     string
	Rating    float32
	RateCount int
}
