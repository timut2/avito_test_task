package service

import (
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
)

type SendService struct {
	sendRepo repository.SendRepo
}

func NewSendService(sendRepo repository.SendRepo) *SendService {
	return &SendService{sendRepo: sendRepo}
}

func (s *SendService) Send(userID int, req models.SendCoinRequest) error {
	err := s.sendRepo.Send(userID, req)
	if err != nil {
		return err
	}
	return nil
}
