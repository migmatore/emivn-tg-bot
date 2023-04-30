package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CashManagerStorage struct {
	pool psql.AtomicPoolClient
}

func NewCashManagerStorage(pool psql.AtomicPoolClient) *CashManagerStorage {
	return &CashManagerStorage{pool: pool}
}

func (s *CashManagerStorage) Insert(ctx context.Context, cashManager domain.CashManager) error {
	q := `insert into cash_managers(username, nickname) values ($1, $2)`

	if _, err := s.pool.Exec(ctx, q, cashManager.Username, cashManager.Nickname); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
