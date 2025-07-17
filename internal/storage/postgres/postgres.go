package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const URI = "postgres://%s:%s@%s:%s/%s"

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(username, password, host, port, basename string) (*Postgres, error) {
	connString := fmt.Sprintf(URI, username, password, host, port, basename)

	cfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	cfg.MaxConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}
	return &Postgres{Pool: pool}, err
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
