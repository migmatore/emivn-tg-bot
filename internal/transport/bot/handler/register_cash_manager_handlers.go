package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerCashManagerHandlers() {
	h.Message(h.CashManagerHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepCashManagerMainMenuHandler))

	// cash manager replenishment requests menu
	h.Message(h.CashManagerHandler.RepReqMenuHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqMenuHandler))

	h.Message(h.CashManagerHandler.RepReqSelectHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqSelectHandler))
	h.Message(h.CashManagerHandler.RepReqActionHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqActionHandler))
	h.Message(h.CashManagerHandler.RepReqConfirmActionHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqConfirmActionHandler))
}
