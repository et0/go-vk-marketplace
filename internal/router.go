package router

import (
	"github.com/et0/go-vk-marketplace/internal/handler"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/et0/go-vk-marketplace/internal/storage"
	"github.com/labstack/echo/v4"
)

func New(db storage.Database, jwtSecret string) *echo.Echo {
	e := echo.New()

	// Инициализация сервисов для работы с бд
	userService := service.NewUserService(db)

	// Инициализация хендлеров
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(userService, jwtSecret)

	e.POST("/signup", userHandler.Signup)
	e.POST("/login", authHandler.Login)

	return e
}
