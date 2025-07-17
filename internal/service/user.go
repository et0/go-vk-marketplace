package service

import (
	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/storage"
)

type UserService struct {
	db storage.Database
}

func NewUserService(db storage.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) FindByUsername(username string) (*model.User, error) {
	return s.db.FindByUsername(username)
}

func (s *UserService) Create(username, password string) (*model.User, error) {
	return s.db.Create(username, password)
}
