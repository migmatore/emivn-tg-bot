package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"emivn-tg-bot/pkg/logging"
	"github.com/mr-linch/go-tg"
)

type SamuraiStorage interface {
	Insert(ctx context.Context, samurai domain.Samurai) error
	GetByUsername(ctx context.Context, username string) (domain.Samurai, error)
	SetChatId(ctx context.Context, username string, id int64) error
}

type SamuraiUserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type SamuraiRoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type SamuraiService struct {
	transactor storage.Transactor

	storage         SamuraiStorage
	userRoleStorage SamuraiUserRoleStorage
	roleStorage     SamuraiRoleStorage
}

func NewSamuraiService(
	transactor storage.Transactor,
	s SamuraiStorage,
	userRole SamuraiUserRoleStorage,
	role SamuraiRoleStorage,
) *SamuraiService {
	return &SamuraiService{
		transactor:      transactor,
		storage:         s,
		userRoleStorage: userRole,
		roleStorage:     role,
	}
}

func (s *SamuraiService) Create(ctx context.Context, dto domain.SamuraiDTO) error {
	samurai := domain.Samurai{
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		DaimyoUsername: dto.DaimyoUsername,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.SamuraiRole.String())
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

		if err := s.storage.Insert(txCtx, samurai); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *SamuraiService) SetChatId(ctx context.Context, username string, id tg.ChatID) error {
	samurai, err := s.storage.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	// TODO: Error refactoring
	if samurai.ChatId != nil {
		logging.GetLogger(ctx).Error("Error: samurai chat id is already set")
	}

	return s.storage.SetChatId(ctx, username, int64(id))
}
