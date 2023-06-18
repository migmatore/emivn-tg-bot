package domain

type SessionStep int8

const (
	SessionStepInit = iota

	// admin steps
	SessionStepAdminMainMenuHandler
	SessionStepHierarchyMenuHandler

	SessionStepCreateEntityMenuHandler

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

	SessionStepShogunHierarchyMenuHandler

	SessionStepShogunCreateEntityMenuHandler

	SessionStepShogunCreateDaimyoNickname
	SessionStepShogunCreateDaimyo
	SessionStepShogunCreateSamurai

	SessionStepShogunCreateSamuraiNickname
	SessionStepShogunChooseSamuraiDaimyo

	SessionStepShogunCreateCashManagerNickname
	SessionStepShogunCreateCashManager

	SessionStepShogunCreateMainOperatorNickname
	SessionStepShogunCreateMainOperator

	// diamyo steps
	SessionStepDaimyoMainMenuHandler
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
	MainOperator         MainOperatorDTO
	Card                 CardDTO
	ReplenishmentRequest ReplenishmentRequestDTO
	SamuraiTurnover      SamuraiTurnoverDTO
	ControllerTurnover   ControllerTurnoverDTO
}
