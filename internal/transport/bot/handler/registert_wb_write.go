package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
)

func (h *Handler) registerDbWriteHandler() {
	h.Message(h.DbWriteHandler.Menu, tgb.Command("db_menu")).
		Message(h.DbWriteHandler.ActionSelect, h.isSessionStep(domain.SessionStepAcionSelect)).
		Message(h.DbWriteHandler.Read, h.isSessionStep(domain.SessionStepReadData)).
		Message(h.DbWriteHandler.Write, h.isSessionStep(domain.SessionStepWriteData)).
		Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
			// handle no command with SessionStepInitial
			return mu.Update.Reply(ctx, mu.Answer("Press /db_menu"))
		}, h.isSessionStep(domain.SessionStepInit))
}
