package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerShogunHandler() {
	h.Message(h.ShogunHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepShogunMainMenuHandler))

	// hierarchy menu
	h.Message(h.ShogunHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepShogunHierarchyMenuHandler))

	// create entity menu
	h.Message(h.ShogunHandler.CreateEntityMenuHandler, h.isSessionStep(domain.SessionStepShogunCreateEntityMenuHandler))

	h.Message(h.ShogunHandler.EnterDaimyoNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateDaimyoNickname))
	h.Message(h.ShogunHandler.DaimyoCreationHandler, h.isSessionStep(domain.SessionStepShogunDaimyoCreationMethod))
	h.Message(h.ShogunHandler.CreateDaimyoHandler, h.isSessionStep(domain.SessionStepShogunCreateDaimyo))

	h.Message(h.ShogunHandler.EnterSamuraiNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateSamuraiNickname))
	h.Message(h.ShogunHandler.ChooseSamuraiDaimyoHandler, h.isSessionStep(domain.SessionStepShogunChooseSamuraiDaimyo))
	h.Message(h.ShogunHandler.SamuraiCreationHandler, h.isSessionStep(domain.SessionStepShogunSamuraiCreationMethod))
	h.Message(h.ShogunHandler.CreateSamuraiHandler, h.isSessionStep(domain.SessionStepShogunCreateSamurai))

	h.Message(h.ShogunHandler.EnterCashManagerNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateCashManagerNickname))
	h.Message(h.ShogunHandler.CashManagerCreationHandler, h.isSessionStep(domain.SessionStepShogunCashManagerCreationMethod))
	h.Message(h.ShogunHandler.CreateCashManagerHandler, h.isSessionStep(domain.SessionStepShogunCreateCashManager))

	h.Message(h.ShogunHandler.EnterMainOperatorNicknameHandler, h.isSessionStep(domain.SessionStepShogunCreateMainOperatorNickname))
	h.Message(h.ShogunHandler.MainOperatorCreationHandler, h.isSessionStep(domain.SessionStepShogunMainOperatorCreationMethod))
	h.Message(h.ShogunHandler.CreateMainOperatorHandler, h.isSessionStep(domain.SessionStepShogunCreateMainOperator))

	// cards menu

	h.Message(h.ShogunHandler.CardsMenuHandler, h.isSessionStep(domain.SessionStepShogunCardsMenuHandler))

	// create card
	h.Message(h.ShogunHandler.ChooseCardBankMenuHandler, h.isSessionStep(domain.SessionStepShogunChooseCardBankHandler))
	h.Message(h.ShogunHandler.EnterCardNameHandler, h.isSessionStep(domain.SessionStepShogunEnterCardNameHandler))
	h.Message(h.ShogunHandler.EnterCardLastDigitsHandler, h.isSessionStep(domain.SessionStepShogunEnterCardLastDigitsHandler))
	h.Message(h.ShogunHandler.SetCardLimitHandler, h.isSessionStep(domain.SessionStepShogunSetCardLimitHandler))
	h.Message(h.ShogunHandler.ChooseCardDaimyoAndCreateHandler, h.isSessionStep(domain.SessionStepShogunChooseCardDaimyoHandler))
}
