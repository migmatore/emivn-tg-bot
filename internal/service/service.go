package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/db_actions"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	DbActionsStorage db_actions.DbActionsStorage
}

type Service struct {
	DbActions   *db_actions.DbActionsService
	AuthService *auth.AuthService
}

func New(deps Deps) *Service {
	return &Service{
		DbActions:   db_actions.NewDbActionsService(deps.DbActionsStorage),
		AuthService: auth.NewAuthService(),
	}
}
