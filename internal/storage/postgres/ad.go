package postgres

import (
	"context"
	"log"

	"github.com/et0/go-vk-marketplace/internal/model"
)

func (p *Postgres) CreateAd(ad *model.Ad) (*model.Ad, error) {
	conn, err := p.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer conn.Release()

	var id int
	err = conn.QueryRow(context.Background(),
		"INSERT INTO ads (title, description, image_url, price, user_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		ad.Title, ad.Description, ad.ImageURL, ad.Price, ad.UserID).Scan(&id)
	if err != nil {
		return nil, err
	}

	var newAd model.Ad
	sql := `SELECT a.id, a.title, a.description, a.image_url, a.price, a.created_at, u.username FROM ads a
JOIN users u ON u.id = a.user_id WHERE a.id = $1`
	err = conn.QueryRow(context.Background(), sql, id).
		Scan(&newAd.ID, &newAd.Title, &newAd.Description, &newAd.ImageURL, &newAd.Price,
			&newAd.CreatedAt, &newAd.User.Username)
	if err != nil {
		return nil, err
	}

	return &newAd, err
}
