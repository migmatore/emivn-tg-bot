package scheduler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/pkg/logging"
	"github.com/jackc/pgx/v4"
	"github.com/migmatore/bakery-shop-api/pkg/utils"
	"time"
)

type SchedulerStorage struct {
	pool psql.AtomicPoolClient
}

func NewSchedulerStorage(pool psql.AtomicPoolClient) *SchedulerStorage {
	return &SchedulerStorage{pool: pool}
}

func (s *SchedulerStorage) Insert(ctx context.Context, task domain.Task) error {
	q := `insert into 
    		tasks(alias, name, arguments, status, schedule, scheduled_at, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, $6, $7, $8)`

	if _, err := s.pool.Exec(
		ctx,
		q,
		task.Alias,
		task.Name,
		task.Arguments,
		task.Status,
		task.Schedule,
		task.ScheduledAt,
		task.CreatedAt,
		task.UpdatedAt,
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

func (s *SchedulerStorage) UpdateTime(ctx context.Context, time1 time.Time, status domain.TaskStatus) (domain.Task, error) {
	q := `update tasks set updated_at=$1 
             where id = (
             	select id from tasks 
             	where scheduled_at < now() and status=$2
             	order by scheduled_at limit 1 for update skip locked) 
             returning id, alias, name, arguments, status, schedule, scheduled_at, created_at, updated_at`

	task := domain.Task{}

	if err := s.pool.QueryRow(ctx, q, time1, status).Scan(
		&task.TaskId,
		&task.Alias,
		&task.Name,
		&task.Arguments,
		&task.Status,
		&task.Schedule,
		&task.ScheduledAt,
		&task.CreatedAt,
		&task.UpdatedAt,
	); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			if err == pgx.ErrNoRows {
				return task, err
			}

			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return task, err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return task, err
	}

	return task, nil
}

func (s *SchedulerStorage) Update(ctx context.Context, task domain.Task) error {
	q := `update tasks set alias=$1, name=$2, arguments=$3, status=$4, schedule=$5, scheduled_at=$6, 
                 created_at=$7, updated_at=$8 
             where id=$9`

	if _, err := s.pool.Exec(
		ctx,
		q,
		task.Alias,
		task.Name,
		task.Arguments,
		task.Status,
		task.Schedule,
		task.ScheduledAt,
		task.CreatedAt,
		task.UpdatedAt,
		task.TaskId,
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

func (s *SchedulerStorage) Delete(ctx context.Context, taskName string) error {
	q := `delete from tasks where name = $1`

	if _, err := s.pool.Exec(ctx, q, taskName); err != nil {
		if err := utils.ParsePgError(err); err != nil {
			logging.GetLogger(ctx).Errorf("Error: %v", err)
			return err
		}

		logging.GetLogger(ctx).Errorf("Query error. %v", err)
		return err
	}

	return nil
}
