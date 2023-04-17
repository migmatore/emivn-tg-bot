package bot

//func (bot *Bot) onStartCmd(ctx context.Context, msg *tgb.MessageUpdate) error {
//	//return msg.Update.Reply(ctx, bot.buildStartMsg().AsSendCall(msg.Chat.ID))
//
//	return msg.Answer("Hi, <b>username</b>!").ParseMode(tg.HTML).DoVoid(ctx)
//}
//
//func (bot *Bot) buildStartMsg() *tgx.TextMessage {
//	return tgx.NewTextMessage("Hi, <b>username</b>!").ParseMode(tg.HTML)
//}
//
//func (bot *Bot) registerGeneralHandlers() {
//	bot.Message(bot.onStartCmd,
//		tgb.Command("start"),
//		tgb.ChatType(tg.ChatTypePrivate),
//	)
//}
