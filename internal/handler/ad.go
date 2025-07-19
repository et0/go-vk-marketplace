package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/et0/go-vk-marketplace/internal/model"
	"github.com/et0/go-vk-marketplace/internal/service"
	"github.com/go-playground/validator/v10"
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

	if err := validator.New().Struct(req); err != nil {
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
		IsMine:      true,
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *AdHandler) GetAll(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit <= 0 {
		limit = 10
	}

	sortBy := c.QueryParam("sort_by")
	if sortBy == "" || (sortBy != "price" && sortBy != "created_at") {
		sortBy = "created_at"
	}

	order := c.QueryParam("order")
	if order == "" || (order != "asc" && order != "desc") {
		order = "desc"
	}

	minPrice, _ := strconv.Atoi(c.QueryParam("min_price"))
	if minPrice < 0 {
		minPrice = 0
	}
	maxPrice, _ := strconv.Atoi(c.QueryParam("max_price"))
	if maxPrice < 0 {
		maxPrice = 0
	}

	ads, err := h.adService.GetAll(page, limit, sortBy, order, minPrice, maxPrice)
	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to fetch listings")
	}

	userID, _ := c.Get("userID").(uint)

	response := make([]model.AdResponse, len(ads))
	for i, ad := range ads {
		isMine := false
		if userID != 0 && ad.UserID == userID {
			isMine = true
		}

		response[i] = model.AdResponse{
			ID:          ad.ID,
			Title:       ad.Title,
			Description: ad.Description,
			ImageURL:    ad.ImageURL,
			Price:       ad.Price,
			CreatedAt:   ad.CreatedAt,
			Author:      ad.User.Username,
			IsMine:      isMine,
		}
	}

	fmt.Println(response)

	return c.JSON(http.StatusOK, response)
}
