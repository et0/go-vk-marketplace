package main

import (
	"fmt"
	"log"

	"github.com/et0/go-vk-marketplace/internal/config"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/et0/go-vk-marketplace/internal/storage/postgres"
)

func main() {
	cfg, err := config.Load("config/local.yaml")
	if err != nil {
		log.Fatalf("Config load: %s", err)
	}

	pg, err := postgres.New(cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Basename)
	if err != nil {
		log.Fatal("DB init failed:", err)
	}
	defer pg.Close()

	userService := service.NewUserService(pg)
	fmt.Println(userService.GetAll())
}
