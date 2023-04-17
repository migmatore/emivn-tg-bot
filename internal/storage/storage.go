package storage

import "emivn-tg-bot/internal/storage/psql"

type Storage struct {
	Transactor *Transact
}

func New(pool psql.AtomicPoolClient) *Storage {
	return &Storage{
		Transactor: NewTransactor(pool),
	}
}
