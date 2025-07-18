package handler

import (
	"fmt"
	"net/http"

	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/labstack/echo/v4"
)

type AdHandler struct {
	userService *service.UserService
	adService   *service.AdService
}

func NewAdHandler(userService *service.UserService, adService *service.AdService) *AdHandler {
	return &AdHandler{userService: userService, adService: adService}
}

func (h *AdHandler) Create(c echo.Context) error {
	userID := c.Get("userID").(uint)

	var req model.AdNewRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ad := &model.Ad{
		Title:       req.Title,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		Price:       req.Price,
		UserID:      userID,
	}

	createdAd, err := h.adService.Create(ad)
	fmt.Println(createdAd, err)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create listing")
	}

	response := model.AdResponse{
		ID:          createdAd.ID,
		Title:       createdAd.Title,
		Description: createdAd.Description,
		ImageURL:    createdAd.ImageURL,
		Price:       createdAd.Price,
		CreatedAt:   createdAd.CreatedAt,
		Author:      createdAd.User.Username,
	}

	return c.JSON(http.StatusCreated, response)
}
