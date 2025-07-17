package service

import (
	"fmt"

	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/storage"
)

type UserService struct {
	db storage.Database
}

func NewUserService(db storage.Database) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetAll() (*model.User, error) {
	user, err := s.db.GetAll()
	if err != nil {
		return user, fmt.Errorf("DB GetAll: %w", err)
	}

	return user, nil
}
