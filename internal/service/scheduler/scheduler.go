package scheduler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"errors"
	"sync"
	"time"
)

type SchedulerStorage interface {
	Insert(ctx context.Context, task domain.Task) error
}

const (
	DefaultSleepDuration = 30 * time.Second
	MinimalSleepDuration = 1 * time.Second
)

type (
	TaskPlan map[string]uint

	TaskFunc func(args domain.FuncArgs) (status domain.TaskStatus, when interface{})

	// TaskFuncsMap - list by TaskFunc's (key - task alias, value - TaskFunc)
	TaskFuncsMap map[string]TaskFunc
)

var ErrFuncNotFoundInTaskFuncsMap = errors.New("function not found in TaskFuncsMap")

type SchedulerService struct {
	storage SchedulerStorage

	sync.RWMutex
	listeners     TaskFuncsMap
	sleepDuration time.Duration
}

func New(storage SchedulerStorage, listeners TaskFuncsMap, sleepDuration time.Duration) *SchedulerService {
	sleep := sleepDuration
	if sleep == 0 {
		sleep = DefaultSleepDuration
	}

	return &SchedulerService{storage: storage, listeners: listeners, sleepDuration: sleep}
}

func (c *SchedulerService) Add(ctx context.Context, dto domain.TaskDTO) error {
	c.RLock()
	defer c.RUnlock()
	if _, ok := c.listeners[dto.Alias]; !ok {
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

	return c.storage.Insert(ctx, task)
}
