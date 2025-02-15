package service

import (
	"testing"

	"github.com/timut2/avito_test_task/internal/repository"
	repoPurchaseMock "github.com/timut2/avito_test_task/internal/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBuyItemError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repoPurchaseMock.NewMockPurchaseRepo(ctrl)

	mockRepo.EXPECT().Insert(1, -1).Return(repository.ErrItemNotFound).Times(1)

	buyService := NewBuyService(mockRepo)

	err := buyService.BuyItem(1, -1)

	assert.Error(t, err)
	assert.NotNil(t, err)

}
