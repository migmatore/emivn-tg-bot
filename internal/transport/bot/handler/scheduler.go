package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/logging"
	"sync"
	"time"
)

const (
	DefaultSleepDuration = 30 * time.Second
	MinimalSleepDuration = 1 * time.Second
)

type SchedulerService interface {
	Add(ctx context.Context, dto domain.TaskDTO) error

	UpdateTime(ctx context.Context, time time.Time, status domain.TaskStatus) (domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
}

type Scheduler struct {
	transactorService TransactorService
	schedulerService  SchedulerService

	sync.RWMutex
	listeners     domain.TaskFuncsMap
	sleepDuration time.Duration
}

func NewScheduler(transactorService TransactorService, schedulerService SchedulerService) *Scheduler {
	return &Scheduler{
		transactorService: transactorService,
		schedulerService:  schedulerService,
	}
}

func (s *Scheduler) Configure(listeners domain.TaskFuncsMap, sleepDuration time.Duration) {
	s.listeners = listeners

	sleep := sleepDuration
	if sleep == 0 {
		sleep = DefaultSleepDuration
	}

	s.sleepDuration = sleep
}

func (s *Scheduler) Run(ctx context.Context) error {
	for {
		if err := s.transactorService.WithinTransaction(ctx, func(txCtx context.Context) error {
			task, err := s.schedulerService.UpdateTime(txCtx, time.Now(), domain.TaskStatusWait)
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
					if err := s.schedulerService.Update(txCtx, &task); err != nil {
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

func (s *Scheduler) exec(ctx context.Context, task *domain.Task, fn domain.TaskFunc) {
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

	if err := s.schedulerService.Update(ctx, task); err != nil {
		logging.GetLogger(ctx).Errorf("Scheduler error: %v", err)
	}
}
