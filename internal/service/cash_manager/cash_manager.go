package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"emivn-tg-bot/pkg/logging"
	"github.com/mr-linch/go-tg"
)

type CashManagerStorage interface {
	Insert(ctx context.Context, cashManager domain.CashManager) error
	GetByShogunUsername(ctx context.Context, username string) (domain.CashManager, error)
	GetByUsername(ctx context.Context, username string) (domain.CashManager, error)
	SetChatId(ctx context.Context, username string, id int64) error
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
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		ShogunUsername: dto.ShogunUsername,
		ChatId:         dto.ChatId,
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

func (s *CashManagerService) SetChatId(ctx context.Context, username string, id tg.ChatID) error {
	cashManager, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	// TODO: Error refactoring
	if cashManager.ChatId != nil {
		logging.GetLogger(ctx).Error("Error: cashManager chat id is already set")
	}

	return s.storage.SetChatId(ctx, username, int64(id))
}
