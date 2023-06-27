package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.DaimyoDTO, error)
	GetByNickname(ctx context.Context, nickname string) (domain.DaimyoDTO, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
}

type MainOperatorService interface {
	Create(ctx context.Context, dto domain.MainOperatorDTO) error
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.CardDTO, error)
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
	GetCardsBalancesByShogun(ctx context.Context, shogunUsername string) ([]string, error)
}

type ReferalService interface {
	Create(ctx context.Context, link string, role string) error
}

type ShogunHandler struct {
	sessionManager *session.Manager[domain.Session]

	daimyoService       DaimyoService
	samuraiService      SamuraiService
	cashManagerService  CashManagerService
	mainOperatorService MainOperatorService
	cardService         CardService
	referalService      ReferalService
}

func NewShogunHandler(
	sm *session.Manager[domain.Session],
	daimyoService DaimyoService,
	samuraiService SamuraiService,
	cashManagerService CashManagerService,
	mainOperatorService MainOperatorService,
	cardService CardService,
	referalService ReferalService,
) *ShogunHandler {
	return &ShogunHandler{
		sessionManager:      sm,
		daimyoService:       daimyoService,
		samuraiService:      samuraiService,
		cashManagerService:  cashManagerService,
		mainOperatorService: mainOperatorService,
		cardService:         cardService,
		referalService:      referalService,
	}
}

func (h *ShogunHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.ShogunMainMenu.Requests:
		return nil

	case domain.ShogunMainMenu.Cards:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCardsMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.ShogunCardsMenu.CreateCard),
				tg.NewKeyboardButton(domain.ShogunCardsMenu.CardsList),
				tg.NewKeyboardButton(domain.ShogunCardsMenu.Limit),
				tg.NewKeyboardButton(domain.ShogunCardsMenu.Balance),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

		return nil

	case domain.ShogunMainMenu.Report:
		return nil

	case domain.ShogunMainMenu.Hierarchy:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunHierarchyMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.ShogunHierarchyMenu.CreateEntity),
				tg.NewKeyboardButton(domain.ShogunHierarchyMenu.InSubordination),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	case domain.ShogunMainMenu.GARA:
		return nil

	case domain.ShogunMainMenu.UploadData:
		return nil

	case domain.ShogunMainMenu.Deposits:
		return nil

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *ShogunHandler) CardsMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
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

		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunChooseCardBankHandler

		return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)

	case domain.ShogunCardsMenu.CardsList:
		return nil
	case domain.ShogunCardsMenu.Limit:
		return nil

	case domain.ShogunCardsMenu.Balance:
		sessionManager := h.sessionManager.Get(ctx)

		cardsBalances, err := h.cardService.GetCardsBalancesByShogun(ctx, string(msg.From.Username))
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

func (h *ShogunHandler) HierarchyMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.ShogunHierarchyMenu.CreateEntity:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateEntityMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.ShogunCreateEntityMenu.CreateDaimyo),
				tg.NewKeyboardButton(domain.ShogunCreateEntityMenu.CreateSamurai),
				tg.NewKeyboardButton(domain.ShogunCreateEntityMenu.CreateCashManager),
				tg.NewKeyboardButton(domain.ShogunCreateEntityMenu.CreateMainOperator),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите сущность, которую хотите создать").ReplyMarkup(kb).DoVoid(ctx)

	case domain.ShogunHierarchyMenu.InSubordination:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunSubordinationMenuHandler

		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.ShogunSubordinationMenu.Daimyo),
				tg.NewKeyboardButton(domain.ShogunSubordinationMenu.Samurai),
				tg.NewKeyboardButton(domain.ShogunSubordinationMenu.MainOperator),
				tg.NewKeyboardButton(domain.ShogunSubordinationMenu.CashManager),
			)...,
		).WithResizeKeyboardMarkup()

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *ShogunHandler) CreateEntityMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.ShogunCreateEntityMenu.CreateDaimyo:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateDaimyoNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.ShogunCreateEntityMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateSamuraiNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.ShogunCreateEntityMenu.CreateCashManager:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateCashManagerNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.ShogunCreateEntityMenu.CreateMainOperator:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateMainOperatorNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
