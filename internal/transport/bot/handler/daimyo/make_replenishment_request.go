package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
)

func (h *DaimyoHandler) RepReqChooseCardHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	cards, err := h.cardService.GetAllByUsername(ctx, msg.Text, string(msg.Chat.Username))
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range cards {
		buttons = append(buttons, tg.NewKeyboardButton(item.Name))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoEnterReplenishmentRequestAmount

	return msg.Answer("Введите название карты из списка, которую хотите пополнить:").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *DaimyoHandler) EnterRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.ReplenishmentRequest.CardName = msg.Text

	// TODO: Search the card in the replenishment requests table.
	// If the record exists, then reject the request.

	sessionManager.Step = domain.SessionStepDaimyoMakeReplenishmentRequest
	return msg.Answer("Введите сумму на пополнение").DoVoid(ctx)
}

func (h *DaimyoHandler) MakeRepReqHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.ParseFloat(msg.Text, 32)
	if err != nil {
		sessionManager.Step = domain.SessionStepDaimyoMakeReplenishmentRequest
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	sessionManager.ReplenishmentRequest.Amount = float32(amount)
	sessionManager.ReplenishmentRequest.DaimyoUsername = string(msg.From.Username)

	// TODO: Check the card's limit
	chatId, err := h.replenishmentRequestService.Create(ctx, sessionManager.ReplenishmentRequest)
	if err != nil {
		return err
	}

	if err := msg.Client.SendMessage(chatId, "RepReq").DoVoid(ctx); err != nil {
		// TODO
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Данные записаны. Напишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *DaimyoHandler) ChangeRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	return nil
}
