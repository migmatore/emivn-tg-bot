package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
	"strings"
)

func (h *AdminHandler) CardBank(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.Card.BankType = msg.Text

	sessionManager.Step = domain.SessionStepCreateCardName

	return msg.Answer(fmt.Sprintf("Введите название карты")).
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *AdminHandler) EnterCardName(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Card.Name = msg.Text

	sessionManager.Step = domain.SessionStepCreateCardLastDigits
	return msg.Answer("Введите последние 4 цифры банковской карты").DoVoid(ctx)
}

func (h *AdminHandler) EnterCardLastDigits(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	lastDigits, err := strconv.Atoi(msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepCreateCardLastDigits
		return msg.Answer("Пожалуйста, введите последние 4 цифры банковской карты").DoVoid(ctx)
	}

	sessionManager.Card.LastDigits = lastDigits

	sessionManager.Step = domain.SessionStepCreateCardDailyLimit
	return msg.Answer("Введите суточный лимит карты").DoVoid(ctx)
}

func (h *AdminHandler) EnterCardDailyLimit(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	dailyLimit, err := strconv.Atoi(msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepCreateCardDailyLimit
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	sessionManager.Card.DailyLimit = dailyLimit

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	var str string

	for _, daimyo := range daimyos {
		str += "@" + daimyo.Username + "\n"
	}

	sessionManager.Step = domain.SessionStepCreateCard
	return msg.Answer(fmt.Sprintf("Введите username даймё, к которому будет привязана карта. \nСписок даёме: \n%s", str)).DoVoid(ctx)
}

func (h *AdminHandler) EnterCardDaimyoUsernameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Card.DaimyoUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.cardService.Create(ctx, sessionManager.Card); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Карта успешно создана. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
