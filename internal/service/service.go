package service

import "emivn-tg-bot/internal/storage"

type Deps struct {
	Transactor storage.Transactor
}

type Service struct {
}

func New(deps Deps) *Service {
	return &Service{}
}
