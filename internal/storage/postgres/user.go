package postgres

import (
	"context"
	"log"

	"github.com/et0/go-vk-marketplace/internal/model"
)

func (p *Postgres) GetAll() (*model.User, error) {
	conn, err := p.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer conn.Release()

	var user model.User

	err = conn.QueryRow(context.Background(), "SELECT username,password FROM users limit 1").Scan(
		&user.Username,
		&user.Password,
	)

	return &user, err
}
