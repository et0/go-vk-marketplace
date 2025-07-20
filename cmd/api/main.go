package main

import (
	"log"

	router "github.com/et0/go-vk-marketplace/internal"
	"github.com/et0/go-vk-marketplace/internal/config"
	"github.com/et0/go-vk-marketplace/internal/storage/postgres"
)

func main() {
	cfg, err := config.Load("config/local.yaml")
	if err != nil {
		log.Fatalf("Config load: %s", err)
	}

	pg, err := postgres.New(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Basename)
	if err != nil {
		log.Fatalf("DB init failed: %s", err)
	}
	defer pg.Close()

	e := router.New(pg, cfg.HTTP.JWTSecret)
	if err := e.Start(":" + cfg.HTTP.Port); err != nil {
		log.Fatalf("Server start failed: %s", err)
	}
}
