package services

import "gotest/repositories"

type PromotionService interface {
	CalculateDiscount(amount int) (int, error)
}

type promotionService struct {
	repo repositories.PromotionRepository
}

func NewPromotionService(repo repositories.PromotionRepository) PromotionService {
	return promotionService{repo: repo}
}

func (s promotionService) CalculateDiscount(amount int) (int, error) {
	if amount <= 0 {
		return 0, ErrZeroAmount
	}

	promotion, err := s.repo.GetPromotion()
	if err != nil {
		return 0, ErrRepository
	}

	if amount >= promotion.PurchaseMin {
		return amount - (promotion.DiscountPercent * amount / 100), nil
	}

	return amount, nil
}
