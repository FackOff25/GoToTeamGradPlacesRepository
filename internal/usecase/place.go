package usecase

import (
	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/domain"
	"github.com/google/uuid"
)


func (uc *UseCase) GetNearbyPlaces(id uuid.UUID, location string) ([]domain.Place, error) {
	// Google Maps API call
	// nearbyPlaces, err := something.GetNearbyPlaces(location)

	userPlaces, err := uc.repo.GetUserPlaces(id)
	if err != nil {
		return nil, err
	}
	return userPlaces, nil
}