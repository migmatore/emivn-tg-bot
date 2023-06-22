package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterDaimyoNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateDaimyoUsername
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterDaimyoUsernameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Daimyo.Username = strings.ReplaceAll(msg.Text, "@", "")

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range shoguns {
		buttons = append(buttons, tg.NewKeyboardButton(item.Nickname))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateDaimyo

	return msg.Answer("Выберите сёгуна, которому будет подчиняться даймё.").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateDaimyoHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	shogun, err := h.shogunService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.Daimyo.ShogunUsername = shogun.Username

	if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Даймё успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
