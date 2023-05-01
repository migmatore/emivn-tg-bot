package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterCashManagerUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	// TODO: create regular expression to check the username is correct
	h.cashManager.Username = strings.ReplaceAll(msg.Text, "@", "")

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManager
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterCashManagerNicknameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.cashManager.Nickname = msg.Text

	if err := h.cashManagerService.Create(ctx, h.cashManager); err != nil {
		return err
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	return msg.Answer("Инкассатор успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
