package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandlers() {
	h.Message(h.AdminHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepAdminMainMenuHandler))
	h.Message(h.AdminHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepHierarchyMenuHandler))
	h.Message(h.AdminHandler.CreateEntityMenuHandler, h.isSessionStep(domain.SessionStepCreateEntityMenuHandler))

	// shogun
	h.Message(h.AdminHandler.EnterShogunNicknameHandler, h.isSessionStep(domain.SessionStepCreateShogunNickname))
	h.Message(h.AdminHandler.ShogunCreationHandler, h.isSessionStep(domain.SessionStepShogunCreationMethod))
	h.Message(h.AdminHandler.EnterShogunUsernameAndCreateHandler, h.isSessionStep(domain.SessionStepCreateShogun))

	// daimyo
	h.Message(h.AdminHandler.EnterDaimyoNicknameHandler, h.isSessionStep(domain.SessionStepCreateDaimyoNickname))
	h.Message(h.AdminHandler.ChooseDaimyoShogunHandler, h.isSessionStep(domain.SessionStepChooseDaimyoShogun))
	h.Message(h.AdminHandler.DaimyoCreationHandler, h.isSessionStep(domain.SessionStepDaimyoCreationMethod))
	h.Message(h.AdminHandler.CreateDaimyoHandler, h.isSessionStep(domain.SessionStepCreateDaimyo))

	// samurai
	h.Message(h.AdminHandler.EnterSamuraiNicknameHandler, h.isSessionStep(domain.SessionStepCreateSamuraiNickname))
	h.Message(h.AdminHandler.ChooseSamuraiDaimyoHandler, h.isSessionStep(domain.SessionStepChooseSamuraiDaimyo))
	h.Message(h.AdminHandler.SamuraiCreationHandler, h.isSessionStep(domain.SessionStepSamuraiCreationMethod))
	h.Message(h.AdminHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepCreateSamurai))

	// cash manager
	h.Message(h.AdminHandler.EnterCashManagerNicknameHandler, h.isSessionStep(domain.SessionStepCreateCashManagerNickname))
	h.Message(h.AdminHandler.ChooseCashManagerShogunHandler, h.isSessionStep(domain.SessionStepChooseCashManagerShogun))
	h.Message(h.AdminHandler.CashManagerCreationHandler, h.isSessionStep(domain.SessionStepCashManagerCreationMethod))
	h.Message(h.AdminHandler.CreateCashManagerHandler, h.isSessionStep(domain.SessionStepCreateCashManager))

	// controller
	h.Message(h.AdminHandler.EnterControllerNicknameHandler, h.isSessionStep(domain.SessionStepCreateControllerNickname))
	h.Message(h.AdminHandler.ControllerCreationHandler, h.isSessionStep(domain.SessionStepControllerCreationMethod))
	h.Message(h.AdminHandler.EnterControllerUsernameAndCreateHandler, h.isSessionStep(domain.SessionStepCreateController))

	// main operator
	h.Message(h.AdminHandler.EnterMainOperatorNicknameHandler, h.isSessionStep(domain.SessionStepCreateMainOperatorNickname))
	h.Message(h.AdminHandler.ChooseMainOperatorShogunHandler, h.isSessionStep(domain.SessionStepChooseMainOperatorShogun))
	h.Message(h.AdminHandler.MainOperatorCreationHandler, h.isSessionStep(domain.SessionStepMainOperatorCreationMethod))
	h.Message(h.AdminHandler.CreateMainOperatorHandler, h.isSessionStep(domain.SessionStepCreateMainOperator))

	// cards menu
	h.Message(h.AdminHandler.CardsChooseShogunHandler, h.isSessionStep(domain.SessionStepAdminCardsChooseShogunHandler))
	h.Message(h.AdminHandler.CardsMenuHandler, h.isSessionStep(domain.SessionStepAdminCardsMenuHandler))

	// create card
	h.Message(h.AdminHandler.ChooseCardBankMenuHandler, h.isSessionStep(domain.SessionStepAdminChooseCardBankHandler))
	h.Message(h.AdminHandler.EnterCardNameHandler, h.isSessionStep(domain.SessionStepAdminEnterCardNameHandler))
	h.Message(h.AdminHandler.EnterCardLastDigitsHandler, h.isSessionStep(domain.SessionStepAdminEnterCardLastDigitsHandler))
	h.Message(h.AdminHandler.SetCardLimitHandler, h.isSessionStep(domain.SessionStepAdminSetCardLimitHandler))
	h.Message(h.AdminHandler.ChooseCardDaimyoAndCreateHandler, h.isSessionStep(domain.SessionStepAdminChooseCardDaimyoHandler))
}
