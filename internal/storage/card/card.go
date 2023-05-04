package card

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type CardStorage struct {
	pool psql.AtomicPoolClient
}

func NewCardStorage(pool psql.AtomicPoolClient) *CardStorage {
	return &CardStorage{pool: pool}
}

func (s *CardStorage) Insert(ctx context.Context, card domain.Card) error {
	q := `insert into cards(name, last_digits, daily_limit, daimyo_username) values ($1, $2, $3, $4)`

	if _, err := s.pool.Exec(ctx, q, card.Name, card.LastDigits, card.DailyLimit, card.DaimyoUsername); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *CardStorage) GetByUsername(ctx context.Context, daimyoUsername string) ([]*domain.Card, error) {
	q := `select id, name, last_digits, daily_limit, daimyo_username from cards where daimyo_username=$1`

	cards := make([]*domain.Card, 0)

	rows, err := s.pool.Query(ctx, q, daimyoUsername)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var card domain.Card

		err := rows.Scan(&card.CardId, &card.Name, &card.LastDigits, &card.DailyLimit, &card.DaimyoUsername)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, nil
}
