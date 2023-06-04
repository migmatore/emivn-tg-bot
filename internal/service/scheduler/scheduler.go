package scheduler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/storage"
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
