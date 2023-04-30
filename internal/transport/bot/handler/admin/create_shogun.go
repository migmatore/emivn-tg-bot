package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterShogunUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	// TODO: create regular expression to check the username is correct
	h.shogun.Username = strings.ReplaceAll(msg.Text, "@", "")

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogun
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterShogunNicknameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.shogun.Nickname = msg.Text

	if err := h.shogunService.Create(ctx, h.shogun); err != nil {
		return err
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	return msg.Answer("Сёгун успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
