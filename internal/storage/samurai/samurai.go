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
	q := `insert into samurai(username, nickname, daimyo_username) values ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, samurai.Username, samurai.Nickname, samurai.DaimyoUsername); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
