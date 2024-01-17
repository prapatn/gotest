package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("succss", func(t *testing.T) {

		//Arrange
		amount := 100
		expected := 80

		service := services.NewPromotionServiceMock()
		service.On("CalculateDiscount", amount).Return(expected, nil)

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		//Assert
		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {
			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}

	})

	t.Run("amount is not number", func(t *testing.T) {

		//Arrange
		amount := "a"
		// expected := 80

		service := services.NewPromotionServiceMock()
		// service.On("CalculateDiscount", amount).Return(expected, nil)

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("service error", func(t *testing.T) {

		//Arrange
		amount := 0
		expected := 0

		service := services.NewPromotionServiceMock()
		service.On("CalculateDiscount", amount).Return(expected, services.ErrZeroAmount)

		handler := handlers.NewPromotionHandler(service)

		app := fiber.New()
		app.Get("/calculate", handler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)

		//Act
		res, _ := app.Test(req)
		defer res.Body.Close()

		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	})

}
