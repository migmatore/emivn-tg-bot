package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *ShogunHandler) EnterMainOperatorNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.MainOperator.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepShogunCreateMainOperator

	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *ShogunHandler) CreateMainOperatorHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.MainOperator.Username = strings.ReplaceAll(msg.Text, "@", "")
	sessionManager.MainOperator.ShogunUsername = string(msg.From.Username)

	if err := h.mainOperatorService.Create(ctx, sessionManager.MainOperator); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Главный оператор успешно создан.\nНапишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}
