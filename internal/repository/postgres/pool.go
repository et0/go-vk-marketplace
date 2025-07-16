package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init(connString string) error {
	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	cfg.MaxConns = 10

	Pool, err = pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return fmt.Errorf("failed to create pool: %w", err)
	}
	return nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
	}
}
