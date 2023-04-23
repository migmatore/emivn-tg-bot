package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type DaimyoStorage struct {
	pool psql.AtomicPoolClient
}

func NewDaimyoStorage(pool psql.AtomicPoolClient) *DaimyoStorage {
	return &DaimyoStorage{pool: pool}
}

func (s *DaimyoStorage) Insert(ctx context.Context, daimyo domain.Daimyo) error {
	q := `INSERT INTO daimyo(username, nickname, shogun_id) VALUES ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, daimyo.Username, daimyo.Nickname, daimyo.ShogunId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
