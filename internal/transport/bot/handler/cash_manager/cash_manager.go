package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type CashManagerHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func NewCashManagerHandler(sm *session.Manager[domain.Session]) *CashManagerHandler {
	return &CashManagerHandler{sessionManager: sm}
}

func (h *CashManagerHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {

	return nil
}
