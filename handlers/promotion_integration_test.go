//go:build integration

package handlers_test

import (
	"fmt"
	"gotest/handlers"
	"gotest/repositories"
	"gotest/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscountIntegrationService(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		repo := repositories.NewPromotionRepositoryMock()
		repo.On("GetPromotion").Return(repositories.Promotion{
			ID:              1,
			DiscountPercent: 20,
			PurchaseMin:     100,
		}, nil)

		service := services.NewPromotionService(repo)
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

}
