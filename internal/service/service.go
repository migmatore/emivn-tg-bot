package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/db_actions"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	DbActionsStorage db_actions.DbActionsStorage
	AuthStorage      auth.AuthStorage
}

type Service struct {
	DbActions *db_actions.DbActionsService
	Auth      *auth.AuthService
}

func New(deps Deps) *Service {
	return &Service{
		DbActions: db_actions.NewDbActionsService(deps.DbActionsStorage),
		Auth:      auth.NewAuthService(deps.AuthStorage),
	}
}
