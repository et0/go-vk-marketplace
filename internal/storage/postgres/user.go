package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/jackc/pgx/v5"
)

func (p *Postgres) FindByUsername(username string) (*model.User, error) {
	conn, err := p.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer conn.Release()

	var user model.User

	err = conn.QueryRow(context.Background(), "SELECT * FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &user, err
}

func (p *Postgres) Create(username, password string) (*model.User, error) {
	conn, err := p.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer conn.Release()

	var id int
	err = conn.QueryRow(context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", username, password).Scan(&id)
	if err != nil {
		return nil, err
	}

	var user model.User
	err = conn.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, err
}
