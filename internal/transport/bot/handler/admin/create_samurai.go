package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterSamuraiNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.Samurai.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateSamuraiUsername
	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *AdminHandler) EnterSamuraiUsernameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Username = strings.ReplaceAll(msg.Text, "@", "")

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, daimyo := range daimyos {
		buttons = append(buttons, tg.NewKeyboardButton(daimyo.Nickname))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateSamurai

	return msg.Answer("Введите имя даймё, которому будет подчиняться самурай.").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateSamuraiHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	daimyo, err := h.daimyoService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.Samurai.DaimyoUsername = daimyo.Username

	if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Самурай успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
