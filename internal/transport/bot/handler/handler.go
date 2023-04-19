package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/transport/bot/handler/db_actions"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type Deps struct {
	sessionManager   *session.Manager[domain.Session]
	DbActionsService db_actions.DbActionsService
}

type Handler struct {
	*tgb.Router
	sessionManager *session.Manager[domain.Session]

	StartHandler   *start.StartHandler
	DbWriteHandler *db_actions.DbActionsHandler
}

func New(deps Deps) *Handler {

	sm := NewSessionManager()

	return &Handler{
		Router:         tgb.NewRouter(),
		sessionManager: sm.Manager,
		StartHandler:   start.NewStartHandler(),
		DbWriteHandler: db_actions.NewDbWriteHandler(sm.Manager, deps.DbActionsService),
	}
}

var (
	genders = []string{
		"Male",
		"Female",
		"Attack Helicopter",
		"Other",
	}
)

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	//bot.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
	//	return tgb.HandlerFunc(func(c context.Context, update *tgb.Update) error {
	//		ctx = logging.ContextWithLogger(ctx)
	//		return next.Handle(ctx, update)
	//	})
	//}))

	h.registerSession()
	h.registerDbActionsHandler()

	return h.Router
}
