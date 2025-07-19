package router

import (
	"github.com/et0/go-vk-marketplace/internal/handler"
	myMD "github.com/et0/go-vk-marketplace/internal/middleware"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/et0/go-vk-marketplace/internal/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(db storage.Database, jwtSecret string) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(myMD.CheckToken(jwtSecret))

	// Инициализация сервисов для работы с бд
	userService := service.NewUserService(db)
	adService := service.NewAdService(db)

	// Инициализация хендлеров
	userHandler := handler.NewUserHandler(userService, jwtSecret)
	adHandler := handler.NewAdHandler(userService, adService)

	e.POST("/signup", userHandler.Signup)
	e.POST("/login", userHandler.Login)
	e.GET("/ads", adHandler.GetAll)
	e.POST("/ads", adHandler.Create, myMD.IsAuth)

	return e
}
