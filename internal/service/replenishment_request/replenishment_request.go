package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
)

type ReplenishmentRequestStorage interface {
	Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error
}

type CashManagerStorage interface {
	GetByShogunUsername(ctx context.Context, username string) (domain.CashManager, error)
}

type DaimyoStorage interface {
	GetByUsername(ctx context.Context, username string) (domain.Daimyo, error)
}

type CardStorage interface {
	GetByName(ctx context.Context, name string) (domain.Card, error)
}

type ReplenishmentRequestStatusStorage interface {
	GetId(ctx context.Context, name string) (int, error)
}

type ReplenishmentRequestService struct {
	storage                           ReplenishmentRequestStorage
	cashManagerStorage                CashManagerStorage
	daimyoStorage                     DaimyoStorage
	cardStorage                       CardStorage
	replenishmentRequestStatusStorage ReplenishmentRequestStatusStorage
}

func NewReplenishmentRequestService(
	storage ReplenishmentRequestStorage,
	cashManager CashManagerStorage,
	daimyo DaimyoStorage,
	card CardStorage,
	replenishmentRequestStatus ReplenishmentRequestStatusStorage,
) *ReplenishmentRequestService {
	return &ReplenishmentRequestService{
		storage:                           storage,
		cashManagerStorage:                cashManager,
		daimyoStorage:                     daimyo,
		cardStorage:                       card,
		replenishmentRequestStatusStorage: replenishmentRequestStatus,
	}
}

func (s *ReplenishmentRequestService) Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error) {
	daimyo, err := s.daimyoStorage.GetByUsername(ctx, dto.DaimyoUsername)
	if err != nil {
		return 0, err
	}

	cashManager, err := s.cashManagerStorage.GetByShogunUsername(ctx, daimyo.ShogunUsername)
	if err != nil {
		return 0, err
	}

	card, err := s.cardStorage.GetByName(ctx, dto.CardName)
	if err != nil {
		return 0, err
	}

	statusId, err := s.replenishmentRequestStatusStorage.GetId(ctx, domain.ActiveRequest.String())
	if err != nil {
		return 0, err
	}

	replenishmentReq := domain.ReplenishmentRequest{
		CashManagerUsername: cashManager.Username,
		DaimyoUsername:      dto.DaimyoUsername,
		CardId:              card.CardId,
		Amount:              dto.Amount,
		StatusId:            statusId,
	}

	if err := s.storage.Insert(ctx, replenishmentReq); err != nil {
		return 0, err
	}

	return tg.ChatID(cashManager.ChatId), nil
}
