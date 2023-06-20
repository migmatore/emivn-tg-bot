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
			return msg.Answer("Активные запросы отсутствуют").DoVoid(ctx)
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, request := range requests {
			card, _ := h.cardService.GetByName(ctx, request.CardName)

			daimyo, _ := h.daimyoService.GetByUsername(ctx, card.DaimyoUsername)

			buttons = append(
				buttons,
				tg.NewKeyboardButton(fmt.Sprintf(
					"%s %d / %s / %d",
					card.Name,
					card.LastDigits,
					daimyo.Nickname,
					int(request.Amount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите запрос").ReplyMarkup(kb).DoVoid(ctx)
	case domain.CashManagerRepRequestsMenu.Objectionable:
		return nil
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
