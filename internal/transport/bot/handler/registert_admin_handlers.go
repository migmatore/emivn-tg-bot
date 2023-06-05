package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandlers() {
	h.Message(h.AdminHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepAdminMainMenuHandler))
	h.Message(h.AdminHandler.HierarchyMenuHandler, h.isSessionStep(domain.SessionStepHierarchyMenuHandler))
	h.Message(h.AdminHandler.CreateEntityMenuHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler))

	// shogun
	h.Message(h.AdminHandler.EnterShogunUsername, h.isSessionStep(domain.SessionStepCreateShogunUsername))
	h.Message(h.AdminHandler.EnterShogunNicknameAndCreate, h.isSessionStep(domain.SessionStepCreateShogun))

	// daimyo
	h.Message(h.AdminHandler.EnterDaimyoUsername, h.isSessionStep(domain.SessionStepCreateDaimyoUsername))
	h.Message(h.AdminHandler.EnterDaimyoNickname, h.isSessionStep(domain.SessionStepCreateDaimyoNickname))
	h.Message(h.AdminHandler.CreateDaimyo, h.isSessionStep(domain.SessionStepCreateDaimyo))

	// samurai
	h.Message(h.AdminHandler.EnterSamuraiUsername, h.isSessionStep(domain.SessionStepCreateSamuraiUsername))
	h.Message(h.AdminHandler.EnterSamuraiNickname, h.isSessionStep(domain.SessionStepCreateSamuraiNickname))
	h.Message(h.AdminHandler.CreateSamurai, h.isSessionStep(domain.SessionStepCreateSamurai))

	// cash manager
	h.Message(h.AdminHandler.EnterCashManagerUsername, h.isSessionStep(domain.SessionStepCreateCashManagerUsername))
	h.Message(h.AdminHandler.EnterCashManagerNickname, h.isSessionStep(domain.SessionStepCreateCashManagerNickname))
	h.Message(h.AdminHandler.CreateCashManager, h.isSessionStep(domain.SessionStepCreateCashManager))

	// card
	h.Message(h.AdminHandler.CardBank, h.isSessionStep(domain.SessionStepCreateCardBank))
	h.Message(h.AdminHandler.EnterCardName, h.isSessionStep(domain.SessionStepCreateCardName))
	h.Message(h.AdminHandler.EnterCardLastDigits, h.isSessionStep(domain.SessionStepCreateCardLastDigits))
	h.Message(h.AdminHandler.EnterCardDailyLimit, h.isSessionStep(domain.SessionStepCreateCardDailyLimit))
	h.Message(h.AdminHandler.EnterCardDaimyoUsernameAndCreate, h.isSessionStep(domain.SessionStepCreateCard))
}
