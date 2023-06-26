package domain

type SessionStep int8

const (
	SessionStepInit = iota

	// admin steps
	// admin main menu
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

	// admin create entity menu
	SessionStepCreateEntityMenuHandler

	// admin create shogun
	SessionStepCreateShogunNickname
	SessionStepCreateShogun

	// admin create daimyo
	SessionStepCreateDaimyoUsername
	SessionStepCreateDaimyoNickname
	SessionStepCreateDaimyo

	// admin create samurai
	SessionStepCreateSamuraiUsername
	SessionStepCreateSamuraiNickname
	SessionStepCreateSamurai

	// admin create cash manager
	SessionStepCreateCashManagerUsername
	SessionStepCreateCashManagerNickname
	SessionStepCreateCashManager

	// admin create cntroller
	SessionStepCreateControllerNickname
	SessionStepCreateController

	// admin create main operator
	SessionStepCreateMainOperatorUsername
	SessionStepCreateMainOperatorNickname
	SessionStepCreateMainOperator

	// shogun steps
	// shogun main menu
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

	// shogun create entity menu
	SessionStepShogunCreateEntityMenuHandler

	// shogun create daimyo
	SessionStepShogunCreateDaimyoNickname
	SessionStepShogunCreateDaimyo

	// shogun create samurai
	SessionStepShogunCreateSamuraiNickname
	SessionStepShogunChooseSamuraiDaimyo
	SessionStepShogunCreateSamurai

	// shogun create cash manager
	SessionStepShogunCreateCashManagerNickname
	SessionStepShogunCreateCashManager

	// shogun create main operator
	SessionStepShogunCreateMainOperatorNickname
	SessionStepShogunCreateMainOperator

	// shogun in subordination menu
	SessionStepShogunSubordinationMenuHandler

	// diamyo steps
	// daimyo main menu
	SessionStepDaimyoMainMenuHandler

	// daimyo replenishment request menu
	SessionStepDaimyoChooseReplenishmentRequestBank

	// daimyo make replenishment request
	SessionStepDaimyoEnterReplenishmentRequestAmount
	SessionStepDaimyoMakeReplenishmentRequest
	SessionStepDaimyoChangeReplenishmentRequestAmount

	// daimyo requests menu
	SessionStepDaimyoRepReqMenuHandler

	// daimyo objectionable replenishment request selection handler
	SessionStepDaimyoObjRepReqSelectHandler

	// daimyo objectionable replenishment request action handler
	SessionStepDaimyoObjRepReqActionHandler

	SessionStepDaimyoRepReqAnotherAmountHandler

	// daimyo report menu
	SessionStepDaimyoReportMenuHandler

	// daimyo make report
	SessionStepDaimyoReportPeriodMenuHandler

	// daimyo hierarchy menu
	SessionStepDaimyoHierarchyMenuHandler

	// daimyo create samurai
	SessionStepDaimyoCreateSamuraiUsername
	SessionStepDaimyoCreateSamuraiNickname

	// samurai steps
	SessionStepSamuraiEnterDataMenuHandler
	SessionStepSamuraiChooseBankMenuHandler
	SessionStepSamuraiCreateTurnoverHandler

	// cash manager steps
	// cash mamager main menu
	SessionStepCashManagerMainMenuHandler

	// cash manager replenishment requests handler
	SessionStepCashManagerRepReqMenuHandler

	// cash manager active replenishment request selection handler
	SessionStepCashManagerActRepReqSelectHandler

	SessionStepCashManagerActRepReqActionHandler

	SessionStepCashManagerActRepReqConfirmActionHandler

	SessionStepCashManagerRepReqAnotherAmountHandler

	SessionStepCashManagerObjRepReqSelectHandler
	SessionStepCashManagerObjRepReqAnotherAmountSelectHandler

	// controller steps
	SessionStepControllerEnterDataMenuHandler
	SessionStepControllerChooseDaimyoMenuHandler
	SessionStepControllerChooseSamuraiMenuHandler
	SessionStepControllerChooseBankMenuHandler
	SessionStepControllerCreateTurnoverHandler

	// main operator steps
	// main operator main menu
	SessionStepMainOperatorMainMenuHandler

	// main operator replenishment request menu
	SessionStepMainOperatorChooseReplenishmentRequestBank

	// main operator make replenishment request
	SessionStepMainOperatorEnterReplenishmentRequestAmount
	SessionStepMainOperatorChangeReplenishmentRequestAmount
	SessionStepMainOperatorMakeReplenishmentRequest
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
