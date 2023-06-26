package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerCashManagerHandlers() {
	h.Message(h.CashManagerHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepCashManagerMainMenuHandler))

	// cash manager replenishment requests menu
	h.Message(h.CashManagerHandler.RepReqMenuHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqMenuHandler))

	h.Message(h.CashManagerHandler.ActRepReqSelectHandler, h.isSessionStep(domain.SessionStepCashManagerActRepReqSelectHandler))
	h.Message(h.CashManagerHandler.ActRepReqActionHandler, h.isSessionStep(domain.SessionStepCashManagerActRepReqActionHandler))
	h.Message(h.CashManagerHandler.ActRepReqConfirmActionHandler, h.isSessionStep(domain.SessionStepCashManagerActRepReqConfirmActionHandler))
	h.Message(h.CashManagerHandler.RepReqAnotherAmountHandler, h.isSessionStep(domain.SessionStepCashManagerRepReqAnotherAmountHandler))

	h.Message(h.CashManagerHandler.ObjRepReqSelectHandler, h.isSessionStep(domain.SessionStepCashManagerObjRepReqSelectHandler))
	h.Message(h.CashManagerHandler.ObjRepReqAnotherAmountSelectionHandler, h.isSessionStep(domain.SessionStepCashManagerObjRepReqAnotherAmountSelectHandler))
}
