package db_actions

import (
	"context"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type DbActionsStorage struct {
	pool psql.AtomicPoolClient
}

func NewDbActionsStorage(pool psql.AtomicPoolClient) *DbActionsStorage {
	return &DbActionsStorage{pool: pool}
}

func (s *DbActionsStorage) Read(ctx context.Context) ([]string, error) {
	//q := `select txt from test where test.id=1`
	//
	//var txt string
	//
	//if err := s.pool.QueryRow(ctx, q).Scan(&txt); err != nil {
	//	if err := utils.ParsePgError(err); err != nil {
	//		logging.GetLogger(ctx).Errorf("Error: %v", err)
	//		return "", err
	//	}
	//
	//	logging.GetLogger(ctx).Errorf("Query error. %v", err)
	//	return "", err
	//}
	//
	//return txt, nil

	q := `select txt from test`

	strs := make([]string, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var str string

		err := rows.Scan(&str)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		strs = append(strs, str)
	}

	return strs, nil
}

func (s *DbActionsStorage) Write(ctx context.Context, text string) error {
	q := `INSERT INTO test(txt) VALUES ($1)`

	var id int

	if err := s.pool.QueryRow(
		ctx,
		q,
		text,
	).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
