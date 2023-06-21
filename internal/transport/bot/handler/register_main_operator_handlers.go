package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerMainOperatorHandlers() {
	h.Message(h.MainOperatorHandler.MainMenuHandler, h.isSessionStep(domain.SessionStepMainOperatorMainMenuHandler))

	// replenishment requests
	h.Message(h.MainOperatorHandler.RepReqChooseCardHandler, h.isSessionStep(domain.SessionStepMainOperatorChooseReplenishmentRequestBank))
	h.Message(h.MainOperatorHandler.EnterRepReqAmountHandler, h.isSessionStep(domain.SessionStepMainOperatorEnterReplenishmentRequestAmount))
	h.Message(h.MainOperatorHandler.MakeRepReqHandler, h.isSessionStep(domain.SessionStepMainOperatorMakeReplenishmentRequest))
	h.Message(h.MainOperatorHandler.ChangeRepReqAmountHandler, h.isSessionStep(domain.SessionStepMainOperatorChangeReplenishmentRequestAmount))
}
