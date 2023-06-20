package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerCashManagerHandlers() {
	h.Message(h.CashManagerHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepCashManagerMainMenuHandler))

	// cash manager replenishment requests menu
	h.Message(h.CashManagerHandler.RepReqMenuHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqMenuHandler))
}
