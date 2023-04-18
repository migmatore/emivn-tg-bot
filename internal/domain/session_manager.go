package domain

type SessionStep int8

const (
	SessionStepInit = iota
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
