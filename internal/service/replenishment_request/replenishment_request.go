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
	GetAllByOwner(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequest, error)
	GetByCardId(ctx context.Context, cardId int) (domain.ReplenishmentRequest, error)
	UpdateStatus(ctx context.Context, cardId int, statusId int) error
	UpdateActualAmount(ctx context.Context, cardId int, amount float32) error
	UpdateRequiredAmount(ctx context.Context, cardId int, amount float32) error
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

type MainOperatorStorage interface {
	GetByUsername(ctx context.Context, username string) (domain.MainOperator, error)
}

type ReplenishmentRequestService struct {
	transactor storage.Transactor

	storage                           ReplenishmentRequestStorage
	cashManagerStorage                CashManagerStorage
	daimyoStorage                     DaimyoStorage
	cardStorage                       CardStorage
	replenishmentRequestStatusStorage ReplenishmentRequestStatusStorage
	mainOperatorStorage               MainOperatorStorage
}

func NewReplenishmentRequestService(
	transactor storage.Transactor,
	storage ReplenishmentRequestStorage,
	cashManager CashManagerStorage,
	daimyo DaimyoStorage,
	card CardStorage,
	replenishmentRequestStatus ReplenishmentRequestStatusStorage,
	mainOperatorStorage MainOperatorStorage,
) *ReplenishmentRequestService {
	return &ReplenishmentRequestService{
		transactor:                        transactor,
		storage:                           storage,
		cashManagerStorage:                cashManager,
		daimyoStorage:                     daimyo,
		cardStorage:                       card,
		replenishmentRequestStatusStorage: replenishmentRequestStatus,
		mainOperatorStorage:               mainOperatorStorage,
	}
}

func (s *ReplenishmentRequestService) Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error) {
	var shogunUsername string

	daimyo, _ := s.daimyoStorage.GetByUsername(ctx, dto.OwnerUsername)
	if daimyo.Username == "" {
		operator, _ := s.mainOperatorStorage.GetByUsername(ctx, dto.OwnerUsername)
		shogunUsername = operator.ShogunUsername
	} else {
		shogunUsername = daimyo.ShogunUsername
	}

	cashManager, err := s.cashManagerStorage.GetByShogunUsername(ctx, shogunUsername)
	if err != nil {
		return 0, err
	}

	card, err := s.cardStorage.GetByName(ctx, dto.CardName)
	if err != nil {
		return 0, err
	}

	statusId, err := s.replenishmentRequestStatusStorage.GetId(ctx, domain.ActiveRequests.String())
	if err != nil {
		return 0, err
	}

	replenishmentReq := domain.ReplenishmentRequest{
		CashManagerUsername: cashManager.Username,
		OwnerUsername:       dto.OwnerUsername,
		CardId:              card.CardId,
		RequiredAmount:      dto.RequiredAmount,
		ActualAmount:        0,
		StatusId:            statusId,
	}

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.storage.Insert(txCtx, replenishmentReq); err != nil {
			return err
		}

		newLimit := card.DailyLimit - int(replenishmentReq.RequiredAmount)

		if err := s.cardStorage.UpdateLimit(txCtx, card.Name, newLimit); err != nil {
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
			RequiredAmount:      request.RequiredAmount,
			ActualAmount:        request.ActualAmount,
			Status:              status,
		}

		requestsDTOs = append(requestsDTOs, &requestDTO)
	}

	return requestsDTOs, nil
}

func (s *ReplenishmentRequestService) GetAllByOwner(
	ctx context.Context,
	username string,
	status string,
) ([]*domain.ReplenishmentRequestDTO, error) {
	requests, err := s.storage.GetAllByOwner(ctx, username, status)
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
			RequiredAmount:      request.RequiredAmount,
			ActualAmount:        request.ActualAmount,
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
		RequiredAmount:      request.RequiredAmount,
		ActualAmount:        request.ActualAmount,
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

func (s *ReplenishmentRequestService) ConfirmRequest(ctx context.Context, dto domain.ReplenishmentRequestDTO) error {
	card, err := s.cardStorage.GetByName(ctx, dto.CardName)
	if err != nil {
		return err
	}

	switch dto.Status {
	case domain.ActiveRequests.String():
		statusId, err := s.replenishmentRequestStatusStorage.GetId(ctx, domain.ObjectionableRequests.String())
		if err != nil {
			return err
		}

		oldRequest, err := s.storage.GetByCardId(ctx, card.CardId)
		if err != nil {
			return err
		}

		if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
			if oldRequest.RequiredAmount != dto.RequiredAmount {
				if err := s.storage.UpdateRequiredAmount(txCtx, card.CardId, dto.RequiredAmount); err != nil {
					return err
				}
			}

			if err := s.storage.UpdateStatus(txCtx, card.CardId, statusId); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return err
		}

	case domain.ObjectionableRequests.String():
		if dto.RequiredAmount != dto.ActualAmount {
			//oldRequest, err := s.storage.GetByCardId(ctx, card.CardId)
			//if err != nil {
			//	return err
			//}

			if err := s.storage.UpdateActualAmount(ctx, card.CardId, dto.ActualAmount); err != nil {
				return err
			}

			return nil
		}

		statusId, err := s.replenishmentRequestStatusStorage.GetId(ctx, domain.CompletedRequests.String())
		if err != nil {
			return err
		}

		oldRequest, err := s.storage.GetByCardId(ctx, card.CardId)
		if err != nil {
			return err
		}

		if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
			if oldRequest.RequiredAmount != dto.RequiredAmount {
				if err := s.storage.UpdateRequiredAmount(txCtx, card.CardId, dto.RequiredAmount); err != nil {
					return err
				}
			}

			if err := s.storage.UpdateActualAmount(txCtx, card.CardId, dto.ActualAmount); err != nil {
				return err
			}

			if err := s.storage.UpdateStatus(txCtx, card.CardId, statusId); err != nil {
				return err
			}

			return nil
		}); err != nil {
			return err
		}
		return nil
	}

	return nil
}
