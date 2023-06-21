package samurai_turnover

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type SamuraiTurnoverStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *SamuraiTurnoverStorage {
	return &SamuraiTurnoverStorage{pool: pool}
}

func (s *SamuraiTurnoverStorage) Insert(ctx context.Context, turnover domain.SamuraiTurnover) error {
	q := `insert into samurai_turnovers(samurai_username, start_date, initial_amount, final_amount, 
                              turnover, bank_type_id) values ($1, $2, $3, $4, $5, $6)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		turnover.SamuraiUsername,
		turnover.StartDate,
		turnover.InitialAmount,
		turnover.FinalAmount,
		turnover.Turnover,
		turnover.BankTypeId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil

	return nil
}

func (s *SamuraiTurnoverStorage) CheckIfExists(
	ctx context.Context,
	date string,
	bankId int,
	samuraiUsername string,
) (bool, error) {
	q := `select exists(select * from samurai_turnovers where start_date=$1 and bank_type_id=$2 and samurai_username=$3)`

	var exists bool

	if err := s.pool.QueryRow(ctx, q, date, bankId, samuraiUsername).Scan(&exists); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return false, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return false, err
	}

	return exists, nil
}

func (s *SamuraiTurnoverStorage) GetByDateAndBank(
	ctx context.Context,
	date string,
	bankId int,
	samuraiUsername string,
) (domain.SamuraiTurnover, error) {
	q := `select id, samurai_username, start_date::text, initial_amount, final_amount, turnover, bank_type_id 
				from samurai_turnovers where start_date = $1 and bank_type_id = $2 and samurai_username = $3`

	turnover := domain.SamuraiTurnover{}

	if err := s.pool.QueryRow(ctx, q, date, bankId, samuraiUsername).Scan(
		&turnover.TurnoverId,
		&turnover.SamuraiUsername,
		&turnover.StartDate,
		&turnover.InitialAmount,
		&turnover.FinalAmount,
		&turnover.Turnover,
		&turnover.BankTypeId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return turnover, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return turnover, err
	}

	return turnover, nil
}

func (s *SamuraiTurnoverStorage) Update(ctx context.Context, turnover domain.SamuraiTurnover) error {
	q := `update samurai_turnovers SET initial_amount=$1, final_amount=$2, turnover=$3 where samurai_username=$4 
                                                                               and bank_type_id=$5 and (id=$6 or start_date=$7)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		turnover.InitialAmount,
		turnover.FinalAmount,
		turnover.Turnover,
		turnover.SamuraiUsername,
		turnover.BankTypeId,
		turnover.TurnoverId,
		turnover.StartDate,
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

func (s *SamuraiTurnoverStorage) GetTurnover(
	ctx context.Context,
	samuraiUsername string,
	date string,
	bankId int,
) (float64, error) {
	q := `select turnover from samurai_turnovers where samurai_username=$1 and start_date=$2 and bank_type_id=$3`

	var turnover float64

	if err := s.pool.QueryRow(ctx, q, samuraiUsername, date, bankId).Scan(&turnover); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return turnover, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return turnover, err
	}

	return turnover, nil
}
