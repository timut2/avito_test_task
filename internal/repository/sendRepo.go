package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/timut2/avito_test_task/internal/models"
)

type SendRepository struct {
	db             *sql.DB
	userRepository UserRepo
}

type SendRepo interface {
	Send(userId int, req models.SendCoinRequest) error
}

func NewSendRepository(db *sql.DB, userRepo UserRepo) *SendRepository {
	return &SendRepository{db: db, userRepository: userRepo}
}

func (r *SendRepository) Send(userId int, req models.SendCoinRequest) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию: %w", err)
	}

	defer func() {
		if err != nil {
			err2 := tx.Rollback()
			if err2 != nil {
				log.Printf("Ошибка при откате транзакции: %v\n", err2)
			}
		}
	}()

	receiver, err := r.userRepository.FindByUserName(req.ToUser)
	if err != nil {
		return fmt.Errorf("не удалось найти получателя: %w", err)
	}

	if receiver.Id == userId {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Printf("Ошибка при откате транзакции: %v\n", err2)
		}
		return fmt.Errorf("нельзя отправлять средства со своего аккаунта на свой же аккаунт")
	}

	var senderBalance int
	err = tx.QueryRow("SELECT coins FROM users WHERE id = $1 FOR UPDATE", userId).Scan(&senderBalance)
	if err != nil {
		return fmt.Errorf("не удалось получить баланс отправителя: %w", err)
	}

	if senderBalance < req.Amount {
		err2 := tx.Rollback()
		if err2 != nil {
			log.Printf("Ошибка при откате транзакции: %v\n", err2)
		}
		return fmt.Errorf("недостаточно монет у отправителя")
	}

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", req.Amount, userId)
	if err != nil {
		return fmt.Errorf("не удалось уменьшить баланс отправителя: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins + $1 WHERE username = $2", req.Amount, req.ToUser)
	if err != nil {
		return fmt.Errorf("не удалось увеличить баланс получателя: %w", err)
	}

	_, err = tx.Exec("INSERT INTO transactions (from_user_id, to_user_id, amount) VALUES ($1, $2, $3)", userId, receiver.Id, req.Amount)
	if err != nil {
		return fmt.Errorf("не удалось записать транзакцию: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("не удалось подтвердить транзакцию: %w", err)
	}

	return nil
}
