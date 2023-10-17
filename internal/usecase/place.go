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
		// 1
		domain.ApiPlace{
			Id:          uuid.MustParse("7d249c01-6844-4a99-b4ed-479bf2e5a639"),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMImrKNC6PKPQrUEI4GvXoqyhF05Pbk7CkYNXZ1=w408-h255-k-no",
			RatingCount: 54861,
			Rating:      4.7,
			Name:        "Московский Кремль",
			PlaceId:     "AF1QipMImrKNC6PKPQrUEI4GvXoqyhF05Pbk7CkYNXZ1",
			Lat:         55.7520233,
			Lng:         37.6174994,
		},
		// 2
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipOHIILw9yVALfLKF4e2sREloOqS3WfPxDxGykRE=w426-h240-k-no",
			RatingCount: 168180,
			Rating:      4.8,
			Name:        "Красная площадь",
			PlaceId:     "AF1QipOHIILw9yVALfLKF4e2sREloOqS3WfPxDxGykRE",
			Lat:         55.753544,
			Lng:         37.621202,
		}, 
		// 3
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMLko3fPi5q_F3w5UbX9-JStBqwB8HaCkLNUkDZ=w408-h544-k-no",
			RatingCount: 23905,
			Rating:      4.7,
			Name:        "Храм Христа Спасителя",
			PlaceId:     "AF1QipMLko3fPi5q_F3w5UbX9",
			Lat:         55.7446368,
			Lng:         37.6054959 ,
		},
		// 4
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipMQhmZY2zME-WBwWCPj9WwZhVml3-6LdKjuCz0q=w408-h272-k-no",
			RatingCount: 244,
			Rating:      4.6,
			Name:        "Старый Арбат",
			PlaceId:     "AF1QipMQhmZY2zME",
			Lat:         55.7500877,
			Lng:         37.5936813,
		},
		// 5
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipPW4xJKa8dLJKGV7d9aBf9pQ6MYgyQLQy5CBAO5=w408-h541-k-no",
			RatingCount: 77646,
			Rating:      4.5,
			Name:        "Парк Зарядье",
			PlaceId:     "AF1QipPW4xJKa8dLJKGV7d9aBf9pQ6MYgyQLQy5CBAO5",
			Lat:         55.751188,
			Lng:         37.627939,
		},
		// 6
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipNo74P1vCnj4cSOD_wRX5wwqgnvw7S29sZ0ylk=w408-h272-k-no",
			RatingCount: 3972,
			Rating:      4.7,
			Name:        "Москва-Сити",
			PlaceId:     "AF1QipNo74P1vCnj4cSODwRX5wwqgnvw7S29sZ0ylk",
			Lat:         55.749451,
			Lng:         37.542824,
		},
		// 7
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipOH6j96KF6yl8YgcmTVuIY98zdERHnHm_O2Hxxm=w408-h306-k-no",
			RatingCount: 35721,
			Rating:      4.7,
			Name:        "Воробьевы горы",
			PlaceId:     "AF1QipOH6j96KF6yl8YgcmTVuIY98zdERHnHm_O2Hxxm",
			Lat:         55.711422,
			Lng:         37.544391,
		},
		// 8
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipPqUNGOF9ej8E8s2NaHTueeqC3dzzrZwNN3UGjT=w408-h272-k-no",
			RatingCount: 130338,
			Rating:      4.7,
			Name:        "ВДНХ",
			PlaceId:     "AF1QipPqUNGOF9ej8E8s2NaHTueeqC3dzzrZwNN3UGjT",
			Lat:         55.826685,
			Lng:         37.638764,
		},
		// 9
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipPDUfGO4t-NLBt7t-kJE1df0i2UAZbmGMg7Fyyh=w408-h270-k-no",
			RatingCount: 30457,
			Rating:      4.8,
			Name:        "Третьяковская галерея",
			PlaceId:     "kJE1df0i2UAZbmGMg7Fyyh",
			Lat:         55.741333, 
			Lng:         37.620555,
		},
		// 10
		domain.ApiPlace{
			Id:          uuid.New(),
			Cover:       "https://lh5.googleusercontent.com/p/AF1QipNAytHSPl5oAcFjI2921AdQnmgULzTevdpPz0NJ=w408-h302-k-no",
			RatingCount: 72971,
			Rating:      4.4,
			Name:        "Московский зоопарк",
			PlaceId:     "AF1QipNAytHSPl5oAcFjI2921AdQnmgULzTevdpPz0NJ",
			Lat:         55.762394,
			Lng:         37.578684,
		},
	)

	return nearbyPlaces, nil
}
