package service

import (
	"errors"

	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

var (
	ErrInvalidPassword = errors.New("invalid password")
)

func (s *AuthService) Authentication(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByUserName(username)

	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			newUser := &models.User{
				Username: username,
				Password: password,
			}
			if err := s.userRepo.Create(newUser); err != nil {
				return nil, err
			}
			return newUser, nil
		}
		return nil, err
	}

	if err := user.CheckPassword(password); err != nil {
		return nil, ErrInvalidPassword
	}

	return user, nil
}
