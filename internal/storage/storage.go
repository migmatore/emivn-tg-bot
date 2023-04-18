package storage

import (
	"emivn-tg-bot/internal/storage/db_actions"
	"emivn-tg-bot/internal/storage/psql"
)

type Storage struct {
	Transactor *Transact
	DbActions  *db_actions.DbActionsStorage
}

func New(pool psql.AtomicPoolClient) *Storage {
	return &Storage{
		Transactor: NewTransactor(pool),
		DbActions:  db_actions.NewDbActionsStorage(pool),
	}
}
