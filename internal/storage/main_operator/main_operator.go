package main_operator

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type MainOperatorStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *MainOperatorStorage {
	return &MainOperatorStorage{pool: pool}
}

func (s *MainOperatorStorage) Insert(ctx context.Context, operator domain.MainOperator) error {
	q := `insert into main_operators(username, nickname, shogun_username) values ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, operator.Username, operator.Nickname, operator.ShogunUsername); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s MainOperatorStorage) GetByUsername(ctx context.Context, username string) (domain.MainOperator, error) {
	q := `select username, nickname, shogun_username from main_operators where username=$1`

	operator := domain.MainOperator{}

	if err := s.pool.QueryRow(ctx, q, username).Scan(
		&operator.Username,
		&operator.Nickname,
		&operator.ShogunUsername,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return operator, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return operator, err
	}

	return operator, nil
}
