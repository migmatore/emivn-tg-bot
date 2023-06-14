package samurai_turnover

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
)

type SamuraiTurnoverStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *SamuraiTurnoverStorage {
	return &SamuraiTurnoverStorage{pool: pool}
}

func (s *SamuraiTurnoverStorage) Insert(ctx context.Context, turnover domain.SamuraiTurnover) error {
	return nil
}
