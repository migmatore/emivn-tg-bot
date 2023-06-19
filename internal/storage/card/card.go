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
	q := `insert into cards(name, daimyo_username, last_digits, daily_limit, balance, bank_type_id) 
				values ($1, $2, $3, $4, $5, $6)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		card.Name,
		card.DaimyoUsername,
		card.LastDigits,
		card.DailyLimit,
		card.Balance,
		card.BankTypeId,
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

func (s *CardStorage) GetAllByUsername(ctx context.Context, bankId int, daimyoUsername string) ([]*domain.Card, error) {
	q := `select id, name, daimyo_username, last_digits, daily_limit, balance, bank_type_id from cards 
                where daimyo_username=$1 and bank_type_id=$2`

	cards := make([]*domain.Card, 0)

	rows, err := s.pool.Query(ctx, q, daimyoUsername, bankId)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var card domain.Card

		err := rows.Scan(
			&card.CardId,
			&card.Name,
			&card.DaimyoUsername,
			&card.LastDigits,
			&card.DailyLimit,
			&card.Balance,
			&card.BankTypeId,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, nil
}

func (s *CardStorage) GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.Card, error) {
	q := `select id, name, daimyo_username, last_digits, daily_limit, balance, bank_type_id from cards
			where daimyo_username in (select username from daimyo where shogun_username = $1)`

	cards := make([]*domain.Card, 0)

	rows, err := s.pool.Query(ctx, q, shogunUsername)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var card domain.Card

		err := rows.Scan(
			&card.CardId,
			&card.Name,
			&card.DaimyoUsername,
			&card.LastDigits,
			&card.DailyLimit,
			&card.Balance,
			&card.BankTypeId,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		cards = append(cards, &card)
	}

	return cards, nil
}

func (s *CardStorage) GetByName(ctx context.Context, name string) (domain.Card, error) {
	q := `select id, name, last_digits, daily_limit, daimyo_username from cards where name=$1`

	card := domain.Card{}

	if err := s.pool.QueryRow(ctx, q, name).Scan(
		&card.CardId,
		&card.Name,
		&card.LastDigits,
		&card.DailyLimit,
		&card.DaimyoUsername,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return card, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return card, err
	}

	return card, nil
}

func (s *CardStorage) GetBankNames(ctx context.Context) ([]*domain.Bank, error) {
	q := `select id, name from bank_types`

	banks := make([]*domain.Bank, 0)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var bank domain.Bank

		err := rows.Scan(&bank.BankId, &bank.Name)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		banks = append(banks, &bank)
	}

	return banks, nil
}

func (s *CardStorage) GetBankIdByName(ctx context.Context, bankName string) (int, error) {
	q := `select id from bank_types where name=$1`

	var id int

	if err := s.pool.QueryRow(ctx, q, bankName).Scan(&id); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return id, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return id, err
	}

	return id, nil
}
