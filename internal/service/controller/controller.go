package controller

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"time"
)

type ControllerStorage interface {
}

type ControllerTurnoverStorage interface {
	Insert(ctx context.Context, turnover domain.ControllerTurnover) error
	CheckIfExists(ctx context.Context, args domain.ControllerArgs) (bool, error)
	GetByDateAndBank(ctx context.Context, args domain.ControllerArgs) (domain.ControllerTurnover, error)
	Update(ctx context.Context, turnover domain.ControllerTurnover) error
}

type CardStorage interface {
	GetBankIdByName(ctx context.Context, bankName string) (int, error)
}

type ControllerService struct {
	transactor storage.Transactor

	storage                   ControllerStorage
	controllerTurnoverStorage ControllerTurnoverStorage
	cardStorage               CardStorage
}

func New(
	transactor storage.Transactor,
	storage ControllerStorage,
	turnoverStorage ControllerTurnoverStorage,
	cardStorage CardStorage,
) *ControllerService {
	return &ControllerService{
		transactor:                transactor,
		storage:                   storage,
		controllerTurnoverStorage: turnoverStorage,
		cardStorage:               cardStorage,
	}
}

func (s *ControllerService) CreateTurnover(ctx context.Context, dto domain.ControllerTurnoverDTO) error {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	bankId, err := s.cardStorage.GetBankIdByName(ctx, dto.BankTypeName)
	if err != nil {
		return err
	}

	exists, err := s.controllerTurnoverStorage.CheckIfExists(ctx, domain.ControllerArgs{
		Date:               yesterday,
		BankId:             bankId,
		ControllerUsername: dto.ControllerUsername,
		SamuraiUsername:    dto.SamuraiUsername,
	})
	if err != nil {
		return err
	}

	initialAmount := dto.FinalAmount

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if !exists {
			turnover := domain.ControllerTurnover{
				ControllerUsername: dto.ControllerUsername,
				SamuraiUsername:    dto.SamuraiUsername,
				StartDate:          yesterday,
				InitialAmount:      0,
				FinalAmount:        dto.FinalAmount,
				Turnover:           dto.FinalAmount,
				BankTypeId:         bankId,
			}

			if err := s.controllerTurnoverStorage.Insert(ctx, turnover); err != nil {
				return err
			}
		} else {
			turnover, err := s.controllerTurnoverStorage.GetByDateAndBank(ctx, domain.ControllerArgs{
				Date:               yesterday,
				BankId:             bankId,
				ControllerUsername: dto.ControllerUsername,
				SamuraiUsername:    dto.SamuraiUsername,
			})
			if err != nil {
				return err
			}

			turnover.FinalAmount = dto.FinalAmount + turnover.InitialAmount
			turnover.Turnover = dto.FinalAmount

			initialAmount = turnover.FinalAmount

			if err := s.controllerTurnoverStorage.Update(ctx, turnover); err != nil {
				return err
			}
		}

		turnover := domain.ControllerTurnover{
			ControllerUsername: dto.ControllerUsername,
			SamuraiUsername:    dto.SamuraiUsername,
			StartDate:          time.Now().Format("2006-01-02"),
			InitialAmount:      initialAmount,
			FinalAmount:        0,
			Turnover:           initialAmount,
			BankTypeId:         bankId,
		}

		newExists, err := s.controllerTurnoverStorage.CheckIfExists(ctx, domain.ControllerArgs{
			Date:               time.Now().Format("2006-01-02"),
			BankId:             bankId,
			ControllerUsername: dto.ControllerUsername,
			SamuraiUsername:    dto.SamuraiUsername,
		})
		if err != nil {
			return err
		}

		if !newExists {
			if err := s.controllerTurnoverStorage.Insert(ctx, turnover); err != nil {
				return err
			}

			return nil
		}

		if err := s.controllerTurnoverStorage.Update(ctx, turnover); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
