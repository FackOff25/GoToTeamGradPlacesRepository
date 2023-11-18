package repository

type QueriesInterface interface {
	GetPlaceUuid(gID string) (string, error)
	GetUserReaction(userId string, placeId string) (bool, bool, error)
}
