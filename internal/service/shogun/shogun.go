package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type ShogunStorage interface {
	Insert(ctx context.Context, shogun domain.Shogun) error
	GetAll(ctx context.Context) ([]*domain.Shogun, error)
}

type ShogunUserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type ShogunRoleStorage interface {
	GetByName(ctx context.Context, role string) (int, error)
}

type ShogunService struct {
	transactor storage.Transactor

	storage  ShogunStorage
	userRole ShogunUserRoleStorage
	role     ShogunRoleStorage
}

func NewShogunService(
	t storage.Transactor,
	s ShogunStorage,
	userRole ShogunUserRoleStorage,
	role ShogunRoleStorage,
) *ShogunService {
	return &ShogunService{
		transactor: t,
		storage:    s,
		userRole:   userRole,
		role:       role,
	}
}

func (s *ShogunService) Create(ctx context.Context, dto domain.ShogunDTO) error {
	shogun := domain.Shogun{
		Username: dto.Username,
		Nickname: dto.Nickname,
	}

	roleId, err := s.role.GetByName(ctx, domain.ShogunRole.String())
	if err != nil {
		return err
	}

	userRole := domain.UserRole{
		Username: dto.Username,
		RoleId:   roleId,
	}

	if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := s.userRole.Insert(ctx, userRole); err != nil {
			return err
		}

		if err := s.storage.Insert(ctx, shogun); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *ShogunService) GetAll(ctx context.Context) ([]*domain.ShogunDTO, error) {
	shoguns, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	shogunsDTOs := make([]*domain.ShogunDTO, 0)

	for _, item := range shoguns {
		shogunDto := domain.ShogunDTO{
			Username: item.Username,
			Nickname: item.Nickname,
		}

		shogunsDTOs = append(shogunsDTOs, &shogunDto)
	}

	return shogunsDTOs, nil
}
