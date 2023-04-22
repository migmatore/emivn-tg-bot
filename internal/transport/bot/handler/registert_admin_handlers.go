package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandler() {
	h.Message(h.AdminHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepAdminMenuHandler)).
		Message(h.AdminHandler.CreateEntityMenuSelectionHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler)).
		Message(h.AdminHandler.EnterShogunUsername, h.isSessionStep(domain.SessionStepCreateShogunUsername)).
		Message(h.AdminHandler.EnterShogunNicknameAndCreate, h.isSessionStep(domain.SessionStepCreateShogun)).
		Message(h.AdminHandler.EnterDaimyoUsername, h.isSessionStep(domain.SessionStepCreateDaimyoUsername)).
		Message(h.AdminHandler.EnterDaimyoNickname, h.isSessionStep(domain.SessionStepCreateDaimyoNickname)).
		Message(h.AdminHandler.CreateDaimyo, h.isSessionStep(domain.SessionStepCreateDaimyo))
}
