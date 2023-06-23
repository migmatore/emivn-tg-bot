package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/logging"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strconv"
	"strings"
	"time"
)

func (h *DaimyoHandler) RepReqChooseCardHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
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

	h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoEnterReplenishmentRequestAmount

	return msg.Answer("Введите название карты из списка, которую хотите пополнить:").
		ReplyMarkup(kb).
		DoVoid(ctx)
}

func (h *DaimyoHandler) EnterRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
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

	sessionManager.Step = domain.SessionStepDaimyoMakeReplenishmentRequest
	return msg.Answer("Введите сумму на пополнение").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}

func (h *DaimyoHandler) MakeRepReqHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	amount, err := strconv.ParseFloat(msg.Text, 32)
	if err != nil {
		sessionManager.Step = domain.SessionStepDaimyoMakeReplenishmentRequest
		return msg.Answer("Пожалуйста, введите суточный лимит карты").DoVoid(ctx)
	}

	sessionManager.ReplenishmentRequest.Amount = float32(amount)
	sessionManager.ReplenishmentRequest.OwnerUsername = string(msg.From.Username)

	card, err := h.cardService.GetByUsername(ctx, string(msg.From.Username))
	if err != nil {
		return err
	}

	if float32(amount) > float32(card.DailyLimit) {
		sessionManager.Step = domain.SessionStepDaimyoChangeReplenishmentRequestAmount

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

	taskName := fmt.Sprintf(
		"%s_%d",
		sessionManager.ReplenishmentRequest.CardName,
		int(sessionManager.ReplenishmentRequest.Amount),
	)

	if err := h.scheduler.Add(ctx, domain.TaskDTO{
		Alias:           "change_card_limit",
		Name:            taskName,
		Arguments:       domain.FuncArgs{"task_name": taskName},
		IntervalMinutes: 0,
		RunAt:           time.Now().Add(time.Second * 30),
	}); err != nil {
		return msg.Answer(err.Error()).DoVoid(ctx)
	}

	daimyo, err := h.daimyoService.GetByUsername(ctx, string(msg.From.Username))
	if err != nil {
		return err
	}

	msg.Client.SendMessage(chatId,
		fmt.Sprintf("Появилась новая заявка на пополнение: %s / %s %d", daimyo.Username, card.Name, card.LastDigits)).
		DoVoid(ctx)

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Данные записаны.\n Напишите /start").
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *DaimyoHandler) ChangeRepReqAmountHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case "Ввести другую сумму":
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoMakeReplenishmentRequest
		return msg.Answer("Введите сумму на пополнение").DoVoid(ctx)
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) NotifyCardLimitChange(ctx context.Context, args domain.FuncArgs) (status domain.TaskStatus, when interface{}) {
	if taskName, ok := args["task_name"]; ok {
		s := strings.Split(taskName.(string), "_")

		cardName := s[0]

		card, err := h.cardService.GetByName(ctx, cardName)
		if err != nil {
			return domain.TaskStatusWait, time.Now().Add(time.Second * 5)
		}

		amount, _ := strconv.Atoi(s[1])

		newCardLimit := card.DailyLimit + amount

		if err := h.cardService.ChangeLimit(ctx, cardName, newCardLimit); err != nil {
			logging.GetLogger(ctx).Errorf("error changing the card limit")
		}

		if err := h.scheduler.Delete(ctx, taskName.(string)); err != nil {
			logging.GetLogger(ctx).Errorf("error deliting the task")
		}

		return domain.TaskStatusDone, nil
	}

	return domain.TaskStatusWait, time.Now().Add(time.Second * 30)
}
