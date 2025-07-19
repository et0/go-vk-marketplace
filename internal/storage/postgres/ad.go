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

func (p *Postgres) GetAll(page, limit int, sortBy, order string, minPrice, maxPrice int) ([]*model.Ad, error) {
	conn, err := p.Pool.Acquire(context.Background())
	if err != nil {
		log.Fatal("DB connect failed:", err)
	}
	defer conn.Release()

	sql := `SELECT
	a.id, a.title, a.description, a.image_url, a.price, a.created_at, a.user_id, u.username 
FROM ads a
LEFT JOIN users u ON u.id = a.user_id WHERE a.price BETWEEN $1 AND $2`

	if maxPrice == 0 || maxPrice < minPrice {
		maxPrice = 2147483647
	}

	sql += " ORDER BY a." + sortBy
	if order == "desc" {
		sql += " DESC"
	} else {
		sql += " ASC"
	}

	sql += " LIMIT $3 OFFSET $4"

	rows, err := conn.Query(context.Background(), sql, minPrice, maxPrice, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}

	ads := make([]*model.Ad, 0, limit)
	for rows.Next() {
		var a model.Ad
		if err := rows.Scan(&a.ID, &a.Title, &a.Description, &a.ImageURL, &a.Price, &a.CreatedAt, &a.UserID, &a.User.Username); err != nil {
			return nil, err
		}
		ads = append(ads, &a)
	}

	return ads, nil
}
