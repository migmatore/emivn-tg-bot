package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterControllerUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Controller.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepCreateController
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterControllerNicknameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Controller.Nickname = msg.Text

	if err := h.controllerService.Create(ctx, sessionManager.Controller); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Контролёр успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
