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

	// shogun steps
	SessionStepShogunMainMenuHandler

	// diamyo steps
	SessionStepDaimyoMenuHandler
	SessionStepEnterReplenishmentRequestCardName
	SessionStepEnterReplenishmentRequestAmount
	SessionStepMakeReplenishmentRequest

	// samurai steps
	SessionStepSamuraiEnterDataMenuHandler
	SessionStepSamuraiChooseBankMenuHandler
	SessionStepSamuraiCreateTurnoverHandler

	SessionStepCashManagerMenuHandler

	// controller steps
	SessionStepControllerEnterDataMenuHandler
	SessionStepControllerChooseDaimyoMenuHandler
	SessionStepControllerChooseSamuraiMenuHandler
	SessionStepControllerChooseBankMenuHandler
	SessionStepControllerCreateTurnoverHandler
)

type Session struct {
	Step SessionStep

	Shogun               ShogunDTO
	Daimyo               DaimyoDTO
	Samurai              SamuraiDTO
	CashManager          CashManagerDTO
	Controller           ControllerDTO
	Card                 CardDTO
	ReplenishmentRequest ReplenishmentRequestDTO
	SamuraiTurnover      SamuraiTurnoverDTO
	ControllerTurnover   ControllerTurnoverDTO
}
