package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterShogunNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Shogun.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateShogun
	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *AdminHandler) EnterShogunUsernameAndCreateHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Shogun.Username = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.shogunService.Create(ctx, sessionManager.Shogun); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Сёгун успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
