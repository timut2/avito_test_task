package repository

import (
	"database/sql"
	"errors"
)

var ErrItemNotFound = errors.New("айтем не найден")

type ItemRepository struct {
	db *sql.DB
}

type ItemRepo interface {
	GetCostByID(itemID int) (int, error)
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) GetCostByID(itemID int) (int, error) {
	var itemCost int
	query := `SELECT price FROM items WHERE id = $1`
	err := r.db.QueryRow(query, itemID).Scan(&itemCost)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrItemNotFound // Return custom error if item does not exist
		}
		return 0, err // Return other database errors
	}

	return itemCost, nil
}
