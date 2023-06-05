package scheduler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"time"
)

type SchedulerStorage interface {
	Insert(ctx context.Context, task domain.Task) error
	UpdateTime(ctx context.Context, time time.Time, status domain.TaskStatus) (domain.Task, error)
	Update(ctx context.Context, task *domain.Task) error
}

type SchedulerService struct {
	storage SchedulerStorage
}

func New(storage SchedulerStorage) *SchedulerService {
	return &SchedulerService{
		storage: storage,
	}
}

func (s *SchedulerService) Create(ctx context.Context, dto domain.TaskDTO) error {
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

func (s *SchedulerService) UpdateTime(ctx context.Context, time time.Time, status domain.TaskStatus) (domain.Task, error) {
	return s.storage.UpdateTime(ctx, time, status)
}

func (s *SchedulerService) Update(ctx context.Context, task *domain.Task) error {
	return s.storage.Update(ctx, task)
}
