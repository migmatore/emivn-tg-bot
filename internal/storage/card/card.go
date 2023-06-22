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
	q := `insert into cards(name, owner_username, last_digits, daily_limit, balance, bank_type_id) 
				values ($1, $2, $3, $4, $5, $6)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		card.Name,
		card.OwnerUsername,
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

func (s *CardStorage) GetAllByUsername(ctx context.Context, bankId int, ownerUsername string) ([]*domain.Card, error) {
	q := `select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards 
                where owner_username=$1 and bank_type_id=$2`

	cards := make([]*domain.Card, 0)

	rows, err := s.pool.Query(ctx, q, ownerUsername, bankId)
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
			&card.OwnerUsername,
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
	q := `select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards
				where owner_username in (select username from daimyo where shogun_username = $1)
		  union
		  select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards
				where owner_username in (select username from main_operators where shogun_username = $1)`

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
			&card.OwnerUsername,
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
	q := `select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards where name=$1`

	card := domain.Card{}

	if err := s.pool.QueryRow(ctx, q, name).Scan(
		&card.CardId,
		&card.Name,
		&card.OwnerUsername,
		&card.LastDigits,
		&card.DailyLimit,
		&card.Balance,
		&card.BankTypeId,
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

func (s *CardStorage) GetByUsername(ctx context.Context, daimyoUsername string) (domain.Card, error) {
	q := `select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards 
                                                                                  where owner_username=$1`

	card := domain.Card{}

	if err := s.pool.QueryRow(ctx, q, daimyoUsername).Scan(
		&card.CardId,
		&card.Name,
		&card.OwnerUsername,
		&card.LastDigits,
		&card.DailyLimit,
		&card.Balance,
		&card.BankTypeId,
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

func (s *CardStorage) GetById(ctx context.Context, cardId int) (domain.Card, error) {
	q := `select id, name, owner_username, last_digits, daily_limit, balance, bank_type_id from cards where id = $1`

	card := domain.Card{}

	if err := s.pool.QueryRow(ctx, q, cardId).Scan(
		&card.CardId,
		&card.Name,
		&card.OwnerUsername,
		&card.LastDigits,
		&card.DailyLimit,
		&card.Balance,
		&card.BankTypeId,
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

func (s *CardStorage) UpdateLimit(ctx context.Context, name string, limit int) error {
	q := `update cards set daily_limit = $1 where name = $2`

	if _, err := s.pool.Exec(
		ctx,
		q,
		limit,
		name,
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

func (s *CardStorage) GetBankById(ctx context.Context, bankId int) (domain.Bank, error) {
	q := `select id, name from bank_types where id=$1`

	bank := domain.Bank{}

	if err := s.pool.QueryRow(ctx, q, bankId).Scan(&bank.BankId, &bank.Name); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return bank, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return bank, err
	}

	return bank, nil
}
