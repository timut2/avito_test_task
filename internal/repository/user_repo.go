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

var (
	ErrUserNotFound = errors.New("user not found")
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	if err := user.HashPassword(); err != nil {
		fmt.Println(err)
		return err
	}

	query := `INSERT INTO USERS (username, password) values ($1, $2)`
	_, err := r.db.Exec(query, user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
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
