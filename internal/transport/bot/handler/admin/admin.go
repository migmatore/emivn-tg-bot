package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type AdminService interface {
	DoAction(ctx context.Context, actionName string, text string) ([]string, error)
}

type Menu struct {
	CreateEntity string
}

type AdminHandler struct {
	sessionManager *session.Manager[domain.Session]
	menu           Menu

	//	service AdminService
}

func NewDbWriteHandler(sm *session.Manager[domain.Session], s AdminService) *AdminHandler {
	return &AdminHandler{
		menu:           Menu{CreateEntity: "CreateEntity"},
		sessionManager: sm,
		//		service:        s,
	}
}

func (h *AdminHandler) Menu(ctx context.Context, msg *tgb.MessageUpdate) error {
	//menu := Menu{
	//	Write: "Write",
	//	Read:  "Read",
	//}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(h.menu.CreateEntity),
		)...,
	).WithResizeKeyboardMarkup()

	h.sessionManager.Get(ctx).Step = domain.SessionStepAdminMenuHandler

	return msg.Answer("Hey, please click a button above").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case h.menu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepBackCreateEntityMenuStep
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	}

	return msg.Answer(fmt.Sprintf("action selected: %s", msg.Text)).DoVoid(ctx)
}

func (h *AdminHandler) CreateEntityMenu(ctx context.Context, msg *tgb.MessageUpdate) error {
	//kb := tg.NewReplyKeyboardMarkup(
	//	tg.NewButtonColumn(
	//		tg.NewKeyboardButton(h.menu.CreateEntity),
	//	)...,
	//).WithResizeKeyboardMarkup()
	//
	//h.sessionManager.Get(ctx).Step = domain.SessionStepAdminMenuHandler
	//
	//return msg.Answer("Hey, please click a button above").
	//	ReplyMarkup(kb).
	//	DoVoid(ctx)
	return nil
}

func (h *AdminHandler) CreateEntityMenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	return nil
}

func (h *AdminHandler) CreateShogun(ctx context.Context, msg *tgb.MessageUpdate) error {
	return nil
}
