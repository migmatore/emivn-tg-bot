package controller

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ControllerStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *ControllerStorage {
	return &ControllerStorage{pool: pool}
}

func (s *ControllerStorage) Insert(ctx context.Context, controller domain.Controller) error {
	q := `INSERT INTO controllers(username, nickname) VALUES ($1, $2)`

	if _, err := s.pool.Exec(ctx, q, controller.Username, controller.Nickname); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
