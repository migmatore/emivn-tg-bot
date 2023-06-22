package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"github.com/mr-linch/go-tg"
)

type ReplenishmentRequestStorage interface {
	Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error
	CheckIfExists(ctx context.Context, cardName string) (bool, error)
	GetAllByCashManager(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequest, error)
	GetByCardId(ctx context.Context, cardId int) (domain.ReplenishmentRequest, error)
	UpdateStatus(ctx context.Context, cardId int, statusId int) error
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
	UpdateLimit(ctx context.Context, name string, limit int) error
}

type ReplenishmentRequestStatusStorage interface {
	GetId(ctx context.Context, name string) (int, error)
	GetById(ctx context.Context, statusId int) (string, error)
}

type ReplenishmentRequestService struct {
	transactor storage.Transactor

	storage                           ReplenishmentRequestStorage
	cashManagerStorage                CashManagerStorage
	daimyoStorage                     DaimyoStorage
	cardStorage                       CardStorage
	replenishmentRequestStatusStorage ReplenishmentRequestStatusStorage
}

func NewReplenishmentRequestService(
	transactor storage.Transactor,
	storage ReplenishmentRequestStorage,
	cashManager CashManagerStorage,
	daimyo DaimyoStorage,
	card CardStorage,
	replenishmentRequestStatus ReplenishmentRequestStatusStorage,
) *ReplenishmentRequestService {
	return &ReplenishmentRequestService{
		transactor:                        transactor,
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

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.storage.Insert(ctx, replenishmentReq); err != nil {
			return err
		}

		newLimit := card.DailyLimit - int(replenishmentReq.Amount)

		if err := s.cardStorage.UpdateLimit(ctx, card.Name, newLimit); err != nil {
			return err
		}

		return nil
	}); err != nil {
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
	status string,
) ([]*domain.ReplenishmentRequestDTO, error) {
	requests, err := s.storage.GetAllByCashManager(ctx, username, status)
	if err != nil {
		return nil, err
	}

	requestsDTOs := make([]*domain.ReplenishmentRequestDTO, 0)

	for _, request := range requests {
		card, err := s.cardStorage.GetById(ctx, request.CardId)
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

func (s *ReplenishmentRequestService) GetByCardName(
	ctx context.Context,
	name string,
) (domain.ReplenishmentRequestDTO, error) {
	card, err := s.cardStorage.GetByName(ctx, name)
	if err != nil {
		return domain.ReplenishmentRequestDTO{}, err
	}

	request, err := s.storage.GetByCardId(ctx, card.CardId)
	if err != nil {
		return domain.ReplenishmentRequestDTO{}, err
	}

	status, err := s.replenishmentRequestStatusStorage.GetById(ctx, request.StatusId)
	if err != nil {
		return domain.ReplenishmentRequestDTO{}, err
	}

	requestDTO := domain.ReplenishmentRequestDTO{
		CashManagerUsername: request.CashManagerUsername,
		OwnerUsername:       request.OwnerUsername,
		CardName:            name,
		Amount:              request.Amount,
		Status:              status,
	}

	return requestDTO, nil
}

func (s *ReplenishmentRequestService) ChangeStatus(ctx context.Context, cardName string, status string) error {
	statusId, err := s.replenishmentRequestStatusStorage.GetId(ctx, status)
	if err != nil {
		return err
	}

	card, err := s.cardStorage.GetByName(ctx, cardName)
	if err != nil {
		return err
	}

	return s.storage.UpdateStatus(ctx, card.CardId, statusId)
}
