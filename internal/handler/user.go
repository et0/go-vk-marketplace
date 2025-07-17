package handler

import (
	"net/http"

	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/service"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (h *UserHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existingUser, err := h.userService.FindByUsername(req.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error", err)
	}
	if existingUser != nil {
		return echo.NewHTTPError(http.StatusConflict, "username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	createdUser, err := h.userService.Create(req.Username, string(hashedPassword))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	response := model.User{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}
