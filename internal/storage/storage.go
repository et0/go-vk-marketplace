package storage

import "github.com/et0/go-vk-marketplace/internal/model"

type Database interface {
	CreateUser(username, password string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByID(id uint) (*model.User, error)

	CreateAd(ad *model.Ad) (*model.Ad, error)
}
