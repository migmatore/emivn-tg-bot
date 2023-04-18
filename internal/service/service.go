package service

import (
	"emivn-tg-bot/internal/service/db_actions"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	DbActionsStorage db_actions.DbActionsStorage
}

type Service struct {
	DbActions *db_actions.DbActionsService
}

func New(deps Deps) *Service {
	return &Service{DbActions: db_actions.NewDbActionsService(deps.DbActionsStorage)}
}
