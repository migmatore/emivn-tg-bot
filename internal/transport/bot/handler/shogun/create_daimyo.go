package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *ShogunHandler) EnterDaimyoNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepShogunCreateDaimyo

	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *ShogunHandler) CreateDaimyoHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Username = strings.ReplaceAll(msg.Text, "@", "")
	sessionManager.Daimyo.ShogunUsername = string(msg.From.Username)

	if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Даймё успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
