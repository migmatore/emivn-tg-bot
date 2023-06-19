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

type ShogunHandler struct {
	sessionManager *session.Manager[domain.Session]

	daimyoService       DaimyoService
	samuraiService      SamuraiService
	cashManagerService  CashManagerService
	mainOperatorService MainOperatorService
}

func NewShogunHandler(
	sm *session.Manager[domain.Session],
	daimyoService DaimyoService,
	samuraiService SamuraiService,
	cashManagerService CashManagerService,
	mainOperatorService MainOperatorService,
) *ShogunHandler {
	return &ShogunHandler{
		sessionManager:      sm,
		daimyoService:       daimyoService,
		samuraiService:      samuraiService,
		cashManagerService:  cashManagerService,
		mainOperatorService: mainOperatorService,
	}
}

func (h *ShogunHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.ShogunMainMenu.Requests:
		return nil

	case domain.ShogunMainMenu.Cards:
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
