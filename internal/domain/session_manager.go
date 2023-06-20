package domain

type SessionStep int8

const (
	SessionStepInit = iota

	// admin steps
	SessionStepAdminMainMenuHandler

	// admin cards menu
	SessionStepAdminCardsChooseShogunHandler
	SessionStepAdminCardsMenuHandler

	// admin create card steps
	SessionStepAdminChooseCardBankHandler
	SessionStepAdminEnterCardNameHandler
	SessionStepAdminEnterCardLastDigitsHandler
	SessionStepAdminSetCardLimitHandler
	SessionStepAdminChooseCardDaimyoHandler

	// admin hierrarchy menu
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

	// shogun steps
	SessionStepShogunMainMenuHandler

	// shogun cards menu
	SessionStepShogunCardsMenuHandler

	// shogun create card steps
	SessionStepShogunChooseCardBankHandler
	SessionStepShogunEnterCardNameHandler
	SessionStepShogunEnterCardLastDigitsHandler
	SessionStepShogunSetCardLimitHandler
	SessionStepShogunChooseCardDaimyoHandler

	// shogun hierarchy menu
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

	SessionStepShogunSubordinationMenuHandler

	// diamyo steps
	SessionStepDaimyoMainMenuHandler

	// daimyo replenishment request menu
	SessionStepDaimyoChooseReplenishmentRequestBank

	// daimyo make replenishment request
	SessionStepDaimyoEnterReplenishmentRequestAmount
	SessionStepDaimyoMakeReplenishmentRequest
	SessionStepDaimyoChangeReplenishmentRequestAmount

	// daimyo report menu
	SessionStepDaimyoReportMenuHandler

	// daimyo make report
	SessionStepDaimyoReportPeriodMenuHandler

	// daimyo hierarchy menu
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
