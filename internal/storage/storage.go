package storage

import "github.com/et0/go-vk-marketplace/internal/model"

type Database interface {
	Create(username, password string) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
}
