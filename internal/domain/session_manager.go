package domain

type SessionStep int8

const (
	SessionStepInit = iota
	SessionStepStart
	SessionStepAdminRole
	SessionStepAdminMenuHandler
	SessionStepCreateEntityButton
	SessionStepCreateEntityHandler
	SessionStepBackAdminMenuButton
	SessionStepCreateShogun
	SessionStepCreateShogunUsername
	SessionStepBackCreateEntityMenuStep
	//SessionStepAcionSelect
	//SessionStepReadData
	//SessionStepWriteData
)

type Session struct {
	Step SessionStep

	Name   string
	Age    int
	Gender string
}
