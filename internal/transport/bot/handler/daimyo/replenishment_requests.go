package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
	"strings"
)

func (h *DaimyoHandler) RepReqMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoRepRequestsMenu.Active:
		requests, err := h.replenishmentRequestService.GetAllByOwner(
			ctx,
			string(msg.From.Username),
			domain.ActiveRequests.String(),
		)
		if err != nil {
			return err
		}

		if len(requests) == 0 {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Активные запросы отсутствуют.\nНапишите /start").
				ReplyMarkup(tg.NewReplyKeyboardRemove()).
				DoVoid(ctx)
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, request := range requests {
			card, _ := h.cardService.GetByName(ctx, request.CardName)

			buttons = append(
				buttons,
				tg.NewKeyboardButton(fmt.Sprintf(
					"%s %d / %d",
					card.Name,
					card.LastDigits,
					int(request.RequiredAmount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		//h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoActRepReqSelectHandler

		return msg.Answer("Выберите запрос").ReplyMarkup(kb).DoVoid(ctx)

	case domain.DaimyoRepRequestsMenu.Objectionable:
		requests, err := h.replenishmentRequestService.GetAllByOwner(
			ctx,
			string(msg.From.Username),
			domain.ObjectionableRequests.String(),
		)
		if err != nil {
			return err
		}

		if len(requests) == 0 {
			h.sessionManager.Reset(h.sessionManager.Get(ctx))

			return msg.Answer("Спорные запросы отсутствуют.\nНапишите /start").
				ReplyMarkup(tg.NewReplyKeyboardRemove()).
				DoVoid(ctx)
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, request := range requests {
			card, _ := h.cardService.GetByName(ctx, request.CardName)

			buttons = append(
				buttons,
				tg.NewKeyboardButton(fmt.Sprintf(
					"%s %d / %d",
					card.Name,
					card.LastDigits,
					int(request.RequiredAmount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoObjRepReqSelectHandler

		return msg.Answer("Выберите запрос").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) ObjRepReqSelectHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	s := strings.Split(msg.Text, "/")
	cardMsg := strings.Split(s[0], " ")

	sessionManager.ReplenishmentRequest.CardName = cardMsg[0]

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.DaimyoConfirmRepRequestsMenu.Replenished),
			tg.NewKeyboardButton(domain.DaimyoConfirmRepRequestsMenu.ReplenishedAnotherAmount),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepDaimyoObjRepReqActionHandler

	return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *DaimyoHandler) ObjRepReqActionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	request, err := h.replenishmentRequestService.GetByCardName(ctx, sessionManager.ReplenishmentRequest.CardName)
	if err != nil {
		return err
	}

	sessionManager.ReplenishmentRequest = request

	switch msg.Text {
	case domain.DaimyoConfirmRepRequestsMenu.Replenished:
		sessionManager.ReplenishmentRequest.ActualAmount = sessionManager.ReplenishmentRequest.RequiredAmount

		if err := h.replenishmentRequestService.ConfirmRequest(ctx, sessionManager.ReplenishmentRequest); err != nil {
			return err
		}

		h.sessionManager.Reset(sessionManager)

		return msg.Answer("Данные записаны.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	case domain.DaimyoConfirmRepRequestsMenu.ReplenishedAnotherAmount:
		sessionManager.Step = domain.SessionStepDaimyoRepReqAnotherAmountHandler

		return msg.Answer("Введите сумму").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) ObjRepReqAnotherAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.Atoi(msg.Text)
	if err != nil {
		return err
	}

	sessionManager.ReplenishmentRequest.ActualAmount = float32(amount)

	if err := h.replenishmentRequestService.ConfirmRequest(ctx, sessionManager.ReplenishmentRequest); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Данные записаны.\nНапишите /start").DoVoid(ctx)
}
