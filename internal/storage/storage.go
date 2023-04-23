package storage

import (
	"emivn-tg-bot/internal/storage/auth"
	"emivn-tg-bot/internal/storage/daimyo"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/internal/storage/role"
	"emivn-tg-bot/internal/storage/shogun"
	"emivn-tg-bot/internal/storage/user_role"
)

type Storage struct {
	Transactor *Transact

	Auth     *auth.AuthStorage
	Shogun   *shogun.ShogunStorage
	Daimyo   *daimyo.DaimyoStorage
	UserRole *user_role.UserRoleStorage
	Role     *role.RoleStorage
}

func New(pool psql.AtomicPoolClient) *Storage {
	return &Storage{
		Transactor: NewTransactor(pool),
		Auth:       auth.NewAuthStorage(pool),
		Shogun:     shogun.NewShogunStorage(pool),
		Daimyo:     daimyo.NewDaimyoStorage(pool),
		UserRole:   user_role.NewUserRoleStorage(pool),
		Role:       role.NewRoleStorage(pool),
	}
}
