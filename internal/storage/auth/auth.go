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

func (s *AuthStorage) UserRole(ctx context.Context, username string) (string, error) {
	q := `select name from roles, user_roles where username = $1 and user_roles.role_id = roles.id`

	var role string

	if err := s.pool.QueryRow(ctx, q, username).Scan(&role); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return "", err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return "", err
	}

	return role, nil
}
