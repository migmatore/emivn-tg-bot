package domain

type SessionStep int8

const (
	SessionStepInit = iota

	// admin steps
	SessionStepAdminMenuHandler

	SessionStepCreateEntityHandler

	SessionStepCreateShogunUsername
	SessionStepCreateShogun

	SessionStepCreateDaimyoUsername
	SessionStepCreateDaimyoNickname
	SessionStepCreateDaimyo

	SessionStepCreateSamuraiUsername
	SessionStepCreateSamuraiNickname
	SessionStepCreateSamurai

	SessionStepCreateCashManagerUsername
	SessionStepCreateCashManagerNickname
	SessionStepCreateCashManager

	SessionStepCreateCardName
	SessionStepCreateCardLastDigits
	SessionStepCreateCardDailyLimit
	SessionStepCreateCard

	SessionStepBackCreateEntityMenuStep

	// diamyo steps
	SessionStepDaimyoMenuHandler
	SessionStepMakeReplenishmentRequest
)

type Session struct {
	Step SessionStep

	Name   string
	Age    int
	Gender string
}
