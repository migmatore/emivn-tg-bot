package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

func (h *CashManagerHandler) RepReqMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.CashManagerRepRequestsMenu.Active:
		requests, err := h.replenishmentRequestService.GetAllByCashManager(
			ctx,
			string(msg.From.Username),
			domain.ActiveRequest.String(),
		)
		if err != nil {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Активные запросы отсутствуют.\nНапишите /start").DoVoid(ctx)
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, request := range requests {
			card, _ := h.cardService.GetByName(ctx, request.CardName)

			var nickname string

			daimyo, _ := h.daimyoService.GetByUsername(ctx, card.OwnerUsername)
			if daimyo.Username == "" {
				operator, _ := h.mainOperatorService.GetByUsername(ctx, card.OwnerUsername)
				nickname = operator.Nickname
			} else {
				nickname = daimyo.Nickname
			}

			buttons = append(
				buttons,
				tg.NewKeyboardButton(fmt.Sprintf(
					"%s %d / %s / %d",
					card.Name,
					card.LastDigits,
					nickname,
					int(request.Amount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите тип запроса").ReplyMarkup(kb).DoVoid(ctx)

	case domain.CashManagerRepRequestsMenu.Objectionable:
		requests, err := h.replenishmentRequestService.GetAllByCashManager(
			ctx,
			string(msg.From.Username),
			domain.ObjectionableRequest.String(),
		)
		if err != nil {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Спорные запросы отсутствуют.\nНапишите /start").DoVoid(ctx)
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, request := range requests {
			card, _ := h.cardService.GetByName(ctx, request.CardName)

			var nickname string

			daimyo, _ := h.daimyoService.GetByUsername(ctx, card.OwnerUsername)
			if daimyo.Username == "" {
				operator, _ := h.mainOperatorService.GetByUsername(ctx, card.OwnerUsername)
				nickname = operator.Nickname
			} else {
				nickname = daimyo.Nickname
			}

			buttons = append(
				buttons,
				tg.NewKeyboardButton(fmt.Sprintf(
					"%s %d / %s / %d",
					card.Name,
					card.LastDigits,
					nickname,
					int(request.Amount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите тип запроса").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
