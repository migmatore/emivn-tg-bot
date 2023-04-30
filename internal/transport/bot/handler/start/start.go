package start

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type AuthService interface {
	//CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	GetRole(ctx context.Context, username string) (string, error)
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
	role, err := s.AuthService.GetRole(ctx, string(msg.Chat.Username))
	if err != nil {
		return msg.Answer("Error").DoVoid(ctx)
	}

	switch role {
	case domain.AdminRole.String():
		s.sessionManager.Get(ctx).Step = domain.SessionStepAdminMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildAdminStartMenu()).
			DoVoid(ctx)
	default:
		return nil
	}
}
