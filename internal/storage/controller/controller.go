package controller

import "emivn-tg-bot/internal/storage/psql"

type ControllerStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *ControllerStorage {
	return &ControllerStorage{pool: pool}
}
