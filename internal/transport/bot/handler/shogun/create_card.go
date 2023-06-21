package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
)

func (h *ShogunHandler) ChooseCardBankMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Card.BankType = msg.Text

	sessionManager.Step = domain.SessionStepShogunEnterCardNameHandler

	return msg.Answer("Введите имя карты").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}

func (h *ShogunHandler) EnterCardNameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Card.Name = msg.Text

	sessionManager.Step = domain.SessionStepShogunEnterCardLastDigitsHandler

	return msg.Answer("Введите 4 последних цифры номера карты").DoVoid(ctx)
}

func (h *ShogunHandler) EnterCardLastDigitsHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	lastDigits, err := strconv.Atoi(msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepShogunEnterCardLastDigitsHandler
		return msg.Answer("Пожалуйста, введите последние 4 цифры номера карты").DoVoid(ctx)
	}

	sessionManager.Card.LastDigits = lastDigits
	sessionManager.Step = domain.SessionStepShogunSetCardLimitHandler

	return msg.Answer("Установите лимит").DoVoid(ctx)
}

func (h *ShogunHandler) SetCardLimitHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	limit, err := strconv.Atoi(msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepShogunSetCardLimitHandler
		return msg.Answer("Пожалуйста, установите лимит").DoVoid(ctx)
	}

	sessionManager.Card.DailyLimit = limit

	daimyos, err := h.daimyoService.GetAllByShogun(ctx, string(msg.From.Username))
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, daimyo := range daimyos {
		buttons = append(buttons, tg.NewKeyboardButton(daimyo.Nickname))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepShogunChooseCardDaimyoHandler

	return msg.Answer("Выберите имя дайме, к которому будет привязана карта").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ShogunHandler) ChooseCardDaimyoAndCreateHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	daimyo, err := h.daimyoService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.Card.OwnerUsername = daimyo.Username

	if err := h.cardService.Create(ctx, sessionManager.Card); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Данные записаны.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
