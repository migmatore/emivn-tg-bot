package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type ShogunService interface {
	Create(ctx context.Context, dto domain.ShogunDTO) error
	GetAll(ctx context.Context) ([]*domain.ShogunDTO, error)
}

type Menu struct {
	CreateEntity string
}

type AdminHandler struct {
	sessionManager *session.Manager[domain.Session]

	shogunService ShogunService

	shogun domain.ShogunDTO
	daimyo domain.DaimyoDTO
}

func NewDbWriteHandler(sm *session.Manager[domain.Session], s ShogunService) *AdminHandler {
	return &AdminHandler{
		sessionManager: sm,
		shogunService:  s,
		shogun:         domain.ShogunDTO{},
		daimyo:         domain.DaimyoDTO{},
	}
}

func (h *AdminHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminMenu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateEntityHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateShogun),
				tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateDaimyo),
				tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateSamurai),
				tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateCashManager),
				tg.NewKeyboardButton(domain.AdminCreateEnityMenu.CreateCard),
			)...,
		).WithResizeKeyboardMarkup()

		//return msg.Answer(fmt.Sprintf("Введите telegram username")).ReplyMarkup(kb).DoVoid(ctx)
		return msg.Update.Reply(ctx, msg.Answer(fmt.Sprintf("Выберите сущность, которую хотите создать")).
			ReplyMarkup(kb))
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *AdminHandler) CreateEntityMenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
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
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardBankInfo

		return msg.Answer(fmt.Sprintf("Введите информацию о банке эмитенте")).
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
