package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandlers() {
	h.Message(h.AdminHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepAdminMainMenuHandler))
	h.Message(h.AdminHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepHierarchyMenuHandler))
	h.Message(h.AdminHandler.CreateEntityMenuHandler, h.isSessionStep(domain.SessionStepCreateEntityMenuHandler))

	// shogun
	h.Message(h.AdminHandler.EnterShogunUsername, h.isSessionStep(domain.SessionStepCreateShogunUsername))
	h.Message(h.AdminHandler.EnterShogunNicknameAndCreate, h.isSessionStep(domain.SessionStepCreateShogun))

	// daimyo
	h.Message(h.AdminHandler.EnterDaimyoUsername, h.isSessionStep(domain.SessionStepCreateDaimyoUsername))
	h.Message(h.AdminHandler.EnterDaimyoNickname, h.isSessionStep(domain.SessionStepCreateDaimyoNickname))
	h.Message(h.AdminHandler.CreateDaimyo, h.isSessionStep(domain.SessionStepCreateDaimyo))

	// samurai
	h.Message(h.AdminHandler.EnterSamuraiUsernameHandler, h.isSessionStep(domain.SessionStepCreateSamuraiUsername))
	h.Message(h.AdminHandler.EnterSamuraiNicknameHandler, h.isSessionStep(domain.SessionStepCreateSamuraiNickname))
	h.Message(h.AdminHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepCreateSamurai))

	// cash manager
	h.Message(h.AdminHandler.EnterCashManagerUsername, h.isSessionStep(domain.SessionStepCreateCashManagerUsername))
	h.Message(h.AdminHandler.EnterCashManagerNickname, h.isSessionStep(domain.SessionStepCreateCashManagerNickname))
	h.Message(h.AdminHandler.CreateCashManager, h.isSessionStep(domain.SessionStepCreateCashManager))

	// controller
	h.Message(h.AdminHandler.EnterControllerUsername, h.isSessionStep(domain.SessionStepCreateControllerUsername))
	h.Message(h.AdminHandler.EnterControllerNicknameAndCreate, h.isSessionStep(domain.SessionStepCreateController))

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
