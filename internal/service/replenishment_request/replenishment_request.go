package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
)

type ReplenishmentRequestStorage interface {
	Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error
	CheckIfExists(ctx context.Context, cardName string) (bool, error)
	GetAllByCashManager(ctx context.Context, username string) ([]*domain.ReplenishmentRequest, error)
}

type CashManagerStorage interface {
	GetByShogunUsername(ctx context.Context, username string) (domain.CashManager, error)
}

type DaimyoStorage interface {
	GetByUsername(ctx context.Context, username string) (domain.Daimyo, error)
}

type CardStorage interface {
	GetByName(ctx context.Context, name string) (domain.Card, error)
	GetById(ctx context.Context, cardId int) (domain.Card, error)
}

type ReplenishmentRequestStatusStorage interface {
	GetId(ctx context.Context, name string) (int, error)
	GetById(ctx context.Context, statusId int) (string, error)
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
	daimyo, err := s.daimyoStorage.GetByUsername(ctx, dto.OwnerUsername)
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
		OwnerUsername:       dto.OwnerUsername,
		CardId:              card.CardId,
		Amount:              dto.Amount,
		StatusId:            statusId,
	}

	if err := s.storage.Insert(ctx, replenishmentReq); err != nil {
		return 0, err
	}

	return tg.ChatID(*cashManager.ChatId), nil
}

func (s *ReplenishmentRequestService) CheckIfExists(ctx context.Context, cardName string) (bool, error) {
	return s.storage.CheckIfExists(ctx, cardName)
}

func (s *ReplenishmentRequestService) GetAllByCashManager(
	ctx context.Context,
	username string,
) ([]*domain.ReplenishmentRequestDTO, error) {
	requests, err := s.storage.GetAllByCashManager(ctx, username)
	if err != nil {
		return nil, err
	}

	requestsDTOs := make([]*domain.ReplenishmentRequestDTO, 0)

	for _, request := range requests {
		card, err := s.cardStorage.GetById(ctx, request.CardId)
		if err != nil {
			return nil, err
		}

		status, err := s.replenishmentRequestStatusStorage.GetById(ctx, request.ReplenishmentRequestId)
		if err != nil {
			return nil, err
		}

		requestDTO := domain.ReplenishmentRequestDTO{
			CashManagerUsername: request.CashManagerUsername,
			OwnerUsername:       request.OwnerUsername,
			CardName:            card.Name,
			Amount:              request.Amount,
			Status:              status,
		}

		requestsDTOs = append(requestsDTOs, &requestDTO)
	}

	return requestsDTOs, nil
}
