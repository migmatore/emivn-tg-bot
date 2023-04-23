package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
)

type DaimyoStorage interface {
	Insert(ctx context.Context, daimyo domain.Daimyo) error
	GetAll(ctx context.Context) ([]*domain.Daimyo, error)
	GetIdByName(ctx context.Context, username string) (int, error)
}

type DaimyoShogunStorage interface {
	GetIdByName(ctx context.Context, username string) (int, error)
}

type DaimyoUserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type DaimyoRoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type DaimyoService struct {
	transactor storage.Transactor

	storage         DaimyoStorage
	shogunStorage   DaimyoShogunStorage
	userRoleStorage DaimyoUserRoleStorage
	roleStorage     DaimyoRoleStorage
}

func NewDaimyoService(
	t storage.Transactor,
	s DaimyoStorage,
	shogunStorage DaimyoShogunStorage,
	userRoleStorage DaimyoUserRoleStorage,
	roleStorage DaimyoRoleStorage,
) *DaimyoService {
	return &DaimyoService{
		transactor:      t,
		storage:         s,
		shogunStorage:   shogunStorage,
		userRoleStorage: userRoleStorage,
		roleStorage:     roleStorage,
	}
}

func (s *DaimyoService) Create(ctx context.Context, dto domain.DaimyoDTO) error {
	shogunId, err := s.shogunStorage.GetIdByName(ctx, dto.ShogunUsername)
	if err != nil {
		return err
	}

	daimyo := domain.Daimyo{
		Username: dto.Username,
		Nickname: dto.Nickname,
		ShogunId: shogunId,
	}

	roleId, err := s.roleStorage.GetIdByName(ctx, domain.DaimyoRole.String())
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

		if err := s.storage.Insert(txCtx, daimyo); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *DaimyoService) GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error) {
	daimyos, err := s.storage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	daimyoDTOs := make([]*domain.DaimyoDTO, 0)

	for _, item := range daimyos {
		daimyoDTO := domain.DaimyoDTO{
			Username:     item.Username,
			Nickname:     item.Nickname,
			CardsBalance: item.CardsBalance,
			// ShogunUsername: TODO: covert shogunId to shogunUsername
		}

		daimyoDTOs = append(daimyoDTOs, &daimyoDTO)
	}

	return daimyoDTOs, nil
}
