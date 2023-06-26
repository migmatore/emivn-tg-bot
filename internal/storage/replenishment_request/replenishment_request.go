package replenishment_request

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"time"
)

type ReplenishmentRequestStorage struct {
	pool psql.AtomicPoolClient
}

func NewReplenishmentRequestStorage(pool psql.AtomicPoolClient) *ReplenishmentRequestStorage {
	return &ReplenishmentRequestStorage{pool: pool}
}

func (s *ReplenishmentRequestStorage) Insert(ctx context.Context, replenishmentReq domain.ReplenishmentRequest) error {
	q := `insert into replenishment_requests(cash_manager_username, owner_username, card_id, required_amount, 
                                   actual_amount, status_id, creation_date, creation_time) 
				values ($1, $2, $3, $4, $5, $6, $7, $8)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		replenishmentReq.CashManagerUsername,
		replenishmentReq.OwnerUsername,
		replenishmentReq.CardId,
		replenishmentReq.RequiredAmount,
		replenishmentReq.ActualAmount,
		replenishmentReq.StatusId,
		time.Now().Format("2006-01-02"),
		time.Now().Format("15:04:05"),
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
                       where card_id = (select id from cards where name=$1 limit 1) 
                       and status_id <> (select id from replenishment_request_status_groups where name = $2)
                       and creation_date = $3)`

	var exists bool

	if err := s.pool.QueryRow(
		ctx,
		q,
		cardName,
		domain.CompletedRequests,
		time.Now().Format("2006-01-02"),
	).Scan(&exists); err != nil {
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
	q := `select id, cash_manager_username, owner_username, card_id, required_amount, actual_amount, status_id, 
       					creation_date::text, creation_time::text from replenishment_requests 
       		   where cash_manager_username = $1 
               and status_id = (select id from replenishment_request_status_groups where name = $2)`

	requests := make([]*domain.ReplenishmentRequest, 0)

	rows, err := s.pool.Query(ctx, q, username, status)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var request domain.ReplenishmentRequest

		err := rows.Scan(
			&request.ReplenishmentRequestId,
			&request.CashManagerUsername,
			&request.OwnerUsername,
			&request.CardId,
			&request.RequiredAmount,
			&request.ActualAmount,
			&request.StatusId,
			&request.CreationDate,
			&request.CreationTime,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		requests = append(requests, &request)
	}

	return requests, nil
}

func (s *ReplenishmentRequestStorage) GetAllByOwner(
	ctx context.Context,
	username string,
	status string,
) ([]*domain.ReplenishmentRequest, error) {
	q := `select id, cash_manager_username, owner_username, card_id, required_amount, actual_amount, status_id, 
       					creation_date::text, creation_time::text from replenishment_requests 
       		   where owner_username = $1 
               and status_id = (select id from replenishment_request_status_groups where name = $2)`

	requests := make([]*domain.ReplenishmentRequest, 0)

	rows, err := s.pool.Query(ctx, q, username, status)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var request domain.ReplenishmentRequest

		err := rows.Scan(
			&request.ReplenishmentRequestId,
			&request.CashManagerUsername,
			&request.OwnerUsername,
			&request.CardId,
			&request.RequiredAmount,
			&request.ActualAmount,
			&request.StatusId,
			&request.CreationDate,
			&request.CreationTime,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		requests = append(requests, &request)
	}

	return requests, nil
}

func (s *ReplenishmentRequestStorage) GetByCardId(
	ctx context.Context,
	cardId int,
) (domain.ReplenishmentRequest, error) {
	q := `select id, cash_manager_username, owner_username, card_id, required_amount, actual_amount, status_id, 
       creation_date::text, creation_time::text from replenishment_requests where card_id = $1 
                    and status_id <> (select id from replenishment_request_status_groups where name = $2)`

	request := domain.ReplenishmentRequest{}

	if err := s.pool.QueryRow(ctx, q, cardId, domain.CompletedRequests.String()).Scan(
		&request.ReplenishmentRequestId,
		&request.CashManagerUsername,
		&request.OwnerUsername,
		&request.CardId,
		&request.RequiredAmount,
		&request.ActualAmount,
		&request.StatusId,
		&request.CreationDate,
		&request.CreationTime,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return request, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return request, err
	}

	return request, nil
}

func (s *ReplenishmentRequestStorage) UpdateStatus(ctx context.Context, cardId int, statusId int) error {
	q := `update replenishment_requests set status_id = $1 where card_id = $2 and creation_date = $3 
                            and status_id <> (select id from replenishment_request_status_groups where name = $4)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		statusId,
		cardId,
		time.Now().Format("2006-01-02"),
		domain.CompletedRequests.String(),
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

func (s *ReplenishmentRequestStorage) UpdateActualAmount(ctx context.Context, cardId int, amount float32) error {
	q := `update replenishment_requests set actual_amount = $1 where card_id = $2 and creation_date = $3 
                        and status_id <> (select id from replenishment_request_status_groups where name = $4)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		amount,
		cardId,
		time.Now().Format("2006-01-02"),
		domain.CompletedRequests.String(),
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

func (s *ReplenishmentRequestStorage) UpdateRequiredAmount(ctx context.Context, cardId int, amount float32) error {
	q := `update replenishment_requests set required_amount = $1 where card_id = $2 and creation_date = $3 
                        and status_id <> (select id from replenishment_request_status_groups where name = $4)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		amount,
		cardId,
		time.Now().Format("2006-01-02"),
		domain.CompletedRequests.String(),
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
