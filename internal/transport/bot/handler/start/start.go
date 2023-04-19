package start

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type AuthService interface {
	Auth(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	Redirect(ctx context.Context, username string) domain.SessionStep
}

type StartHandler struct {
	sessionManager *session.Manager[domain.Session]

	AuthService AuthService
}

func NewStartHandler(sm *session.Manager[domain.Session], s AuthService) *StartHandler {
	return &StartHandler{
		sessionManager: sm, AuthService: s,
	}
}

func (s *StartHandler) Start(ctx context.Context, msg *tgb.MessageUpdate) error {
	//auth, _ := s.AuthService.Auth(ctx, string(msg.From.Username), domain.AdminRole)
	//
	//if auth {
	//	return msg.Answer("You are welcome!!!!!").DoVoid(ctx)
	//}
	//
	s.sessionManager.Get(ctx).Step = s.AuthService.Redirect(ctx, string(msg.From.Username))

	return nil
}
