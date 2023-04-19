package auth

import (
	"context"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type AuthStorage struct {
	pool psql.AtomicPoolClient
}

func NewAuthStorage(pool psql.AtomicPoolClient) *AuthStorage {
	return &AuthStorage{pool: pool}
}

func (s *AuthStorage) CheckAuth(ctx context.Context, username string) (bool, error) {
	q := `select exists(select * from user_roles where username = $1)`

	var isExist bool

	if err := s.pool.QueryRow(ctx, q, username).Scan(&isExist); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return false, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return false, err
	}

	return isExist, nil
}
