package handler

import (
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
)

func (h *Handler) registerStartHandler() {
	h.Message(h.StartHandler.Start, tgb.Command("start"), h.isSessionStep(domain.SessionStepInit))
}
