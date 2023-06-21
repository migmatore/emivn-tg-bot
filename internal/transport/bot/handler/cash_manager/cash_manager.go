package cash_manager

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type ReplenishmentRequestService interface {
	GetAllByCashManager(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequestDTO, error)
}

type CardService interface {
	GetByName(ctx context.Context, name string) (domain.CardDTO, error)
}

type DaimyoService interface {
	GetByUsername(ctx context.Context, username string) (domain.DaimyoDTO, error)
}

type MainOperatorService interface {
	GetByUsername(ctx context.Context, username string) (domain.MainOperatorDTO, error)
}

type CashManagerHandler struct {
	sessionManager *session.Manager[domain.Session]

	replenishmentRequestService ReplenishmentRequestService
	cardService                 CardService
	daimyoService               DaimyoService
	mainOperatorService         MainOperatorService
}

func New(
	sm *session.Manager[domain.Session],
	replenishmentRequestService ReplenishmentRequestService,
	cardService CardService,
	daimyoService DaimyoService,
	mainOperatorService MainOperatorService,
) *CashManagerHandler {
	return &CashManagerHandler{
		sessionManager:              sm,
		replenishmentRequestService: replenishmentRequestService,
		cardService:                 cardService,
		daimyoService:               daimyoService,
		mainOperatorService:         mainOperatorService,
	}
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
