package samurai_turnover

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"time"
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

func (s *SamuraiTurnoverStorage) CheckIfExists(ctx context.Context, time time.Time) (bool, error) {

	return true, nil
}
