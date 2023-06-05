package domain

type SessionStep int8

const (
	SessionStepInit = iota

	// admin steps
	SessionStepAdminMainMenuHandler
	SessionStepHierarchyMenuHandler

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

	SessionStepCreateCardBank
	SessionStepCreateCardName
	SessionStepCreateCardLastDigits
	SessionStepCreateCardDailyLimit
	SessionStepCreateCard

	SessionStepBackCreateEntityMenuStep

	// diamyo steps
	SessionStepDaimyoMenuHandler
	SessionStepEnterReplenishmentRequestCardName
	SessionStepEnterReplenishmentRequestAmount
	SessionStepMakeReplenishmentRequest

	SessionStepSamuraiMenuHandler

	SessionStepCashManagerMenuHandler
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
