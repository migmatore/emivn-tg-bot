package referal

import (
	"context"
	"emivn-tg-bot/internal/domain"
)

type ReferalStorage interface {
	Insert(ctx context.Context, link string, roleId int) error
	GetByLink(ctx context.Context, link string) (domain.Referal, error)
	Delete(ctx context.Context, link string) error
}

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
}

type ReferalService struct {
	storage     ReferalStorage
	roleStorage RoleStorage
}

func New(storage ReferalStorage, roleStorage RoleStorage) *ReferalService {
	return &ReferalService{storage: storage, roleStorage: roleStorage}
}

func (s *ReferalService) Create(ctx context.Context, link string, role string) error {
	roleId, err := s.roleStorage.GetIdByName(ctx, role)
	if err != nil {
		return err
	}

	if err := s.storage.Insert(ctx, link, roleId); err != nil {
		return err
	}

	return nil
}
