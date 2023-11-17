package queries

const (
	getPlaceUuidQuery    = `SELECT id FROM places WHERE place_id = $1;`
	getUserPlaceReaction = `SELECT like_mark, visited_mark FROM users_places WHERE place_id = $1 AND user_id = $2;`
)

func (q *Queries) GetPlaceUuid(gID string) (string, error) {
	var uuid string
	row := q.Pool.QueryRow(q.Ctx, getPlaceUuidQuery, gID)

	err := row.Scan(&uuid)
	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (q *Queries) GetUserReaction(userId string, placeId string) (bool, bool, error) {
	var likeFlag, visitedFlag bool
	r := q.Pool.QueryRow(q.Ctx, getUserPlaceReaction, &placeId, &userId)

	err := r.Scan(&likeFlag, &visitedFlag)
	return likeFlag, visitedFlag, err
}
