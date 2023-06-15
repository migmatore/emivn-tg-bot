package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerSamuraiHandler() {
	h.Message(h.SamuraiHandler.EnterDataMenuHandler, h.isSessionStep(domain.SessionStepSamuraiEnterDataMenuHandler))
	h.Message(h.SamuraiHandler.ChooseBankMenuHandler, h.isSessionStep(domain.SessionStepSamuraiChooseBankMenuHandler))
}
