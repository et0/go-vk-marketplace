package storage

import "github.com/et0/go-vk-marketplace/internal/model"

type Database interface {
	GetAll() (*model.User, error)
}
