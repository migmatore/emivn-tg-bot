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

	SessionStepCreateShogunUsername
	SessionStepCreateShogun

	SessionStepCreateDaimyoUsername
	SessionStepCreateDaimyoNickname
	SessionStepCreateDaimyo

	SessionStepCreateSamuraiUsername
	SessionStepCreateSamurai

	SessionStepCreateCashManagerUsername
	SessionStepCreateCashManager

	SessionStepCreateCardBankInfo
	SessionStepCreateCard

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
