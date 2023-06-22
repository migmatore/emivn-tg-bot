package start

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
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

func (h *StartHandler) Start(ctx context.Context, msg *tgb.MessageUpdate) error {
	role, err := h.AuthService.GetRole(ctx, string(msg.Chat.Username))
	if err != nil {
		return msg.Answer("Error").DoVoid(ctx)
	}

	switch role {
	case domain.AdminRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepAdminMainMenuHandler

		//if err := s.scheduler.Add(ctx, domain.TaskDTO{
		//	Alias:           "notify_samurai",
		//	Name:            "test task",
		//	Arguments:       domain.FuncArgs{"id": 6109520093},
		//	IntervalMinutes: 0,
		//	RunAt:           time.Now(),
		//}); err != nil {
		//	return msg.Answer(err.Error()).DoVoid(ctx)
		//}
		//
		//if err := s.scheduler.Add(ctx, domain.TaskDTO{
		//	Alias:           "notify_samurai",
		//	Name:            "test task 2",
		//	Arguments:       domain.FuncArgs{"id": 1093658711},
		//	IntervalMinutes: 0,
		//	RunAt:           time.Now(),
		//}); err != nil {
		//	return msg.Answer(err.Error()).DoVoid(ctx)
		//}

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildAdminStartMenu()).
			DoVoid(ctx)

	case domain.ShogunRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunMainMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildShogunStartMenu()).
			DoVoid(ctx)

	case domain.DaimyoRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoMainMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildDaimyoStartMenu()).
			DoVoid(ctx)

	case domain.SamuraiRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepSamuraiEnterDataMenuHandler

		if err := h.SamuraiService.SetChatId(ctx, string(msg.Chat.Username), msg.Chat.ID); err != nil {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Ошибка").DoVoid(ctx)
		}

		return msg.Answer("Введите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").
			ReplyMarkup(buildSamuraiStartMenu()).
			DoVoid(ctx)

	case domain.CashManagerRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerMainMenuHandler

		if err := h.CashManagerService.SetChatId(ctx, string(msg.Chat.Username), msg.Chat.ID); err != nil {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Ошибка").DoVoid(ctx)
		}

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildCashManagerStartMenu()).
			DoVoid(ctx)

	case domain.ControllerRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepControllerEnterDataMenuHandler

		return msg.Answer("Введите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").
			ReplyMarkup(buildControllerStartMenu()).
			DoVoid(ctx)

	case domain.MainOperatorRole.String():
		h.sessionManager.Get(ctx).Step = domain.SessionStepMainOperatorMainMenuHandler

		return msg.Answer("Пожалуйста, выберите действие").
			ReplyMarkup(buildMainOperatorStartMenu()).
			DoVoid(ctx)

	default:
		return nil
	}
}
