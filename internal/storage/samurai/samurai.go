package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
)

type SamuraiStorage struct {
	pool psql.AtomicPoolClient
}

func NewSamuraiStorage(pool psql.AtomicPoolClient) *SamuraiStorage {
	return &SamuraiStorage{pool: pool}
}

func (s *SamuraiStorage) Insert(ctx context.Context, samurai domain.Samurai) error {
	q := `insert into samurai(username, nickname, daimyo_username) values ($1, $2, $3)`

	if _, err := s.pool.Exec(ctx, q, samurai.Username, samurai.Nickname, samurai.DaimyoUsername); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *SamuraiStorage) GetByUsername(ctx context.Context, username string) (domain.Samurai, error) {
	q := `select username, nickname, daimyo_username, turnover_per_shift, chat_id from samurai where username=$1`

	samurai := domain.Samurai{}

	if err := s.pool.QueryRow(ctx, q, username).Scan(
		&samurai.Username,
		&samurai.Nickname,
		&samurai.DaimyoUsername,
		&samurai.TurnoverPerShift,
		&samurai.ChatId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return samurai, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return samurai, err
	}

	return samurai, nil
}

func (s *SamuraiStorage) GetByNickname(ctx context.Context, nickname string) (domain.Samurai, error) {
	q := `select username, nickname, daimyo_username, turnover_per_shift, chat_id from samurai where nickname=$1`

	samurai := domain.Samurai{}

	if err := s.pool.QueryRow(ctx, q, nickname).Scan(
		&samurai.Username,
		&samurai.Nickname,
		&samurai.DaimyoUsername,
		&samurai.TurnoverPerShift,
		&samurai.ChatId,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return samurai, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return samurai, err
	}

	return samurai, nil
}

func (s *SamuraiStorage) SetChatId(ctx context.Context, username string, id int64) error {
	q := `update samurai set chat_id=$1 where username=$2`

	if _, err := s.pool.Exec(ctx, q, id, username); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}

func (s *SamuraiStorage) GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.Samurai, error) {
	q := `select username, nickname, daimyo_username, turnover_per_shift, chat_id from samurai where daimyo_username=$1`

	samurais := make([]*domain.Samurai, 0)

	rows, err := s.pool.Query(ctx, q, daimyoUsername)
	if err != nil {
		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var samurai domain.Samurai

		err := rows.Scan(
			&samurai.Username,
			&samurai.Nickname,
			&samurai.DaimyoUsername,
			&samurai.TurnoverPerShift,
			&samurai.ChatId,
		)
		if err != nil {
			logging.GetLogger(ctx).Errorf("Query error. %v", err)
			return nil, err
		}

		samurais = append(samurais, &samurai)
	}

	return samurais, nil
}

func (s *SamuraiStorage) UpdateUsername(ctx context.Context, old string, new string) error {
	q := `update samurai set username=$1 where username=$2`

	if _, err := s.pool.Exec(ctx, q, new, old); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
