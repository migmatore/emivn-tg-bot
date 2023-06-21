package main_operator

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
)

func (h *MainOperatorHandler) RepReqChooseCardHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	cards, err := h.cardService.GetAllByUsername(ctx, msg.Text, string(msg.From.Username))
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

	h.sessionManager.Get(ctx).Step = domain.SessionStepMainOperatorEnterReplenishmentRequestAmount

	return msg.Answer("Введите название карты из списка, которую хотите пополнить:").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *MainOperatorHandler) EnterRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.ReplenishmentRequest.CardName = msg.Text

	// If the record exists, then reject the request.
	exists, err := h.replenishmentRequestService.CheckIfExists(ctx, msg.Text)
	if err != nil {
		return err
	}

	if exists {
		h.sessionManager.Reset(sessionManager)
		return msg.Answer("Вы не можете использовать данную карту, так как она активна/в споре.\nНапишите /start").
			DoVoid(ctx)
	}

	sessionManager.Step = domain.SessionStepMainOperatorMakeReplenishmentRequest
	return msg.Answer("Введите сумму на пополнение").DoVoid(ctx)
}

func (h *MainOperatorHandler) MakeRepReqHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.ParseFloat(msg.Text, 32)
	if err != nil {
		sessionManager.Step = domain.SessionStepMainOperatorMakeReplenishmentRequest
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	sessionManager.ReplenishmentRequest.Amount = float32(amount)
	sessionManager.ReplenishmentRequest.OwnerUsername = string(msg.From.Username)

	card, err := h.cardService.GetByUsername(ctx, string(msg.From.Username))
	if err != nil {
		return err
	}

	if float32(amount) > float32(card.DailyLimit) {
		sessionManager.Step = domain.SessionStepMainOperatorChangeReplenishmentRequestAmount

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton("Ввести другую сумму"),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer(fmt.Sprintf("Вы превысили лимит. Остаток по карте:\n%d", card.DailyLimit)).
			ReplyMarkup(kb).
			DoVoid(ctx)
	}

	chatId, err := h.replenishmentRequestService.Create(ctx, sessionManager.ReplenishmentRequest)
	if err != nil {
		return err
	}

	operator, err := h.mainOperatorService.GetByUsername(ctx, string(msg.From.Username))
	if err != nil {
		return err
	}

	msg.Client.SendMessage(chatId,
		fmt.Sprintf("Появилась новая заявка на пополнение: %s / %s %d", operator.Nickname, card.Name, card.LastDigits)).
		DoVoid(ctx)

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Данные записаны.\n Напишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *MainOperatorHandler) ChangeRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case "Ввести другую сумму":
		h.sessionManager.Get(ctx).Step = domain.SessionStepMainOperatorMakeReplenishmentRequest
		return msg.Answer("Введите сумму на пополнение").DoVoid(ctx)
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
