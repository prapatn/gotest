package services_test

import (
	"errors"
	"gotest/repositories"
	"gotest/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
		{name: "error 0", purchaseMin: 100, discountPercent: 20, amount: 0, expected: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {

			//Arrage
			repository := repositories.NewPromotionRepositoryMock()
			repository.On("GetPromotion").Return(
				repositories.Promotion{
					ID:              1,
					PurchaseMin:     c.purchaseMin,
					DiscountPercent: c.discountPercent,
				},
				nil,
			)

			service := services.NewPromotionService(repository)

			//Act
			discount, err := service.CalculateDiscount(c.amount)
			expected := c.expected

			//Assert
			assert.Equal(t, expected, discount)

			if c.amount == 0 {
				assert.ErrorIs(t, services.ErrZeroAmount, err)
				repository.AssertNotCalled(t, "GetPromotion")
			}
		})
	}

	t.Run("repo error", func(t *testing.T) {
		//Arrage
		repository := repositories.NewPromotionRepositoryMock()
		repository.On("GetPromotion").Return(
			repositories.Promotion{},
			errors.New(""),
		)
		service := services.NewPromotionService(repository)

		//Act
		_, err := service.CalculateDiscount(100)
		assert.ErrorIs(t, services.ErrRepository, err)

	})
}
