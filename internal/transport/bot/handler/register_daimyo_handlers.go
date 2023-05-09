package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerDaimyoHandler() {
	h.Message(h.DaimyoHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepDaimyoMenuHandler)).
		Message(h.DaimyoHandler.MakeReplenishmentRequest, h.isSessionStep(domain.SessionStepMakeReplenishmentRequest)).
		Message(h.DaimyoHandler.EnterReplenishmentRequestAmount, h.isSessionStep(domain.SessionStepEnterReplenishmentRequestAmount)).
		Message(h.DaimyoHandler.MakeReplenishmentRequest, h.isSessionStep(domain.SessionStepMakeReplenishmentRequest))
}
