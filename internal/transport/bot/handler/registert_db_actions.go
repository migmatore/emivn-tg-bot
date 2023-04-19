package handler

import (
	"emivn-tg-bot/internal/domain"
)

func (h *Handler) registerDbActionsHandler() {
	h.Message(h.DbWriteHandler.Menu, h.isSessionStep(domain.SessionStepStart)).
		Message(h.DbWriteHandler.ActionSelect, h.isSessionStep(domain.SessionStepAcionSelect)).
		Message(h.DbWriteHandler.Read, h.isSessionStep(domain.SessionStepReadData)).
		Message(h.DbWriteHandler.Write, h.isSessionStep(domain.SessionStepWriteData))
	//		Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
	//			// handle no command with SessionStepInitial
	//			return mu.Update.Reply(ctx, mu.Answer("Press /db_menu"))
	//		}, h.isSessionStep(domain.SessionStepInit))
}
