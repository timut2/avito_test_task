package service

import (
	"testing"

	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
	repoUserMock "github.com/timut2/avito_test_task/internal/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoUserMock.NewMockUserRepo(ctrl)

	username := "testuser"
	password := "password123"
	user := &models.User{Username: username, Password: password}

	mockRepo.EXPECT().FindByUserName(username).Return(nil, repository.ErrUserNotFound).Times(1)
	mockRepo.EXPECT().Create(user).Return(1, nil).Times(1) // Mock user creation with ID 1

	authService := NewAuthService(mockRepo)

	newUser, err := authService.Authentication(username, password)

	assert.NoError(t, err)
	assert.NotNil(t, newUser)
	assert.Equal(t, username, newUser.Username)
	assert.Equal(t, password, newUser.Password)
}
