package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *ShogunHandler) EnterSamuraiNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Samurai.Nickname = msg.Text

	daimyos, err := h.daimyoService.GetAllByShogun(ctx, string(msg.From.Username))
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

	sessionManager.Step = domain.SessionStepShogunChooseSamuraiDaimyo

	return msg.Answer("Выберите даймё, которому будет подчиняться самурай").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ShogunHandler) ChooseSamuraiDaimyoHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.DaimyoUsername = msg.Text

	sessionManager.Step = domain.SessionStepShogunCreateSamurai
	return msg.Answer("Введите telegram username").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}

func (h *ShogunHandler) CreateSamuraiHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Samurai.Username = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Самурай успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
