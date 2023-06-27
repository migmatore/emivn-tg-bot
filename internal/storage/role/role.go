package role

import (
	"context"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type RoleStorage struct {
	pool psql.AtomicPoolClient
}

func NewRoleStorage(pool psql.AtomicPoolClient) *RoleStorage {
	return &RoleStorage{pool: pool}
}

func (s *RoleStorage) GetIdByName(ctx context.Context, role string) (int, error) {
	q := `select id from roles where name=$1`

	var id int

	if err := s.pool.QueryRow(ctx, q, role).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return 0, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return 0, err
	}

	return id, nil
}

func (s *RoleStorage) GetById(ctx context.Context, roleId int) (string, error) {
	q := `select name from roles where id=$1`

	var name string

	if err := s.pool.QueryRow(ctx, q, roleId).Scan(&name); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return name, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return name, err
	}

	return name, nil
}
