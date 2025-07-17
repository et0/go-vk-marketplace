package router

import (
	"github.com/et0/go-vk-marketplace/internal/handler"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/et0/go-vk-marketplace/internal/storage"
	"github.com/labstack/echo/v4"
)

func New(db storage.Database) *echo.Echo {
	e := echo.New()

	// Инициализация сервисов для работы с бд
	userService := service.NewUserService(db)

	// Инициализация хендлеров
	userHandler := handler.NewUserHandler(userService)

	e.POST("/register", userHandler.Register)

	return e
}
