package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerDaimyoHandler() {
	h.Message(h.DaimyoHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepDaimyoMenuHandler))
}
