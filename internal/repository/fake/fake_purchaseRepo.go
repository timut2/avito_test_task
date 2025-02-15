package fake

import (
	"errors"

	"github.com/timut2/avito_test_task/internal/models"
)

type fakePurchaseRepo struct {
	userRepo     *fakeUserRepo
	itemRepo     *fakeItemRepo
	PurchaseRepo map[int]int
}

func NewFakePurchaseRepo() *fakePurchaseRepo {
	userRepo := &fakeUserRepo{
		users: map[int64]*models.User{
			1: {Username: "Alice", Coins: 1000},
			2: {Username: "Bob", Coins: 1000},
		},
	}
	return &fakePurchaseRepo{
		userRepo:     userRepo,
		itemRepo:     NewFakeItemRepo(),
		PurchaseRepo: make(map[int]int),
	}
}

// Now takes a userID argument to return the correct balance
func (r *fakePurchaseRepo) GiveInfo(userID, itemID int) (int, int) {
	balance, _ := r.userRepo.BalanceById(userID)
	itemCost, _ := r.itemRepo.GetCostByID(itemID)
	return balance, itemCost
}

func (r *fakePurchaseRepo) Insert(userID, itemID int) error {
	userBalance, err := r.userRepo.BalanceById(userID)
	if err != nil {
		return err
	}

	itemCost, err := r.itemRepo.GetCostByID(itemID)
	if err != nil {
		return err
	}

	if userBalance < itemCost {
		return errors.New("недостаточно монет, чтобы купить товар")
	}

	// Deduct coins from user
	r.userRepo.users[int64(userID)].Coins -= itemCost

	// Store purchase
	r.PurchaseRepo[userID] = itemID
	return nil
}
