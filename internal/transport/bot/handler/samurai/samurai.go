package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type SamuraiHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func NewSamuraiHandler(sm *session.Manager[domain.Session]) *SamuraiHandler {
	return &SamuraiHandler{sessionManager: sm}
}

func (h *SamuraiHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {

	return nil
}
