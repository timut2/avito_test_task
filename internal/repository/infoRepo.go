package repository

import (
	"database/sql"
	"errors"

	"github.com/timut2/avito_test_task/internal/models"
)

type InfoRepository struct {
	db *sql.DB
}

func NewInfoRepository(db *sql.DB) *InfoRepository {
	return &InfoRepository{db: db}
}

type InfoRepo interface {
	Get(userID int) (*models.InfoResponse, error)
}

func (r *InfoRepository) Get(userID int) (*models.InfoResponse, error) {
	var info models.InfoResponse
	query := `SELECT coins FROM users WHERE id = $1`
	err := r.db.QueryRow(query, userID).Scan(&info.Coins)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("пользователь не найден")
		}
		return nil, errors.New("ошибка при получении данных о пользователе")
	}

	info.Inventory = []models.Item{}
	queryInventory := `SELECT items.name, COUNT(purchases.item_id) FROM purchases JOIN items ON items.id = purchases.item_id WHERE user_id = $1 GROUP BY items.name`
	rows, err := r.db.Query(queryInventory, userID)
	if err != nil {
		return nil, errors.New("ошибка при получении инвентаря пользователя")
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Type, &item.Quantity); err != nil {
			return nil, errors.New("ошибка при обработке данных инвентаря")
		}
		info.Inventory = append(info.Inventory, item)
	}

	info.CoinHistory.Received = []models.Transaction{}
	queryReceived := `SELECT users.username, amount FROM transactions JOIN users ON users.id = transactions.from_user_id WHERE to_user_id = $1`
	rowsReceived, err := r.db.Query(queryReceived, userID)
	if err != nil {
		return nil, errors.New("ошибка при получении истории полученных монет")
	}
	defer rowsReceived.Close()

	for rowsReceived.Next() {
		var transaction models.Transaction
		if err := rowsReceived.Scan(&transaction.FromUser, &transaction.Amount); err != nil {
			return nil, errors.New("ошибка при обработке данных полученных транзакций")
		}
		info.CoinHistory.Received = append(info.CoinHistory.Received, transaction)
	}

	info.CoinHistory.Send = []models.Transaction{}
	querySent := `SELECT users.username, amount FROM transactions JOIN users ON users.id = transactions.to_user_id WHERE from_user_id = $1`
	rowsSent, err := r.db.Query(querySent, userID)
	if err != nil {
		return nil, errors.New("ошибка при получении истории отправленных монет")
	}
	defer rowsSent.Close()

	for rowsSent.Next() {
		var transaction models.Transaction
		if err := rowsSent.Scan(&transaction.ToUser, &transaction.Amount); err != nil {
			return nil, errors.New("ошибка при обработке данных отправленных транзакций")
		}
		info.CoinHistory.Send = append(info.CoinHistory.Send, transaction)
	}

	return &info, nil
}
