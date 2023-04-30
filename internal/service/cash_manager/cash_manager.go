package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type CashManagerStorage interface {
	Insert(ctx context.Context, cashManager domain.CashManager) error
}

type CashManagerUserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type CashManagerRoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type CashManagerService struct {
	transactor storage.Transactor

	storage         CashManagerStorage
	userRoleStorage CashManagerUserRoleStorage
	roleStorage     CashManagerRoleStorage
}

func NewCashManagerService(
	t storage.Transactor,
	s CashManagerStorage,
	userRole CashManagerUserRoleStorage,
	role CashManagerRoleStorage,
) *CashManagerService {
	return &CashManagerService{
		transactor:      t,
		storage:         s,
		userRoleStorage: userRole,
		roleStorage:     role,
	}
}

func (s *CashManagerService) Create(ctx context.Context, dto domain.CashManagerDTO) error {
	cashManager := domain.CashManager{
		Username: dto.Username,
		Nickname: dto.Nickname,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.CashManagerRole.String())
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

		if err := s.storage.Insert(txCtx, cashManager); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
