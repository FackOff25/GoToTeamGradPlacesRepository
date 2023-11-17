package usecase

import (
	"context"

	"github.com/FackOff25/GoToTeamGradPlacesRepository/internal/repository"
)

type UseCase struct {
	Repo *repository.Repo
	Ctx  context.Context
}
