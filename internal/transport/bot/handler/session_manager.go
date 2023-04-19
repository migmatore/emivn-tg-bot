package handler

import (
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type SessionManager struct {
	*session.Manager[domain.Session]
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		session.NewManager(domain.Session{
			Step: domain.SessionStepStart,
		}),
	}
}

func (h *Handler) registerSession() {
	h.Router.Use(h.sessionManager)
}

func (h *Handler) isSessionStep(state domain.SessionStep) tgb.Filter {
	return h.sessionManager.Filter(func(session *domain.Session) bool {
		return session.Step == state
	})
}
