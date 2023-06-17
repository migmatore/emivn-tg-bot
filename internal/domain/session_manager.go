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

	SessionStepCreateControllerUsername
	SessionStepCreateController

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
	SessionStepDaimyoEnterReplenishmentRequestCardName
	SessionStepDaimyoEnterReplenishmentRequestAmount
	SessionStepDaimyoMakeReplenishmentRequest

	SessionStepDaimyoReportMenuHandler
	SessionStepDaimyoReportPeriodMenuHandler

	SessionStepDaimyoHierarchyMenuHandler

	// samurai steps
	SessionStepSamuraiEnterDataMenuHandler
	SessionStepSamuraiChooseBankMenuHandler
	SessionStepSamuraiCreateTurnoverHandler

	SessionStepDaimyoCreateSamuraiUsername
	SessionStepDaimyoCreateSamuraiNickname

	// cash manager steps
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
