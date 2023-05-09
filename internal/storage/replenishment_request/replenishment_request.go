package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
)

type ReplenishmentRequestStorage struct {
	pool psql.AtomicPoolClient
}

func NewReplenishmentRequestStorage(pool psql.AtomicPoolClient) *ReplenishmentRequestStorage {
	return &ReplenishmentRequestStorage{pool: pool}
}

func (s *ReplenishmentRequestStorage) Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error {
	return nil
}
