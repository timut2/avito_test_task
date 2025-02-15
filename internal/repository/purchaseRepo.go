package repository

import (
	"database/sql"
	"fmt"
	"log"
)

type PurchaseRepository struct {
	db             *sql.DB
	userRepository UserRepo
	ItemRepository ItemRepo
}

type PurchaseRepo interface {
	Insert(userID, itemID int) error
}

func NewPurchaseRepository(db *sql.DB, userRepo *UserRepository, itemRepo *ItemRepository) *PurchaseRepository {
	return &PurchaseRepository{db: db, userRepository: userRepo, ItemRepository: itemRepo}
}

func (r *PurchaseRepository) Insert(userID, itemID int) error {
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

	userBalance, err := r.userRepository.BalanceById(userID)
	if err != nil {
		return fmt.Errorf("не удалось найти получателя: %w", err)
	}

	itemCost, err := r.ItemRepository.GetCostByID(itemID)
	if err != nil {
		return err
	}

	if userBalance < itemCost {
		return fmt.Errorf("недостаточно монет, чтобы купить товар")
	}

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", itemCost, userID)
	if err != nil {
		return fmt.Errorf("не удалось уменьшить баланс отправителя: %w", err)
	}

	_, err = tx.Exec("INSERT INTO purchases (user_id, item_id) VALUES ($1, $2)", userID, itemID)
	if err != nil {
		return fmt.Errorf("не удалось записать транзакцию: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("не удалось подтвердить транзакцию: %w", err)
	}

	return nil
}
