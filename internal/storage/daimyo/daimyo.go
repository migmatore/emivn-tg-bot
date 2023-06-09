package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type DaimyoStorage struct {
	pool psql.AtomicPoolClient
}

func NewDaimyoStorage(pool psql.AtomicPoolClient) *DaimyoStorage {
	return &DaimyoStorage{pool: pool}
}

func (s *DaimyoStorage) Insert(ctx context.Context, daimyo domain.Daimyo) error {
	q := `insert into daimyo(username, nickname, shogun_username) values ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, daimyo.Username, daimyo.Nickname, daimyo.ShogunUsername); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *DaimyoStorage) GetAll(ctx context.Context) ([]*domain.Daimyo, error) {
	q := `select shogun_username, username, nickname from daimyo`

	daimyos := make([]*domain.Daimyo, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var daimyo domain.Daimyo

		err := rows.Scan(&daimyo.ShogunUsername, &daimyo.Username, &daimyo.Nickname)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		daimyos = append(daimyos, &daimyo)
	}

	return daimyos, nil
}

func (s *DaimyoStorage) GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.Daimyo, error) {
	q := `select shogun_username, username, nickname from daimyo where shogun_username=$1`

	daimyos := make([]*domain.Daimyo, 0)

	rows, err := s.pool.Query(ctx, q, shogunUsername)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var daimyo domain.Daimyo

		err := rows.Scan(&daimyo.ShogunUsername, &daimyo.Username, &daimyo.Nickname)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		daimyos = append(daimyos, &daimyo)
	}

	return daimyos, nil
}

func (s *DaimyoStorage) GetByUsername(ctx context.Context, username string) (domain.Daimyo, error) {
	q := `select username, nickname, cards_balance, shogun_username from daimyo where username=$1`

	daimyo := domain.Daimyo{}

	if err := s.pool.QueryRow(ctx, q, username).Scan(
		&daimyo.Username,
		&daimyo.Nickname,
		&daimyo.CardsBalance,
		&daimyo.ShogunUsername,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return daimyo, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return daimyo, err
	}

	return daimyo, nil
}

func (s *DaimyoStorage) GetByNickname(ctx context.Context, nickname string) (domain.Daimyo, error) {
	q := `select username, nickname, cards_balance, shogun_username from daimyo where nickname=$1`

	daimyo := domain.Daimyo{}

	if err := s.pool.QueryRow(ctx, q, nickname).Scan(
		&daimyo.Username,
		&daimyo.Nickname,
		&daimyo.CardsBalance,
		&daimyo.ShogunUsername,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return daimyo, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return daimyo, err
	}

	return daimyo, nil
}

func (s *DaimyoStorage) UpdateUsername(ctx context.Context, old string, new string) error {
	q := `update daimyo set username=$1 where username=$2`

	if _, err := s.pool.Exec(ctx, q, new, old); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
