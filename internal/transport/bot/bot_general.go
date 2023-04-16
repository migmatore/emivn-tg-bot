package bot

import (
	"context"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg-bot/pkg/tgx"
	"github.com/mr-linch/go-tg/tgb"
)

func (bot *Bot) onStartCmd(ctx context.Context, msg *tgb.MessageUpdate) error {
	return msg.Update.Reply(ctx, bot.buildStartMsg().AsSendCall(msg.Chat.ID))
}

func (bot *Bot) buildStartMsg() *tgx.TextMessage {
	return tgx.NewTextMessage("Hi, <b>username</b>!").ParseMode(tg.HTML)
}

func (bot *Bot) registerGeneralHandlers() {
	bot.Message(bot.onStartCmd,
		tgb.Command("start"),
		tgb.ChatType(tg.ChatTypePrivate),
	)
}
