package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterDaimyoUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepCreateDaimyoNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterDaimyoNickname(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Daimyo.Nickname = msg.Text

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range shoguns {
		buttons = append(buttons, tg.NewKeyboardButton(item.Username))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateDaimyo

	return msg.Answer("Введите username сёгуна, которому будет подчиняться даймё.").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateDaimyo(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.ShogunUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Даймё успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
