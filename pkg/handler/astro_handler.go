package handler

import (
	"AstrologService/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetImageOfTheDay(c echo.Context) error {
	today := time.Now()
	fmt.Println(today)
	image, err := getImageOfTheDayFromNASA(os.Getenv("API_KEY"), today)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	_, err = h.service.InsertImage(context.Background(), image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, image)
}

func (h *Handler) GetImageByDate(c echo.Context) error {
	dateStr := c.Param("date") // Используйте Param для получения даты из URL
	if dateStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Date parameter is empty"})
	}

	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format: " + err.Error()})
	}
	fmt.Println("handler:", date)
	image, err := h.service.GetImageByDate(context.Background(), date)
	if err != nil {

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error: " + err.Error()})
	}

	return c.JSON(http.StatusOK, image)
}

func (h *Handler) GetAllImages(c echo.Context) error {
	images, err := h.service.GetAllImages(context.Background())
	if err != nil {
		log.Printf("Ошибка:&v", err)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}

	return c.JSON(http.StatusOK, images)
}

func (h *Handler) DeleteImageByDate(c echo.Context) error {
	dateStr := c.Param("date")
	date, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		log.Printf("Internal server error: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid date format"})
	}

	rowsAffected, err := h.service.DeleteImageByDate(context.Background(), date)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Image not found"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Image deleted"})
}

func getImageOfTheDayFromNASA(apiKey string, date time.Time) (*model.ImageModel, error) {
	nasaAPIURL := os.Getenv("NASA_API")
	queryParams := url.Values{}
	queryParams.Add("api_key", apiKey)
	queryParams.Add("date", date.Format(time.DateOnly))
	nasaAPIURL += "?" + queryParams.Encode()
	resp, err := http.Get(nasaAPIURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("NASA API returned non-200 status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("NASA API returned non-200 status code: %d", resp.StatusCode)
	}

	var apodImage model.ImageModel
	err = json.NewDecoder(resp.Body).Decode(&apodImage)
	if err != nil {
		return nil, err
	}

	return &apodImage, nil
}
