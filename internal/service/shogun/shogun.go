package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type ShogunStorage interface {
	Insert(ctx context.Context, shogun domain.Shogun) error
	GetAll(ctx context.Context) ([]*domain.Shogun, error)
	GetByNickname(ctx context.Context, nickname string) (domain.Shogun, error)
	GetByUsername(ctx context.Context, username string) (domain.Shogun, error)
	UpdateUsername(ctx context.Context, old string, new string) error
}

type ShogunUserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type ShogunRoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type ShogunService struct {
	transactor storage.Transactor

	storage         ShogunStorage
	userRoleStorage ShogunUserRoleStorage
	roleStorage     ShogunRoleStorage
}

func NewShogunService(
	t storage.Transactor,
	s ShogunStorage,
	userRole ShogunUserRoleStorage,
	role ShogunRoleStorage,
) *ShogunService {
	return &ShogunService{
		transactor:      t,
		storage:         s,
		userRoleStorage: userRole,
		roleStorage:     role,
	}
}

func (s *ShogunService) Create(ctx context.Context, dto domain.ShogunDTO) error {
	shogun := domain.Shogun{
		Username: dto.Username,
		Nickname: dto.Nickname,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.ShogunRole.String())
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

		if err := s.storage.Insert(txCtx, shogun); err != nil {
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

func (s *ShogunService) GetByNickname(ctx context.Context, nickname string) (domain.ShogunDTO, error) {
	shogun, err := s.storage.GetByNickname(ctx, nickname)
	if err != nil {
		return domain.ShogunDTO{}, err
	}

	shogunDTO := domain.ShogunDTO{
		Username: shogun.Username,
		Nickname: shogun.Nickname,
	}

	return shogunDTO, nil
}
