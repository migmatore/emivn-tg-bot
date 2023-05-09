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
	SessionStepEnterReplenishmentRequestAmount
	SessionStepMakeReplenishmentRequest
)

type Session struct {
	Step SessionStep

	Shogun               ShogunDTO
	Daimyo               DaimyoDTO
	Samurai              SamuraiDTO
	CashManager          CashManagerDTO
	Card                 CardDTO
	ReplenishmentRequest ReplenishmentRequestDTO
}
