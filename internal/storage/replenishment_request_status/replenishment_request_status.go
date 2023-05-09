package replenishment_request_status

import (
	"context"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ReplenishmentRequestStatusStorage struct {
	pool psql.AtomicPoolClient
}

func NewReplenishmentRequestStatusStorage(pool psql.AtomicPoolClient) *ReplenishmentRequestStatusStorage {
	return &ReplenishmentRequestStatusStorage{
		pool: pool,
	}
}

func (s *ReplenishmentRequestStatusStorage) GetId(ctx context.Context, name string) (int, error) {
	q := `select id from replenishment_request_status where name=$1`

	var id int

	if err := s.pool.QueryRow(ctx, q, name).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return id, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return id, err
	}

	return id, nil
}
