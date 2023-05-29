package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CashManagerStorage struct {
	pool psql.AtomicPoolClient
}

func NewCashManagerStorage(pool psql.AtomicPoolClient) *CashManagerStorage {
	return &CashManagerStorage{pool: pool}
}

func (s *CashManagerStorage) Insert(ctx context.Context, cashManager domain.CashManager) error {
	q := `insert into cash_managers(username, nickname, shogun_username, chat_id) values ($1, $2, $3, $4)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		cashManager.Username,
		cashManager.Nickname,
		cashManager.ShogunUsername,
		cashManager.ChatId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *CashManagerStorage) GetByShogunUsername(ctx context.Context, username string) (domain.CashManager, error) {
	q := `select username, nickname, shogun_username, chat_id from cash_managers where shogun_username=$1`

	cashManager := domain.CashManager{}

	if err := s.pool.QueryRow(ctx, q, username).Scan(
		&cashManager.Username,
		&cashManager.Nickname,
		&cashManager.ShogunUsername,
		&cashManager.ChatId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return cashManager, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return cashManager, err
	}

	return cashManager, nil
}
