package handler

import (
	"emivn-tg-bot/internal/domain"
	"sync"
	"time"
)

//type SchedulerService interface {
//	Configure(listeners domain.TaskFuncsMap, sleepDuration time.Duration)
//	Add(ctx context.Context, dto domain.TaskDTO) error
//	Run(ctx context.Context) error
//}

const (
	DefaultSleepDuration = 30 * time.Second
	MinimalSleepDuration = 1 * time.Second
)

type Scheduler struct {
	schedulerService SchedulerService

	sync.RWMutex
	listeners     domain.TaskFuncsMap
	sleepDuration time.Duration
}

func NewScheduler(schedulerService SchedulerService) *Scheduler {
	return &Scheduler{
		schedulerService: schedulerService,
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
