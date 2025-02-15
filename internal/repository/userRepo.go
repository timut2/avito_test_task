package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/timut2/avito_test_task/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepo interface {
	Create(user *models.User) (int, error)
	FindByUserName(username string) (*models.User, error)
	BalanceById(userId int) (int, error)
}

var ErrUserNotFound = errors.New("пользователь не найден")

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) (int, error) {
	if err := user.HashPassword(); err != nil {
		fmt.Println("Error hashing password:", err)
		return -1, err
	}

	query := `INSERT INTO USERS (username, password) VALUES ($1, $2) RETURNING id`
	var newUserID int
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&newUserID)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return -1, err
	}

	return newUserID, nil
}

func (r *UserRepository) FindByUserName(username string) (*models.User, error) {
	var user models.User

	query := `SELECT id, username, password, coins FROM users where username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.Id, &user.Username, &user.Password, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) BalanceById(userId int) (int, error) {
	var user models.User

	query := `SELECT coins FROM users where id = $1`
	err := r.db.QueryRow(query, userId).Scan(&user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrUserNotFound
		}
		return 0, err
	}

	return user.Coins, nil
}
