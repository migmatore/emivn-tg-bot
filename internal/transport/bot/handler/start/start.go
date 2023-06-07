package start

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"time"
)

type AuthService interface {
	//CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	GetRole(ctx context.Context, username string) (string, error)
}

type SamuraiService interface {
	SetChatId(ctx context.Context, username string, id tg.ChatID) error
}

type CashManagerService interface {
	SetChatId(ctx context.Context, username string, id tg.ChatID) error
}

type Scheduler interface {
	Add(ctx context.Context, dto domain.TaskDTO) error
}

type StartHandler struct {
	sessionManager *session.Manager[domain.Session]

	AuthService        AuthService
	SamuraiService     SamuraiService
	CashManagerService CashManagerService
	scheduler          Scheduler
}

func NewStartHandler(
	sm *session.Manager[domain.Session],
	authService AuthService,
	samuraiService SamuraiService,
	cashManagerService CashManagerService,
	scheduler Scheduler,
) *StartHandler {
	return &StartHandler{
		sessionManager:     sm,
		AuthService:        authService,
		SamuraiService:     samuraiService,
		CashManagerService: cashManagerService,
		scheduler:          scheduler,
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

		if err := s.scheduler.Add(ctx, domain.TaskDTO{
			Alias:           "notify_samurai",
			Name:            "test task",
			Arguments:       nil,
			IntervalMinutes: 1,
			RunAt:           time.Now(),
		}); err != nil {
			return msg.Answer(err.Error()).DoVoid(ctx)
		}

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

		if err := s.SamuraiService.SetChatId(ctx, string(msg.Chat.Username), msg.Chat.ID); err != nil {
			s.sessionManager.Reset(s.sessionManager.Get(ctx))

			return msg.Answer("Ошибка").DoVoid(ctx)
		}

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)
	case domain.CashManagerRole.String():
		s.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerMenuHandler

		if err := s.CashManagerService.SetChatId(ctx, string(msg.Chat.Username), msg.Chat.ID); err != nil {
			s.sessionManager.Reset(s.sessionManager.Get(ctx))

			return msg.Answer("Ошибка").DoVoid(ctx)
		}

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)
	default:
		return nil
	}
}
