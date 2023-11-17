package repository

import (
	// "github.com/jackc/pgx/v5"
	"context"
)

type Repo struct {
	ctx context.Context
	QueriesInterface
}

func New(queries QueriesInterface, ctx context.Context) *Repo {
	return &Repo{
		QueriesInterface: queries,
		ctx:              ctx,
	}
}
