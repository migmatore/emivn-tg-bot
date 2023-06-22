package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterMainOperatorNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.MainOperator.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateMainOperatorUsername

	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *AdminHandler) EnterMainOperatorUsernameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.MainOperator.Username = strings.ReplaceAll(msg.Text, "@", "")

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range shoguns {
		buttons = append(buttons, tg.NewKeyboardButton(item.Nickname))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateMainOperator

	return msg.Answer("Выберите сёгуна, которому будет подчиняться главный оператор.").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *AdminHandler) CreateMainOperatorHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	shogun, err := h.shogunService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.MainOperator.ShogunUsername = shogun.Username

	if err := h.mainOperatorService.Create(ctx, sessionManager.MainOperator); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Главный оператор успешно создан.\nНапишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}
