package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
)

func (h *DaimyoHandler) EnterCardName(ctx context.Context, msg *tgb.MessageUpdate) error {
	cards, err := h.cardService.GetByUsername(ctx, msg.Text, string(msg.Chat.Username))
	if err != nil {
		return err
	}

	var str string

	for i, card := range cards {
		str += fmt.Sprintf("%d. %s\n", i+1, card.Name)
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepEnterReplenishmentRequestAmount

	return msg.Answer(fmt.Sprintf("Введите название карты из списка, которую хотите пополнить: \n%s", str)).
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *DaimyoHandler) EnterReplenishmentRequestAmount(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.ReplenishmentRequest.CardName = msg.Text

	sessionManager.Step = domain.SessionStepMakeReplenishmentRequest
	return msg.Answer("Введите сумму на пополнение").DoVoid(ctx)
}

func (h *DaimyoHandler) MakeReplenishmentRequest(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.ParseFloat(msg.Text, 32)
	if err != nil {
		sessionManager.Step = domain.SessionStepMakeReplenishmentRequest
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	sessionManager.ReplenishmentRequest.Amount = float32(amount)
	sessionManager.ReplenishmentRequest.DaimyoUsername = string(msg.From.Username)

	chatId, err := h.replenishmentRequestService.Create(ctx, sessionManager.ReplenishmentRequest)
	if err != nil {
		return err
	}

	if err := msg.Client.SendMessage(chatId, "RepReq").DoVoid(ctx); err != nil {
		// TODO
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Запрос на пополнение успешно создан. Напишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}
