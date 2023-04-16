package bot

import (
	"context"
	"github.com/mr-linch/go-tg/tgb"
)

type Bot struct {
	*tgb.Router
}

func New(ctx context.Context) (*Bot, error) {
	bot := &Bot{Router: tgb.NewRouter()}

	//bot.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
	//	return tgb.HandlerFunc(func(c context.Context, update *tgb.Update) error {
	//		ctx = logging.ContextWithLogger(ctx)
	//		return next.Handle(ctx, update)
	//	})
	//}))

	bot.registerHandlers()

	return bot, nil

}

func (bot *Bot) registerHandlers() *Bot {
	bot.registerGeneralHandlers()

	return bot
}
