package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"log"
	"time"
)

type DaimyoStorage interface {
	Insert(ctx context.Context, daimyo domain.Daimyo) error
	GetAll(ctx context.Context) ([]*domain.Daimyo, error)
	GetByUsername(ctx context.Context, username string) (domain.Daimyo, error)
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
	userRoleStorage DaimyoUserRoleStorage
	roleStorage     DaimyoRoleStorage
}

func NewDaimyoService(
	t storage.Transactor,
	s DaimyoStorage,
	userRoleStorage DaimyoUserRoleStorage,
	roleStorage DaimyoRoleStorage,
) *DaimyoService {
	return &DaimyoService{
		transactor:      t,
		storage:         s,
		userRoleStorage: userRoleStorage,
		roleStorage:     roleStorage,
	}
}

func (s *DaimyoService) Create(ctx context.Context, dto domain.DaimyoDTO) error {
	daimyo := domain.Daimyo{
		Username:       dto.Username,
		Nickname:       dto.Nickname,
		ShogunUsername: dto.ShogunUsername,
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
			Username:       item.Username,
			Nickname:       item.Nickname,
			CardsBalance:   item.CardsBalance,
			ShogunUsername: item.ShogunUsername,
		}

		daimyoDTOs = append(daimyoDTOs, &daimyoDTO)
	}

	return daimyoDTOs, nil
}

func (s *DaimyoService) Notify(args domain.FuncArgs) (status domain.TaskStatus, when interface{}) {
	if name, ok := args["name"]; ok {
		log.Println("PrintWithArgs:", time.Now(), name)
		return domain.TaskStatusDeferred, time.Now().Add(time.Second * 10)
	}

	log.Print("Not found name arg in func args")

	return domain.TaskStatusDeferred, time.Now().Add(time.Second * 10)
}
