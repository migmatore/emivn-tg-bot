package user_role

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type UserRoleStorage struct {
	pool psql.AtomicPoolClient
}

func NewUserRoleStorage(pool psql.AtomicPoolClient) *UserRoleStorage {
	return &UserRoleStorage{pool: pool}
}

func (s *UserRoleStorage) Insert(ctx context.Context, user domain.UserRole) error {
	q := `INSERT INTO user_roles(username, role_id) VALUES ($1, $2)`

	if _, err := s.pool.Exec(ctx, q, user.Username, user.RoleId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
