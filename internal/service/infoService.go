package service

import (
	"github.com/timut2/avito_test_task/internal/models"
	"github.com/timut2/avito_test_task/internal/repository"
)

type InfoService struct {
	infoRepo repository.InfoRepo
}

func NewInfoService(infoRepo repository.InfoRepo) *InfoService {
	return &InfoService{infoRepo: infoRepo}
}

func (s *InfoService) GetInformation(userId int) (*models.InfoResponse, error) {
	info, err := s.infoRepo.Get(userId)
	if err != nil {
		return nil, err
	}
	return info, nil
}
