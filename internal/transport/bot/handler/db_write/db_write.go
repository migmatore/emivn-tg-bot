package db_write

import (
	"context"
	"emivn-tg-bot/pkg/logging"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type DbWriteHandler struct {
}

type Menu struct {
	Write string
	Read  string
}

func NewDbWriteHandler() *DbWriteHandler {
	return &DbWriteHandler{}
}

func (h *DbWriteHandler) Menu(ctx context.Context, msg *tgb.MessageUpdate) error {
	menu := Menu{
		Write: "Write",
		Read:  "Read",
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(menu.Write),
			tg.NewKeyboardButton(menu.Read),
		)...,
	).WithResizeKeyboardMarkup()

	return msg.Answer("Hey, please click a button above").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *DbWriteHandler) Write(ctx context.Context, msg *tgb.MessageUpdate) error {
	return msg.Answer("db write").DoVoid(ctx)
}

func (h *DbWriteHandler) Read(ctx context.Context, msg *tgb.MessageUpdate) error {
	logging.GetLogger(ctx).Infof("%s", msg.Text)

	return msg.Update.Reply(ctx, msg.Answer("write data"))
}
