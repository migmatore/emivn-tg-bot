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
		s.sessionManager.Get(ctx).Step = domain.SessionStepAdminMainMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildAdminStartMenu()).
			DoVoid(ctx)
	case domain.DaimyoRole.String():
		s.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)
	case domain.SamuraiRole.String():
		s.sessionManager.Get(ctx).Step = domain.SessionStepSamuraiMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)
	case domain.CashManagerRole.String():
		s.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)
	default:
		return nil
	}
}
