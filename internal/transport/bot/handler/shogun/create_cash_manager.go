package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *ShogunHandler) EnterCashManagerNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.CashManager.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepShogunCreateCashManager

	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *ShogunHandler) CreateCashManagerHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.CashManager.Username = strings.ReplaceAll(msg.Text, "@", "")
	sessionManager.CashManager.ShogunUsername = string(msg.From.Username)

	if err := h.cashManagerService.Create(ctx, sessionManager.CashManager); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Инкассатор успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
