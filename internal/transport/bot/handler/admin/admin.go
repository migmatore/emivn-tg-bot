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

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
}

type AdminHandler struct {
	sessionManager *session.Manager[domain.Session]

	shogunService      ShogunService
	daimyoService      DaimyoService
	samuraiService     SamuraiService
	cashManagerService CashManagerService
	cardService        CardService
}

func NewAdminHandler(
	sm *session.Manager[domain.Session],
	shogunService ShogunService,
	daimyoService DaimyoService,
	samuraiService SamuraiService,
	cashManagerService CashManagerService,
	cardService CardService,
) *AdminHandler {
	return &AdminHandler{
		sessionManager:     sm,
		shogunService:      shogunService,
		daimyoService:      daimyoService,
		samuraiService:     samuraiService,
		cashManagerService: cashManagerService,
		cardService:        cardService,
	}
}

func (h *AdminHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {

	switch msg.Text {
	case domain.AdminMainMenu.Hierarchy:
		h.sessionManager.Get(ctx).Step = domain.SessionStepHierarchyMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.AdminHierarchyMenu.CreateEntity),
				tg.NewKeyboardButton(domain.AdminHierarchyMenu.InSubordination),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *AdminHandler) HierarchyMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminHierarchyMenu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateEntityHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateShogun),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateDaimyo),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateSamurai),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateCashManager),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateCard),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите сущность, которую хотите создать").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *AdminHandler) CreateEntityMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminCreateEntityMenu.CreateShogun:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateDaimyo:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateDaimyoUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateCashManager:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateCard:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardBank

		banks, err := h.cardService.GetBankNames(ctx)
		if err != nil {
			return err
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, item := range banks {
			buttons = append(buttons, tg.NewKeyboardButton(item.Name))
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer(fmt.Sprintf("Выберите банк")).
			ReplyMarkup(kb).
			DoVoid(ctx)
	//case domain.AdminCreateEntityMenu.Back:
	//	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	//	return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
