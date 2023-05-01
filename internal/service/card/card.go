package card

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type CardStorage interface {
	Insert(ctx context.Context, card domain.Card) error
}

type CardService struct {
	transactor storage.Transactor

	storage CardStorage
}

func NewCardService(t storage.Transactor, s CardStorage) *CardService {
	return &CardService{
		transactor: t,
		storage:    s,
	}
}

func (s *CardService) Create(ctx context.Context, dto domain.CardDTO) error {
	card := domain.Card{
		Name:           dto.Name,
		LastDigits:     dto.LastDigits,
		DailyLimit:     dto.DailyLimit,
		DaimyoUsername: dto.DaimyoUsername,
	}

	return s.storage.Insert(ctx, card)
}
