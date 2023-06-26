package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
	"strings"
)

func (h *CashManagerHandler) RepReqMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.CashManagerRepRequestsMenu.Active:
		requests, err := h.replenishmentRequestService.GetAllByCashManager(
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
					int(request.RequiredAmount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerActRepReqSelectHandler

		return msg.Answer("Выберите запрос").ReplyMarkup(kb).DoVoid(ctx)

	case domain.CashManagerRepRequestsMenu.Objectionable:
		requests, err := h.replenishmentRequestService.GetAllByCashManager(
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
					int(request.RequiredAmount),
				)),
			)
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerObjRepReqSelectHandler

		return msg.Answer("Выберите запрос").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *CashManagerHandler) ActRepReqSelectHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	s := strings.Split(msg.Text, "/")
	cardMsg := strings.Split(s[0], " ")

	request, err := h.replenishmentRequestService.GetByCardName(ctx, cardMsg[0])
	if err != nil {
		return err
	}

	sessionManager.ReplenishmentRequest = request

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.CashManagerConfirmRepRequestsMenu.Confirm),
			tg.NewKeyboardButton(domain.CashManagerConfirmRepRequestsMenu.Cancel),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCashManagerActRepReqActionHandler

	return msg.Answer("Подтвердите действие").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *CashManagerHandler) ActRepReqActionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.CashManagerConfirmRepRequestsMenu.Confirm:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerActRepReqConfirmActionHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton("Да"),
				tg.NewKeyboardButton("Нет"),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)
	case domain.CashManagerConfirmRepRequestsMenu.Cancel:
		return nil
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *CashManagerHandler) ActRepReqConfirmActionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case "Да":
		sessionManager := h.sessionManager.Get(ctx)

		//if err := h.replenishmentRequestService.ChangeStatus(
		//	ctx,
		//	sessionManager.ReplenishmentRequest.CardName,
		//	domain.ObjectionableRequests.String(),
		//); err != nil {
		//	return err
		//}

		//sessionManager.ReplenishmentRequest.ActualAmount = sessionManager.ReplenishmentRequest.RequiredAmount

		if err := h.replenishmentRequestService.ConfirmRequest(ctx, sessionManager.ReplenishmentRequest); err != nil {
			return err
		}

		h.sessionManager.Reset(sessionManager)

		return msg.Answer("Данные записаны.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	case "Нет":
		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerRepReqAnotherAmountHandler

		return msg.Answer("Введите сумму").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *CashManagerHandler) ObjRepReqSelectHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	s := strings.Split(msg.Text, "/")
	cardMsg := strings.Split(s[0], " ")

	request, err := h.replenishmentRequestService.GetByCardName(ctx, cardMsg[0])
	if err != nil {
		return err
	}

	sessionManager.ReplenishmentRequest = request

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton("Пополнено на другую сумму"),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepCashManagerObjRepReqAnotherAmountSelectHandler

	return msg.Answer("Подтвердите действие").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *CashManagerHandler) ObjRepReqAnotherAmountSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerRepReqAnotherAmountHandler

	return msg.Answer("Введите сумму").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}

func (h *CashManagerHandler) RepReqAnotherAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.Atoi(msg.Text)
	if err != nil {
		return err
	}

	//sessionManager.ReplenishmentRequest.ActualAmount += float32(amount)
	sessionManager.ReplenishmentRequest.RequiredAmount = float32(amount)
	//sessionManager.ReplenishmentRequest.ActualAmount = float32(amount)

	if err := h.replenishmentRequestService.ConfirmRequest(ctx, sessionManager.ReplenishmentRequest); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Данные записаны.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
