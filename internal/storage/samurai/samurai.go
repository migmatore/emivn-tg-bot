package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type SamuraiStorage struct {
	pool psql.AtomicPoolClient
}

func NewSamuraiStorage(pool psql.AtomicPoolClient) *SamuraiStorage {
	return &SamuraiStorage{pool: pool}
}

func (s *SamuraiStorage) Insert(ctx context.Context, samurai domain.Samurai) error {
	q := `INSERT INTO samurai(username, nickname, daimyo_id) VALUES ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, samurai.Username, samurai.Nickname, samurai.DaimyoId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
