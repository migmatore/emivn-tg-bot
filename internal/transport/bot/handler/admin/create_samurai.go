package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterSamuraiUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	// TODO: create regular expression to check the username is correct
	sessionManager.Samurai.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepCreateSamuraiNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterSamuraiNickname(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Nickname = msg.Text

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range daimyos {
		buttons = append(buttons, tg.NewKeyboardButton(item.Username))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateSamurai

	return msg.Answer("Введите username даймё, которому будет подчиняться самурай.").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateSamurai(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Samurai.DaimyoUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Самурай успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
