package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ReplenishmentRequestStorage struct {
	pool psql.AtomicPoolClient
}

func NewReplenishmentRequestStorage(pool psql.AtomicPoolClient) *ReplenishmentRequestStorage {
	return &ReplenishmentRequestStorage{pool: pool}
}

func (s *ReplenishmentRequestStorage) Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error {
	q := `insert into replenishment_requests(cash_manager_username, daimyo_username, card_id, amount, status_id) 
		values ($1, $2, $3, $4, $5)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		replenishmentReq.CashManagerUsername,
		replenishmentReq.DaimyoUsername,
		replenishmentReq.CardId,
		replenishmentReq.Amount,
		replenishmentReq.StatusId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
