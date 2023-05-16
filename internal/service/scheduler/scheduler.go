package scheduler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
	"emivn-tg-bot/pkg/logging"
	"errors"
	"sync"
	"time"
)

type SchedulerStorage interface {
	Insert(ctx context.Context, task domain.Task) error
	UpdateTime(ctx context.Context, time time.Time, status domain.TaskStatus) (domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
}

const (
	DefaultSleepDuration = 30 * time.Second
	MinimalSleepDuration = 1 * time.Second
)

var ErrFuncNotFoundInTaskFuncsMap = errors.New("function not found in TaskFuncsMap")

type SchedulerService struct {
	transactor storage.Transactor
	storage    SchedulerStorage

	sync.RWMutex
	listeners     domain.TaskFuncsMap
	sleepDuration time.Duration
}

func New(t storage.Transactor, storage SchedulerStorage) *SchedulerService {

	return &SchedulerService{
		transactor:    t,
		storage:       storage,
		listeners:     make(map[string]domain.TaskFunc, 0),
		sleepDuration: DefaultSleepDuration,
	}
}

func (s *SchedulerService) Configure(listeners domain.TaskFuncsMap, sleepDuration time.Duration) {
	s.listeners = listeners

	sleep := sleepDuration
	if sleep == 0 {
		sleep = DefaultSleepDuration
	}

	s.sleepDuration = sleep
}

func (s *SchedulerService) Add(ctx context.Context, dto domain.TaskDTO) error {
	s.RLock()
	defer s.RUnlock()
	if _, ok := s.listeners[dto.Alias]; !ok {
		return ErrFuncNotFoundInTaskFuncsMap
	}

	task := domain.Task{
		Alias:       dto.Alias,
		Name:        dto.Name,
		Arguments:   dto.Arguments.String(),
		Status:      domain.TaskStatusWait,
		Schedule:    dto.IntervalMinutes,
		ScheduledAt: dto.RunAt,
	}

	return s.storage.Insert(ctx, task)
}

func (s *SchedulerService) Run(ctx context.Context) error {
	for {
		if err := s.transactor.WithinTransaction(ctx, func(txCtx context.Context) error {
			task, err := s.storage.UpdateTime(txCtx, time.Now(), domain.TaskStatusWait)
			if err != nil {
				return err
			}

			if task.TaskId == 0 {
				time.Sleep(s.sleepDuration)
			} else {
				if fn, ok := s.listeners[task.Alias]; ok {
					go s.exec(txCtx, &task, fn)
				} else {
					task.Status = domain.TaskStatusDeferred
					if err := s.storage.Update(txCtx, &task); err != nil {
						return err
					}
				}
			}

			return nil
		}); err != nil {
			return err
		}
	}
}

func (s *SchedulerService) exec(ctx context.Context, task *domain.Task, fn domain.TaskFunc) {
	funcArgs := task.ParseArgs()

	status, when := fn(funcArgs)
	switch status { // nolint:exhaustive TaskStatusInProgress = default
	case domain.TaskStatusDone, domain.TaskStatusWait, domain.TaskStatusDeferred:
		task.Status = status
	default:
		task.Status = domain.TaskStatusDeferred
	}

	switch v := when.(type) {
	case time.Duration:
		task.ScheduledAt = task.ScheduledAt.Add(v)
	case time.Time:
		task.ScheduledAt = v
	default:
		if task.Schedule > 0 {
			d := time.Minute * time.Duration(task.Schedule)
			task.ScheduledAt = task.ScheduledAt.Add(time.Now().Sub(task.ScheduledAt).Truncate(d) + d)
		} else {
			task.Status = domain.TaskStatusDeferred
		}
	}

	if err := s.storage.Update(ctx, task); err != nil {
		logging.GetLogger(ctx).Errorf("Scheduler error: %v", err)
	}
}
