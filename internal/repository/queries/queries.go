package queries

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Queries struct {
	Ctx context.Context
	pgxpool.Pool
}
