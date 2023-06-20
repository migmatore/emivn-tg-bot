package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type CashManagerHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func New(sm *session.Manager[domain.Session]) *CashManagerHandler {
	return &CashManagerHandler{sessionManager: sm}
}

func (h *CashManagerHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.CashManagerMainMenu.Requests:
		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.CashManagerRepRequestsMenu.Active),
				tg.NewKeyboardButton(domain.CashManagerRepRequestsMenu.Objectionable),
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepCashManagerRepReqMenuHandler

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	case domain.CashManagerMainMenu.WithdrawalRequests:
		return nil

	case domain.CashManagerMainMenu.RemainingFunds:
		return nil

	case domain.CashManagerMainMenu.CurrentBalance:
		return nil

	case domain.CashManagerMainMenu.ReplenishmentList:
		return nil

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
