package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/transport/bot/handler/admin"
	"emivn-tg-bot/internal/transport/bot/handler/daimyo"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

// TODO: Refactor DI
type Deps struct {
	sessionManager *session.Manager[domain.Session]

	AuthService        start.AuthService
	ShogunService      admin.ShogunService
	DaimyoService      admin.DaimyoService
	SamuraiService     admin.SamuraiService
	CashManagerService admin.CashManagerService
	CardService        admin.CardService
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
		DaimyoHandler: daimyo.NewDaimyoHandler(sm.Manager),
	}
}

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	h.registerSession()
	h.registerStartHandlers()
	h.registerAdminHandlers()

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
