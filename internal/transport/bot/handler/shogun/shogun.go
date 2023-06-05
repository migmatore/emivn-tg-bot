package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type ShogunHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func NewShogunHandler(sm *session.Manager[domain.Session]) *ShogunHandler {
	return &ShogunHandler{sessionManager: sm}
}

func (h *ShogunHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {

	return nil
}
