package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/shogun"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	AuthStorage     auth.AuthStorage
	ShogunStorage   shogun.ShogunStorage
	UserRoleStorage shogun.ShogunUserRoleStorage
	RoleStorage     shogun.ShogunRoleStorage
}

type Service struct {
	Auth   *auth.AuthService
	Shogun *shogun.ShogunService
}

func New(deps Deps) *Service {
	return &Service{
		Auth:   auth.NewAuthService(deps.AuthStorage),
		Shogun: shogun.NewShogunService(deps.Transactor, deps.ShogunStorage, deps.UserRoleStorage, deps.RoleStorage),
	}
}
