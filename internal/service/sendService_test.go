package service

import (
	"errors"
	"testing"

	"github.com/timut2/avito_test_task/internal/models"
	repoSendMock "github.com/timut2/avito_test_task/internal/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoSendMock.NewMockSendRepo(ctrl)

	req := models.SendCoinRequest{
		ToUser: "",
		Amount: 100,
	}

	mockRepo.EXPECT().Send(1, req).Return(errors.New("не удалось найти получателя:")).Times(1)

	sendService := NewSendService(mockRepo)

	err := sendService.Send(1, req)

	assert.Error(t, err)
	assert.Equal(t, err.Error(), "не удалось найти получателя:")
}
