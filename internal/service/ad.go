package service

import (
	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/storage"
)

type AdService struct {
	db storage.Database
}

func NewAdService(db storage.Database) *AdService {
	return &AdService{db: db}
}

func (s *AdService) Create(ad *model.Ad) (*model.Ad, error) {
	return s.db.CreateAd(ad)
}

func (s *AdService) GetAll(page, limit int, sortBy, order string, minPrice, maxPrice int) ([]*model.Ad, error) {
	return s.db.GetAll(page, limit, sortBy, order, minPrice, maxPrice)
}
