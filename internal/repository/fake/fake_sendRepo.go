package fake

import (
	"errors"
	"strconv"

	"github.com/timut2/avito_test_task/internal/models"
)

type fakeSendRepo struct {
	userRepo     *fakeUserRepo
	transactions []models.Transaction
}

func NewFakeSendRepo() *fakeSendRepo {
	userRepo := &fakeUserRepo{
		users: map[int64]*models.User{
			1: {Username: "Alice", Coins: 1000},
			2: {Username: "Bob", Coins: 1000},
		},
	}
	return &fakeSendRepo{
		userRepo:     userRepo,
		transactions: []models.Transaction{},
	}
}

func (r *fakeSendRepo) TransactionInfo(userID int) (int, int) {
	balance, err := r.userRepo.BalanceById(userID)
	if err != nil {
		return 0, 0
	}
	balance2, err := r.userRepo.BalanceById(2) // Or another user ID
	if err != nil {
		return 0, 0
	}
	return balance, balance2
}

func (r *fakeSendRepo) Send(userId int, req models.SendCoinRequest) error {
	// Ensure sender exists
	sender, exists := r.userRepo.users[int64(userId)]
	if !exists {
		return errors.New("отправитель не найден")
	}

	// Ensure receiver exists
	var receiver *models.User
	for _, user := range r.userRepo.users {
		if user.Username == req.ToUser {
			receiver = user
			break
		}
	}
	if receiver == nil {
		return errors.New("получатель не найден")
	}

	// Prevent self-transfer
	if sender.Username == receiver.Username {
		return errors.New("нельзя отправлять средства самому себе")
	}

	// Ensure sender has enough coins
	if sender.Coins < req.Amount {
		return errors.New("недостаточно монет у отправителя")
	}

	// Perform transaction
	sender.Coins -= req.Amount
	receiver.Coins += req.Amount

	// Log transaction (convert int64 to string)
	r.transactions = append(r.transactions, models.Transaction{
		FromUser: strconv.FormatInt(int64(userId), 10),
		ToUser:   strconv.FormatInt(int64(receiver.Id), 10),
		Amount:   req.Amount,
	})

	return nil
}
