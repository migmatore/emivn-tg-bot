package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/daimyo"
	"emivn-tg-bot/internal/service/role"
	"emivn-tg-bot/internal/service/shogun"
	"emivn-tg-bot/internal/service/user_role"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	AuthStorage     auth.AuthStorage
	ShogunStorage   shogun.ShogunStorage
	DaimyoStorage   daimyo.DaimyoStorage
	UserRoleStorage user_role.UserRoleStorage
	RoleStorage     role.RoleStorage
}

type Service struct {
	Auth   *auth.AuthService
	Shogun *shogun.ShogunService
	Daimyo *daimyo.DaimyoService
}

func New(deps Deps) *Service {
	return &Service{
		Auth:   auth.NewAuthService(deps.AuthStorage),
		Shogun: shogun.NewShogunService(deps.Transactor, deps.ShogunStorage, deps.UserRoleStorage, deps.RoleStorage),
		Daimyo: daimyo.NewDaimyoService(
			deps.Transactor,
			deps.DaimyoStorage,
			deps.ShogunStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
	}
}
