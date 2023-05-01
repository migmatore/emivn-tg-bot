package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
)

func (h *AdminHandler) EnterCardName(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.card.Name = msg.Text

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardLastDigits
	return msg.Answer("Введите последние 4 цифры банковской карты").DoVoid(ctx)
}

func (h *AdminHandler) EnterCardLastDigits(ctx context.Context, msg *tgb.MessageUpdate) error {
	lastDigits, err := strconv.Atoi(msg.Text)
	if err != nil {
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardLastDigits
		return msg.Answer("Пожалуйста, введите последние 4 цифры банковской карты").DoVoid(ctx)
	}

	h.card.LastDigits = lastDigits

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardDailyLimit
	return msg.Answer("Введите суточный лимит карты").DoVoid(ctx)
}

func (h *AdminHandler) EnterCardDailyLimit(ctx context.Context, msg *tgb.MessageUpdate) error {
	dailyLimit, err := strconv.Atoi(msg.Text)
	if err != nil {
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardDailyLimit
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	h.card.DailyLimit = dailyLimit

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	var str string

	for _, daimyo := range daimyos {
		str += "@" + daimyo.Username + "\n"
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCard
	return msg.Answer(fmt.Sprintf("Введите username даймё, к которому будет привязана карта. \nСписок даёме: \n%s", str)).DoVoid(ctx)
}

func (h *AdminHandler) EnterCardDaimyoUsernameAndCreate(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.card.DaimyoUsername = msg.Text

	if err := h.cardService.Create(ctx, h.card); err != nil {
		return err
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	return msg.Answer("Карта успешно создана. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
