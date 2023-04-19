package domain

type SessionStep int8

const (
	SessionStepStart = iota
	SessionStepInit
	SessionStepAcionSelect
	SessionStepReadData
	SessionStepWriteData
)

type Session struct {
	Step SessionStep

	Name   string
	Age    int
	Gender string
}
