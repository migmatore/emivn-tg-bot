package handler

import (
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
)

func (h *Handler) registerAdminHandler() {
	h.Message(h.AdminHandler.Menu, tgb.Command("/menu")).
		Message(h.AdminHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepAdminMenuHandler)).
		Message(h.AdminHandler.CreateEntityMenu, h.isSessionStep(domain.SessionStepCreateEntityButton)).
		Message(h.AdminHandler.CreateEntityMenuSelectionHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler)).
		Message(h.AdminHandler.CreateShogun, h.isSessionStep(domain.SessionStepCreateShogun))
}
