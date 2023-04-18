package db_actions

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/logging"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type DbActionsService interface {
	DoAction(ctx context.Context, actionName string, text string) ([]string, error)
}

type DbActionsHandler struct {
	sessionManager *session.Manager[domain.Session]

	service DbActionsService
}

type Menu struct {
	Write string
	Read  string
}

func NewDbWriteHandler(sessionManager *session.Manager[domain.Session], s DbActionsService) *DbActionsHandler {
	return &DbActionsHandler{sessionManager: sessionManager, service: s}
}

func (h *DbActionsHandler) Menu(ctx context.Context, msg *tgb.MessageUpdate) error {
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

func (h *DbActionsHandler) ActionSelect(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case "Read":
		h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect

		strs, err := h.service.DoAction(ctx, "Read", "")
		if err != nil {
			logging.GetLogger(ctx).Errorf("service error %s", err)
		}

		str := ""
		for _, s := range strs {
			str += s + "\n"
		}

		return msg.Answer(fmt.Sprintf("readed \n%s", tg.HTML.Code(str))).ParseMode(tg.HTML).DoVoid(ctx)
	case "Write":
		h.sessionManager.Get(ctx).Step = domain.SessionStepWriteData
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect
	}
	return msg.Answer(fmt.Sprintf("action selected: %s", msg.Text)).DoVoid(ctx)
}

func (h *DbActionsHandler) Read(ctx context.Context, msg *tgb.MessageUpdate) error {
	//h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect
	//
	//txt, err := h.service.DoAction(ctx, "Read", "")
	//if err != nil {
	//	logging.GetLogger(ctx).Errorf("service error %s", err)
	//}

	return msg.Answer(fmt.Sprintf("readed")).DoVoid(ctx)
}

func (h *DbActionsHandler) Write(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.sessionManager.Get(ctx).Step = domain.SessionStepAcionSelect

	_, err := h.service.DoAction(ctx, "Write", msg.Text)
	if err != nil {
		logging.GetLogger(ctx).Errorf("service error %s", err)
	}

	return msg.Answer(fmt.Sprintf("writed", msg.Text)).DoVoid(ctx)
}
