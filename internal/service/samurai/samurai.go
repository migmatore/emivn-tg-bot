package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type SamuraiStorage interface {
	Insert(ctx context.Context, samurai domain.Samurai) error
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
