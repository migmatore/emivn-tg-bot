package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerAdminHandler() {
	h.Message(h.AdminHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepAdminMenuHandler)).
		Message(h.AdminHandler.CreateEntityMenuSelectionHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler)).
		Message(h.AdminHandler.CreateShogun, h.isSessionStep(domain.SessionStepCreateShogun))
}
