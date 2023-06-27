package referal

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type ReferalStorage struct {
	pool psql.AtomicPoolClient
}

func New(pool psql.AtomicPoolClient) *ReferalStorage {
	return &ReferalStorage{pool: pool}
}

func (s *ReferalStorage) Insert(ctx context.Context, link string, roleId int) error {
	q := `insert into referals(link, role_id) values ($1, $2)`

	if _, err := s.pool.Exec(ctx, q, link, roleId); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *ReferalStorage) GetByLink(ctx context.Context, link string) (domain.Referal, error) {
	q := `select id, link, role_id from referals where link=$1`

	referal := domain.Referal{}

	if err := s.pool.QueryRow(ctx, q, link).Scan(
		&referal.ReferalId,
		&referal.Link,
		&referal.RoleId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return referal, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return referal, err
	}

	return referal, nil
}

func (s *ReferalStorage) Delete(ctx context.Context, link string) error {
	q := `delete from referals where link = $1`

	if _, err := s.pool.Exec(ctx, q, link); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
