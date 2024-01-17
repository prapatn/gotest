package services

import "github.com/stretchr/testify/mock"

type promotionServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *promotionServiceMock {
	return &promotionServiceMock{}
}

func (m *promotionServiceMock) CalculateDiscount(amount int) (int, error) {
	agrs := m.Called(amount)
	return agrs.Int(0), agrs.Error(1)
}
