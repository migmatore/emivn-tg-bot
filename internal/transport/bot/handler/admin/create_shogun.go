package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterShogunUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Shogun.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepCreateShogun
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterShogunNicknameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Shogun.Nickname = msg.Text

	if err := h.shogunService.Create(ctx, sessionManager.Shogun); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Сёгун успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
