package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type CardService interface {
	GetByUsername(ctx context.Context, daimyoUsername string) ([]*domain.CardDTO, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
}

type CashManagerService interface {
	//GetByDaimyoUsername(ctx context.Context, username string) (domain.CashManagerDTO, error)
}

type DaimyoHandler struct {
	sessionManager *session.Manager[domain.Session]

	cardService                 CardService
	replenishmentRequestService ReplenishmentRequestService
	cashManagerService          CashManagerService

	//replenishmentRequest domain.ReplenishmentRequestDTO
}

func NewDaimyoHandler(
	sm *session.Manager[domain.Session],
	cardService CardService,
	replenishmentRequestService ReplenishmentRequestService,
	cashManagerService CashManagerService,
) *DaimyoHandler {
	return &DaimyoHandler{
		sessionManager:              sm,
		cardService:                 cardService,
		replenishmentRequestService: replenishmentRequestService,
		cashManagerService:          cashManagerService,
		//replenishmentRequest:        domain.ReplenishmentRequestDTO{},
	}
}

func (h *DaimyoHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoMenu.MakeReplenishmentRequest:
		cards, err := h.cardService.GetByUsername(ctx, string(msg.Chat.Username))
		if err != nil {
			return err
		}

		var str string

		for i, card := range cards {
			str += fmt.Sprintf("%d. %s\n", i+1, card.Name)
		}

		h.sessionManager.Get(ctx).Step = domain.SessionStepEnterReplenishmentRequestAmount

		return msg.Answer(fmt.Sprintf("Введите название карты из списка, которую хотите пополнить: \n%s", str)).DoVoid(ctx)
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

//
//func (h *DaimyoHandler) CreateEntityMenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
//	switch msg.Text {
//	case domain.AdminCreateEnityMenu.CreateShogun:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEnityMenu.CreateDaimyo:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateDaimyoUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEnityMenu.CreateSamurai:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEnityMenu.CreateCashManager:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEnityMenu.CreateCard:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardName
//
//		return msg.Answer(fmt.Sprintf("Введите название карты")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//	//case domain.AdminCreateEnityMenu.Back:
//	//	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
//	//	return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
//	default:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
//		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
//	}
//}
