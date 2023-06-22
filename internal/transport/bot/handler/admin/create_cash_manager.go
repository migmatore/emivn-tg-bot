package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterCashManagerNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.CashManager.Nickname = msg.Text

	sessionManager.Step = domain.SessionStepCreateCashManagerUsername
	return msg.Answer("Введите telegram username").DoVoid(ctx)
}

func (h *AdminHandler) EnterCashManagerUsernameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.CashManager.Username = strings.ReplaceAll(msg.Text, "@", "")

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, shogun := range shoguns {
		buttons = append(buttons, tg.NewKeyboardButton(shogun.Nickname))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCreateCashManager
	return msg.Answer("Выберите сёгуна, к которому будет привязан инкассатор").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateCashManagerHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	shogun, err := h.shogunService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.CashManager.ShogunUsername = shogun.Username

	if err := h.cashManagerService.Create(ctx, sessionManager.CashManager); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Инкассатор успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
