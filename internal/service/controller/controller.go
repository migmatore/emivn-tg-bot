package controller

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type ControllerStorage interface {
}

type ControllerService struct {
	transactor storage.Transactor

	storage ControllerStorage
}

func New(transactor storage.Transactor, storage ControllerStorage) *ControllerService {
	return &ControllerService{
		transactor: transactor,
		storage:    storage,
	}
}

func (s *ControllerService) CreateTurnover(ctx context.Context, dto domain.ControllerTurnoverDTO) error {

	return nil
}
