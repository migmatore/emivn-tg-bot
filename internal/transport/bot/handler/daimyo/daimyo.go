package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type DaimyoHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func NewDaimyoHandler(
	sm *session.Manager[domain.Session],
) *DaimyoHandler {
	return &DaimyoHandler{
		sessionManager: sm,
	}
}

func (h *DaimyoHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoMenu.MakeReplenishmentRequest:
		h.sessionManager.Get(ctx).Step = domain.SessionStepMakeReplenishmentRequest

		//return msg.Answer(fmt.Sprintf("Введите telegram username")).ReplyMarkup(kb).DoVoid(ctx)
		return msg.Answer("").DoVoid(ctx)
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) CreateEntityMenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminCreateEnityMenu.CreateShogun:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEnityMenu.CreateDaimyo:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateDaimyoUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEnityMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEnityMenu.CreateCashManager:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEnityMenu.CreateCard:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardName

		return msg.Answer(fmt.Sprintf("Введите название карты")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)
	//case domain.AdminCreateEnityMenu.Back:
	//	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	//	return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
