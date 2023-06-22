package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterControllerNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.Controller.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateController
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterControllerUsernameAndCreateHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Controller.Username = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.controllerService.Create(ctx, sessionManager.Controller); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Контролёр успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
