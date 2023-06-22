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
	GetByNickname(ctx context.Context, nickname string) (domain.ShogunDTO, error)
}

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.DaimyoDTO, error)
	GetByNickname(ctx context.Context, nickname string) (domain.DaimyoDTO, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
}

type ControllerService interface {
	Create(ctx context.Context, dto domain.ControllerDTO) error
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.CardDTO, error)
	GetCardsBalancesByShogun(ctx context.Context, shogunUsername string) ([]string, error)
}

type AdminHandler struct {
	sessionManager *session.Manager[domain.Session]

	shogunService      ShogunService
	daimyoService      DaimyoService
	samuraiService     SamuraiService
	cashManagerService CashManagerService
	controllerService  ControllerService
	cardService        CardService
}

func NewAdminHandler(
	sm *session.Manager[domain.Session],
	shogunService ShogunService,
	daimyoService DaimyoService,
	samuraiService SamuraiService,
	cashManagerService CashManagerService,
	controllerService ControllerService,
	cardService CardService,
) *AdminHandler {
	return &AdminHandler{
		sessionManager:     sm,
		shogunService:      shogunService,
		daimyoService:      daimyoService,
		samuraiService:     samuraiService,
		cashManagerService: cashManagerService,
		controllerService:  controllerService,
		cardService:        cardService,
	}
}

func (h *AdminHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminMainMenu.Cards:
		shoguns, err := h.shogunService.GetAll(ctx)
		if err != nil {
			return err
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, shogun := range shoguns {
			buttons = append(buttons, tg.NewKeyboardButton(shogun.Nickname))
		}

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				buttons...,
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepAdminCardsChooseShogunHandler

		return msg.Answer("Выберите сёгуна").ReplyMarkup(kb).DoVoid(ctx)

	case domain.AdminMainMenu.RequestsList:
		return nil

	case domain.AdminMainMenu.Report:
		return nil

	case domain.AdminMainMenu.Hierarchy:
		h.sessionManager.Get(ctx).Step = domain.SessionStepHierarchyMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.AdminHierarchyMenu.CreateEntity),
				tg.NewKeyboardButton(domain.AdminHierarchyMenu.InSubordination),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	case domain.AdminMainMenu.CreateGARA:
		return nil

	case domain.AdminMainMenu.UploadData:
		return nil

	case domain.AdminMainMenu.Deposits:
		return nil

	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *AdminHandler) CardsChooseShogunHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	shogun, err := h.shogunService.GetByNickname(ctx, msg.Text)
	if err != nil {
		return err
	}

	sessionManager.Shogun.Username = shogun.Username

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.AdminCardsMenu.CreateCard),
			tg.NewKeyboardButton(domain.AdminCardsMenu.CardsList),
			tg.NewKeyboardButton(domain.AdminCardsMenu.Limit),
			tg.NewKeyboardButton(domain.AdminCardsMenu.Balance),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepAdminCardsMenuHandler

	return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *AdminHandler) CardsMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.ShogunCardsMenu.CreateCard:
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

		h.sessionManager.Get(ctx).Step = domain.SessionStepAdminChooseCardBankHandler

		return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)

	case domain.ShogunCardsMenu.CardsList:
		return nil
	case domain.ShogunCardsMenu.Limit:
		return nil
	case domain.ShogunCardsMenu.Balance:
		sessionManager := h.sessionManager.Get(ctx)

		cardsBalances, err := h.cardService.GetCardsBalancesByShogun(ctx, sessionManager.Shogun.Username)
		if err != nil {
			return err
		}

		for _, cardBalance := range cardsBalances {
			msg.Answer(cardBalance).DoVoid(ctx)
		}

		h.sessionManager.Reset(sessionManager)
		return msg.Answer("Балансы карт на данный момент.\nНапишите /start").DoVoid(ctx)
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *AdminHandler) HierarchyMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.AdminHierarchyMenu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateEntityMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateShogun),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateDaimyo),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateSamurai),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateCashManager),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateController),
				tg.NewKeyboardButton(domain.AdminCreateEntityMenu.CreateMainOperator),
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
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateDaimyo:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateDaimyoNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateCashManager:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.AdminCreateEntityMenu.CreateController:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateControllerNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
