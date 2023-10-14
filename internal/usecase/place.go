package usecase

import (
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/google/uuid"
)

func (uc *UseCase) GetNearbyPlaces(id uuid.UUID, location string) ([]domain.ApiPlace, error) {
	// Google Maps API call
	// nearbyPlaces, err := something.GetNearbyPlaces(location)
	nearbyPlaces := make([]domain.ApiPlace, 0)

	nearbyPlaces = append(nearbyPlaces,
		domain.ApiPlace{
			Id:          uuid.MustParse("7d249c01-6844-4a99-b4ed-479bf2e5a639"),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMImrKNC6PKPQrUEI4GvXoqyhF05Pbk7CkYNXZ1=w408-h255-k-no",
			RatingCount: 54861,
			Rating:      4.7,
			Name:        "Московский Кремль",
			PlaceId:     "ChIJgUbEo8cfqokR5lP9_Wh_DaM",
			Lat:         -33.866489,
			Lng:         1.123,
		},
		domain.ApiPlace{
			Id:          uuid.MustParse("7d249c01-6844-4a99-b4ed-479bf2e5a639"),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMImrKNC6PKPQrUEI4GvXoqyhF05Pbk7CkYNXZ1=w408-h255-k-no",
			RatingCount: 54861,
			Rating:      4.7,
			Name:        "Московский Кремль",
			PlaceId:     "ChIJgUbEo8cfqokR5lP9_Wh_DaM",
			Lat:         -33.866489,
			Lng:         1.123,
		}, domain.ApiPlace{
			Id:          uuid.MustParse("7d249c01-6844-4a99-b4ed-479bf2e5a639"),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMImrKNC6PKPQrUEI4GvXoqyhF05Pbk7CkYNXZ1=w408-h255-k-no",
			RatingCount: 54861,
			Rating:      4.7,
			Name:        "Московский Кремль",
			PlaceId:     "ChIJgUbEo8cfqokR5lP9_Wh_DaM",
			Lat:         -33.866489,
			Lng:         1.123,
		},
	)

	return nearbyPlaces, nil
}
