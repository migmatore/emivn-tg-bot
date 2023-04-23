package user_role

import (
	"context"
	"emivn-tg-bot/internal/domain"
)

type UserRoleStorage interface {
	Insert(ctx context.Context, user domain.UserRole) error
}

type UserRoleSerivce struct {
	storage UserRoleStorage
}

func NewUserRoleServie(s UserRoleStorage) *UserRoleSerivce {
	return &UserRoleSerivce{storage: s}
}
