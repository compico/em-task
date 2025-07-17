package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionConfig interface {
	GetDsn() string
}

func NewConnection(ctx context.Context, cfg ConnectionConfig) (*pgxpool.Pool, func(), error) {
	pool, err := pgxpool.New(ctx, cfg.GetDsn())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, pool.Close, nil
}
