package service

import (
	"github.com/timut2/avito_test_task/internal/repository"
)

type BuyService struct {
	PurchaseRepository repository.PurchaseRepo
}

func NewBuyService(purchasesRepo repository.PurchaseRepo) *BuyService {
	return &BuyService{PurchaseRepository: purchasesRepo}
}

func (s *BuyService) BuyItem(userID, itemID int) error {
	err := s.PurchaseRepository.Insert(userID, itemID)
	if err != nil {
		return err
	}

	return nil
}
