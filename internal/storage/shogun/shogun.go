package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ShogunStorage struct {
	pool psql.AtomicPoolClient
}

func NewShogunStorage(pool psql.AtomicPoolClient) *ShogunStorage {
	return &ShogunStorage{pool: pool}
}

func (s *ShogunStorage) Insert(ctx context.Context, shogun domain.Shogun) error {
	q := `INSERT INTO shoguns(username, nickname) VALUES ($1, $2)`

	if _, err := s.pool.Exec(ctx, q, shogun.Username, shogun.Nickname); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *ShogunStorage) GetAll(ctx context.Context) ([]*domain.Shogun, error) {
	q := `select username, nickname from shoguns`

	shoguns := make([]*domain.Shogun, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var shogun domain.Shogun

		err := rows.Scan(&shogun.Username, &shogun.Nickname)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		shoguns = append(shoguns, &shogun)
	}

	return shoguns, nil
}

func (s *ShogunStorage) GetByUsername(ctx context.Context, username string) (domain.Shogun, error) {
	q := `select username, nickname from shoguns where username=$1`

	shogun := domain.Shogun{}

	if err := s.pool.QueryRow(ctx, q, username).Scan(&shogun.Username, &shogun.Nickname); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return shogun, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return shogun, err
	}

	return shogun, nil
}

func (s *ShogunStorage) GetByNickname(ctx context.Context, nickname string) (domain.Shogun, error) {
	q := `select username, nickname from shoguns where nickname=$1`

	shogun := domain.Shogun{}

	if err := s.pool.QueryRow(ctx, q, nickname).Scan(&shogun.Username, &shogun.Nickname); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return shogun, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return shogun, err
	}

	return shogun, nil
}

func (s *ShogunStorage) UpdateUsername(ctx context.Context, old string, new string) error {
	q := `update shoguns set username=$1 where username=$2`

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
