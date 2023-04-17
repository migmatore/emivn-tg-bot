package handler

import (
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type Deps struct {
}

type Handler struct {
	*tgb.Router
	StartHandler *StartHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Router:       tgb.NewRouter(),
		StartHandler: NewStartHandler(),
	}
}

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	//bot.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
	//	return tgb.HandlerFunc(func(c context.Context, update *tgb.Update) error {
	//		ctx = logging.ContextWithLogger(ctx)
	//		return next.Handle(ctx, update)
	//	})
	//}))

	h.Router.Message(h.StartHandler.Start, tgb.Command("start"), tgb.ChatType(tg.ChatTypePrivate))

	return h.Router
}
