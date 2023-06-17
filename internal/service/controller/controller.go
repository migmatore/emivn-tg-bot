package controller

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"time"
)

type ControllerStorage interface {
	Insert(ctx context.Context, cotroller domain.Controller) error
}

type ControllerTurnoverStorage interface {
	Insert(ctx context.Context, turnover domain.ControllerTurnover) error
	CheckIfExists(ctx context.Context, args domain.ControllerArgs) (bool, error)
	GetByDateAndBank(ctx context.Context, args domain.ControllerArgs) (domain.ControllerTurnover, error)
	Update(ctx context.Context, turnover domain.ControllerTurnover) error

	GetTurnover(ctx context.Context, samuraiUsername string, date string, bankId int) (float64, error)
}

type CardStorage interface {
	GetBankIdByName(ctx context.Context, bankName string) (int, error)
}

type UserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type ControllerService struct {
	transactor storage.Transactor

	storage                   ControllerStorage
	controllerTurnoverStorage ControllerTurnoverStorage
	cardStorage               CardStorage
	userRoleStorage           UserRoleStorage
	roleStorage               RoleStorage
}

func New(
	transactor storage.Transactor,
	storage ControllerStorage,
	turnoverStorage ControllerTurnoverStorage,
	cardStorage CardStorage,
	userRoleStorage UserRoleStorage,
	roleStorage RoleStorage,
) *ControllerService {
	return &ControllerService{
		transactor:                transactor,
		storage:                   storage,
		controllerTurnoverStorage: turnoverStorage,
		cardStorage:               cardStorage,
		userRoleStorage:           userRoleStorage,
		roleStorage:               roleStorage,
	}
}

func (s *ControllerService) Create(ctx context.Context, dto domain.ControllerDTO) error {
	controller := domain.Controller{
		Username: dto.Username,
		Nickname: dto.Nickname,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.ControllerRole.String())
	if err != nil {
		return err
	}

	userRole := domain.UserRole{
		Username: dto.Username,
		RoleId:   roleId,
	}

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.userRoleStorage.Insert(txCtx, userRole); err != nil {
			return err
		}

		if err := s.storage.Insert(txCtx, controller); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
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
