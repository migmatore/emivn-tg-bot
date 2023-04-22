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

	//	service AdminService
}

func NewDbWriteHandler(sm *session.Manager[domain.Session], s AdminService) *AdminHandler {
	return &AdminHandler{
		sessionManager: sm,
		//		service:        s,
	}
}

func (h *AdminHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminMenu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateEntityHandler
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateShogun),
			tg.NewKeyboardButton(domain.AdminCreateEnityMenu.Back),
		)...,
	).WithResizeKeyboardMarkup()

	return msg.Answer(fmt.Sprintf("action selected: %s", msg.Text)).ReplyMarkup(kb).DoVoid(ctx)
}

func (h *AdminHandler) CreateEntityMenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminCreateEnityMenu.CreateShogun:
		return msg.Answer("create shogun").DoVoid(ctx)
	case domain.AdminCreateEnityMenu.Back:
		h.sessionManager.Get(ctx).Step = domain.SessionStepAdminMenuHandler
		return msg.Answer("Главное меню").ReplyMarkup(domain.Menu).DoVoid(ctx)
	}

	return nil
}

func (h *AdminHandler) CreateShogun(ctx context.Context, msg *tgb.MessageUpdate) error {
	return nil
}
