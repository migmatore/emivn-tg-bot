package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandlers() {
	h.Message(h.AdminHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepAdminMenuHandler)).
		Message(h.AdminHandler.CreateEntityMenuSelectionHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler)).
		// shogun
		Message(h.AdminHandler.EnterShogunUsername, h.isSessionStep(domain.SessionStepCreateShogunUsername)).
		Message(h.AdminHandler.EnterShogunNicknameAndCreate, h.isSessionStep(domain.SessionStepCreateShogun)).
		// daimyo
		Message(h.AdminHandler.EnterDaimyoUsername, h.isSessionStep(domain.SessionStepCreateDaimyoUsername)).
		Message(h.AdminHandler.EnterDaimyoNickname, h.isSessionStep(domain.SessionStepCreateDaimyoNickname)).
		Message(h.AdminHandler.CreateDaimyo, h.isSessionStep(domain.SessionStepCreateDaimyo)).
		// samurai
		Message(h.AdminHandler.EnterSamuraiUsername, h.isSessionStep(domain.SessionStepCreateSamuraiUsername)).
		Message(h.AdminHandler.EnterSamuraiNickname, h.isSessionStep(domain.SessionStepCreateSamuraiNickname)).
		Message(h.AdminHandler.CreateSamurai, h.isSessionStep(domain.SessionStepCreateSamurai)).
		// cash manager
		Message(h.AdminHandler.EnterCashManagerUsername, h.isSessionStep(domain.SessionStepCreateCashManagerUsername)).
		Message(h.AdminHandler.EnterCashManagerNickname, h.isSessionStep(domain.SessionStepCreateCashManagerNickname)).
		Message(h.AdminHandler.CreateCashManager, h.isSessionStep(domain.SessionStepCreateCashManager)).
		// card
		Message(h.AdminHandler.CardBank, h.isSessionStep(domain.SessionStepCreateCardBank)).
		Message(h.AdminHandler.EnterCardName, h.isSessionStep(domain.SessionStepCreateCardName)).
		Message(h.AdminHandler.EnterCardLastDigits, h.isSessionStep(domain.SessionStepCreateCardLastDigits)).
		Message(h.AdminHandler.EnterCardDailyLimit, h.isSessionStep(domain.SessionStepCreateCardDailyLimit)).
		Message(h.AdminHandler.EnterCardDaimyoUsernameAndCreate, h.isSessionStep(domain.SessionStepCreateCard))
}
