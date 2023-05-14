package domain

import (
	"encoding/json"
	"log"
	"time"
)

type TaskStatus uint

const (
	TaskStatusWait TaskStatus = iota
	TaskStatusDeferred
	TaskStatusInProgress
	TaskStatusDone
)

type FuncArgs map[string]interface{}

type Task struct {
	TaskId      int
	Alias       string
	Name        string
	Arguments   string
	Status      TaskStatus
	Schedule    uint
	ScheduledAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Task) ParseArgs() FuncArgs {
	if t.Arguments == "" {
		return nil
	}

	args := make(FuncArgs)

	err := json.Unmarshal([]byte(t.Arguments), &args)
	if err != nil {
		log.Print("ParseArgs() err:", err)
	}

	return args
}

type TaskDTO struct {
	Alias           string
	Name            string
	Arguments       FuncArgs
	Status          TaskStatus
	IntervalMinutes uint
	RunAt           time.Time
}

func (args *FuncArgs) String() string {
	str, err := json.Marshal(args)
	if err != nil {
		log.Print("FuncArgs.String() err:", err)
	}

	return string(str)
}
