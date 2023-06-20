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
	q := `insert into replenishment_requests(cash_manager_username, owner_username, card_id, amount, status_id) 
		values ($1, $2, $3, $4, $5)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		replenishmentReq.CashManagerUsername,
		replenishmentReq.OwnerUsername,
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

func (s *ReplenishmentRequestStorage) CheckIfExists(ctx context.Context, cardName string) (bool, error) {
	q := `select exists(select * from replenishment_requests 
                       where card_id = (select id from cards where name=$1 limit 1))`

	var exists bool

	if err := s.pool.QueryRow(ctx, q, cardName).Scan(&exists); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return false, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return false, err
	}

	return exists, nil
}

func (s *ReplenishmentRequestStorage) GetAllByCashManager(
	ctx context.Context,
	username string,
	status string,
) ([]*domain.ReplenishmentRequest, error) {
	q := `select id, cash_manager_username, owner_username, card_id, amount, status_id from replenishment_requests
               where cash_manager_username = $1 
               and status_id = (select id from replenishment_request_status_groups where name = $2)`

	replenishmentRequests := make([]*domain.ReplenishmentRequest, 0)

	rows, err := s.pool.Query(ctx, q, username, status)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var replenishmentRequest domain.ReplenishmentRequest

		err := rows.Scan(
			&replenishmentRequest.ReplenishmentRequestId,
			&replenishmentRequest.CashManagerUsername,
			&replenishmentRequest.OwnerUsername,
			&replenishmentRequest.CardId,
			&replenishmentRequest.Amount,
			&replenishmentRequest.StatusId,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		replenishmentRequests = append(replenishmentRequests, &replenishmentRequest)
	}

	return replenishmentRequests, nil
}
