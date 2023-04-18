package db_write

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type DbWriteHandler struct {
	sessionManager *session.Manager[domain.Session]
}

type Menu struct {
	Write string
	Read  string
}

func NewDbWriteHandler(sessionManager *session.Manager[domain.Session]) *DbWriteHandler {
	return &DbWriteHandler{sessionManager: sessionManager}
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

	h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect

	return msg.Answer("Hey, please click a button above").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *DbWriteHandler) ActionSelect(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case "Read":
		h.sessionManager.Get(ctx).Step = domain.SessionStepReadData
	case "Write":
		h.sessionManager.Get(ctx).Step = domain.SessionStepWriteData
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect
	}
	return msg.Answer(fmt.Sprintf("action selected: %s", msg.Text)).DoVoid(ctx)
}

func (h *DbWriteHandler) Read(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect

	return msg.Answer(fmt.Sprintf("read", msg.Text)).DoVoid(ctx)
}

func (h *DbWriteHandler) Write(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.sessionManager.Get(ctx).Step = domain.SessionStepInit

	return msg.Answer(fmt.Sprintf("write", msg.Text)).DoVoid(ctx)
}
