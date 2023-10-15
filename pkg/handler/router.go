package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRouter() *echo.Echo {
	router := echo.New()

	// Маршруты для изображений APOD
	router.GET("/api/apod", h.GetImageOfTheDay)
	router.GET("/api/:date", h.GetImageByDate)
	router.GET("/api/apod/album", h.GetAllImages)

	return router
}
