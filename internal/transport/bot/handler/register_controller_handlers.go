package handler

import "emivn-tg-bot/internal/domain"

func (h *Handler) registerControllerHandler() {
	h.Message(h.ControllerHandler.EnterDataMenuHandler, h.isSessionStep(domain.SessionStepControllerEnterDataMenuHandler))
	h.Message(h.ControllerHandler.ChooseDaimyoMenuHandler, h.isSessionStep(domain.SessionStepControllerChooseDaimyoMenuHandler))
	h.Message(h.ControllerHandler.ChooseSamuraiMenuHandler, h.isSessionStep(domain.SessionStepControllerChooseSamuraiMenuHandler))
	h.Message(h.ControllerHandler.ChooseBankMenuHandler, h.isSessionStep(domain.SessionStepControllerChooseBankMenuHandler))
	h.Message(h.ControllerHandler.CreateTurnoverMenuHandler, h.isSessionStep(domain.SessionStepControllerCreateTurnoverHandler))
}
