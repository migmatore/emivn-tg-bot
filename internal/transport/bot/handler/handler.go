package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/transport/bot/handler/admin"
	"emivn-tg-bot/internal/transport/bot/handler/daimyo"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"log"
	"time"
)

type AuthService interface {
	//CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	GetRole(ctx context.Context, username string) (string, error)
}

type ShogunService interface {
	Create(ctx context.Context, dto domain.ShogunDTO) error
	GetAll(ctx context.Context) ([]*domain.ShogunDTO, error)
}

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetByUsername(ctx context.Context, daimyoUsername string) ([]*domain.CardDTO, error)
}

// TODO: Refactor DI
type Deps struct {
	sessionManager *session.Manager[domain.Session]

	AuthService        AuthService
	ShogunService      ShogunService
	DaimyoService      DaimyoService
	SamuraiService     SamuraiService
	CashManagerService CashManagerService
	CardService        CardService
}

type Handler struct {
	*tgb.Router
	sessionManager *session.Manager[domain.Session]

	StartHandler  *start.StartHandler
	AdminHandler  *admin.AdminHandler
	DaimyoHandler *daimyo.DaimyoHandler
}

func New(deps Deps) *Handler {

	sm := NewSessionManager()

	return &Handler{
		Router:         tgb.NewRouter(),
		sessionManager: sm.Manager,
		StartHandler:   start.NewStartHandler(sm.Manager, deps.AuthService),
		AdminHandler: admin.NewAdminHandler(
			sm.Manager,
			deps.ShogunService,
			deps.DaimyoService,
			deps.SamuraiService,
			deps.CashManagerService,
			deps.CardService,
		),
		DaimyoHandler: daimyo.NewDaimyoHandler(sm.Manager, deps.CardService),
	}
}

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	h.Router.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
		return tgb.HandlerFunc(func(ctx context.Context, update *tgb.Update) error {
			defer func(started time.Time) {
				log.Printf("%#v [%s]", update, time.Since(started))
			}(time.Now())

			return next.Handle(ctx, update)
		})
	}))
	h.registerSession()
	h.registerStartHandlers()
	h.registerAdminHandlers()
	h.registerDaimyoHandler()

	//h.Message(h.StartHandler.Start, tgb.Command("start")).
	//	Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
	//		return mu.Update.Reply(ctx, mu.Answer("Напишите /start"))
	//	}, h.isSessionStep(domain.SessionStepInit)).
	//	Message(h.AdminHandler.Menu, h.isSessionStep(domain.SessionStepAdminRole)).
	//	Message(h.AdminHandler.MenuSelectionHandler, h.isSessionStep(domain.SessionStepAdminMenuHandler)).
	//	Message(h.AdminHandler.CreateEntityMenu, h.isSessionStep(domain.SessionStepCreateEntityButton)).
	//	Message(h.AdminHandler.CreateEntityMenuSelectionHandler, h.isSessionStep(domain.SessionStepCreateEntityHandler)).
	//	Message(h.AdminHandler.CreateShogun, h.isSessionStep(domain.SessionStepCreateShogun))

	return h.Router
}
