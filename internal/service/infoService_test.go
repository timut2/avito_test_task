package service

import (
	"testing"

	"github.com/timut2/avito_test_task/internal/repository"
	repoInfoMock "github.com/timut2/avito_test_task/internal/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetInformation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoInfoMock.NewMockInfoRepo(ctrl)

	mockRepo.EXPECT().Get(-1).Return(nil, repository.ErrUserNotFound).Times(1)

	infoService := NewInfoService(mockRepo)

	info, err := infoService.GetInformation(-1)

	assert.Error(t, err)
	assert.Nil(t, info)
	assert.Equal(t, err.Error(), "пользователь не найден")
}
