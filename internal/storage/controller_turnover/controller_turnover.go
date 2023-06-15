package controller_turnover

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ControllerTurnoverStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *ControllerTurnoverStorage {
	return &ControllerTurnoverStorage{pool: pool}
}

func (s *ControllerTurnoverStorage) Insert(ctx context.Context, turnover domain.ControllerTurnover) error {
	q := `insert into controller_turnovers(controller_username, samurai_username, start_date, initial_amount, 
                                 final_amount, turnover, bank_type_id) values ($1, $2, $3, $4, $5, $6, $7)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		turnover.ControllerUsername,
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
}

func (s *ControllerTurnoverStorage) CheckIfExists(
	ctx context.Context,
	args domain.ControllerArgs,
) (bool, error) {
	q := `select exists(select * from controller_turnovers where start_date=$1 and bank_type_id=$2 
                                                   and controller_username=$3 and samurai_username=$4)`

	var exists bool

	if err := s.pool.QueryRow(ctx, q, args.Date, args.BankId, args.ControllerUsername, args.SamuraiUsername).Scan(&exists); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return false, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return false, err
	}

	return exists, nil
}

func (s *ControllerTurnoverStorage) GetByDateAndBank(
	ctx context.Context,
	args domain.ControllerArgs,
) (domain.ControllerTurnover, error) {
	q := `select id, controller_username, samurai_username, start_date::text, initial_amount, final_amount, turnover, 
       bank_type_id from controller_turnovers where start_date = $1 and bank_type_id = $2 and controller_username = $3
                                                and samurai_username = $4`

	turnover := domain.ControllerTurnover{}

	if err := s.pool.QueryRow(ctx, q, args.Date, args.BankId, args.ControllerUsername, args.SamuraiUsername).Scan(
		&turnover.TurnoverId,
		&turnover.ControllerUsername,
		&turnover.SamuraiUsername,
		&turnover.StartDate,
		&turnover.InitialAmount,
		&turnover.FinalAmount,
		&turnover.TurnoverId,
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

// TODO: Add a rewrite for identical samurai
func (s *ControllerTurnoverStorage) Update(ctx context.Context, turnover domain.ControllerTurnover) error {
	q := `update controller_turnovers SET initial_amount=$1, final_amount=$2, turnover=$3 where controller_username=$4
                                            and samurai_username=$5 and bank_type_id=$6 and (id=$7 or start_date=$8)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		turnover.InitialAmount,
		turnover.FinalAmount,
		turnover.Turnover,
		turnover.ControllerUsername,
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
