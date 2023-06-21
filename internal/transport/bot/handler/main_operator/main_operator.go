package main_operator

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type MainOperatorService interface {
	GetByUsername(ctx context.Context, username string) (domain.MainOperatorDTO, error)
}

type CardService interface {
	GetAllByUsername(ctx context.Context, bankName string, ownerUsername string) ([]*domain.CardDTO, error)
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
	GetByUsername(ctx context.Context, ownerUsername string) (domain.CardDTO, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
	CheckIfExists(ctx context.Context, cardName string) (bool, error)
}

type MainOperatorHandler struct {
	sessionManager *session.Manager[domain.Session]

	cardService                 CardService
	replenishmentRequestService ReplenishmentRequestService
	mainOperatorService         MainOperatorService
}

func New(
	sm *session.Manager[domain.Session],
	cardService CardService,
	replenishmentRequestService ReplenishmentRequestService,
	mainOperatorService MainOperatorService,
) *MainOperatorHandler {
	return &MainOperatorHandler{
		sessionManager:              sm,
		cardService:                 cardService,
		replenishmentRequestService: replenishmentRequestService,
		mainOperatorService:         mainOperatorService,
	}
}

func (h *MainOperatorHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.MainOperatorMainMenu.MakeReplenishmentRequest:
		banks, err := h.cardService.GetBankNames(ctx)
		if err != nil {
			return err
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, bank := range banks {
			buttons = append(buttons, tg.NewKeyboardButton(bank.Name))
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepMainOperatorChooseReplenishmentRequestBank

		return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)

	case domain.MainOperatorMainMenu.Requests:
		return nil

	case domain.MainOperatorMainMenu.FillReport:
		return nil

	case domain.MainOperatorMainMenu.WithdrawalRequest:
		return nil

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
