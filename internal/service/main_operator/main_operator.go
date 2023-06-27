package main_operator

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type MainOperatorStorage interface {
	Insert(ctx context.Context, operator domain.MainOperator) error
	GetByUsername(ctx context.Context, username string) (domain.MainOperator, error)
	UpdateUsername(ctx context.Context, old string, new string) error
}

type UserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type MainOperatorService struct {
	transactor storage.Transactor

	storage         MainOperatorStorage
	userRoleStorage UserRoleStorage
	roleStorage     RoleStorage
}

func New(
	transactor storage.Transactor,
	storage MainOperatorStorage,
	userRoleStorage UserRoleStorage,
	roleStorage RoleStorage,
) *MainOperatorService {
	return &MainOperatorService{
		transactor:      transactor,
		storage:         storage,
		userRoleStorage: userRoleStorage,
		roleStorage:     roleStorage,
	}
}

func (s *MainOperatorService) Create(ctx context.Context, dto domain.MainOperatorDTO) error {
	operator := domain.MainOperator{
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		ShogunUsername: dto.ShogunUsername,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.MainOperatorRole.String())
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

		if err := s.storage.Insert(txCtx, operator); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *MainOperatorService) GetByUsername(ctx context.Context, username string) (domain.MainOperatorDTO, error) {
	operator, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return domain.MainOperatorDTO{}, err
	}

	operatorDTO := domain.MainOperatorDTO{
		Username:       operator.Username,
		Nickname:       operator.Nickname,
		ShogunUsername: operator.ShogunUsername,
	}

	return operatorDTO, nil
}
