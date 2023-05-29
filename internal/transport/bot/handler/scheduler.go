package handler

//type SchedulerService interface {
//	Configure(listeners domain.TaskFuncsMap, sleepDuration time.Duration)
//	Add(ctx context.Context, dto domain.TaskDTO) error
//	Run(ctx context.Context) error
//}

type Scheduler struct {
	schedulerService SchedulerService
}

func NewScheduler(schedulerService SchedulerService) *Scheduler {
	return &Scheduler{
		schedulerService: schedulerService,
	}
}
