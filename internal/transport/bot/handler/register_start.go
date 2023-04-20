package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
)

func (h *Handler) registerStartHandler() {
	h.Message(h.StartHandler.Start, tgb.Command("start")).
		Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
			return mu.Update.Reply(ctx, mu.Answer("Напишите /start"))
		}, h.isSessionStep(domain.SessionStepInit))
}
