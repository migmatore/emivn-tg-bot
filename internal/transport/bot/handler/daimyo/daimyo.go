package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type CardService interface {
	GetAllByUsername(ctx context.Context, bankName string, daimyoUsername string) ([]*domain.CardDTO, error)
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
	GetByUsername(ctx context.Context, daimyoUsername string) (domain.CardDTO, error)
	GetByName(ctx context.Context, name string) (domain.CardDTO, error)
	ChangeLimit(ctx context.Context, name string, limit int) error
	GetLimits(ctx context.Context, owner string) ([]string, error)
}

type DaimyoService interface {
	CreateSamuraiReport(ctx context.Context, date string) ([]string, error)
	GetByUsername(ctx context.Context, username string) (domain.DaimyoDTO, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
	CheckIfExists(ctx context.Context, cardName string) (bool, error)
	GetAllByOwner(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequestDTO, error)
	GetByCardName(ctx context.Context, name string) (domain.ReplenishmentRequestDTO, error)
	ConfirmRequest(ctx context.Context, dto domain.ReplenishmentRequestDTO) error
}

type CashManagerService interface {
	//GetByDaimyoUsername(ctx context.Context, username string) (domain.CashManagerDTO, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
	GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.SamuraiDTO, error)
}

type ReferalService interface {
	Create(ctx context.Context, link string, role string) error
}

type Scheduler interface {
	Add(ctx context.Context, dto domain.TaskDTO) error
	Delete(ctx context.Context, taskName string) error
}

type DaimyoHandler struct {
	sessionManager *session.Manager[domain.Session]

	cardService                 CardService
	daimyoService               DaimyoService
	replenishmentRequestService ReplenishmentRequestService
	cashManagerService          CashManagerService
	samuraiService              SamuraiService
	referalService              ReferalService

	scheduler Scheduler
}

func NewDaimyoHandler(
	sm *session.Manager[domain.Session],
	cardService CardService,
	daimyoService DaimyoService,
	replenishmentRequestService ReplenishmentRequestService,
	cashManagerService CashManagerService,
	samuraiService SamuraiService,
	referalService ReferalService,
	scheduler Scheduler,
) *DaimyoHandler {
	return &DaimyoHandler{
		sessionManager:              sm,
		cardService:                 cardService,
		daimyoService:               daimyoService,
		replenishmentRequestService: replenishmentRequestService,
		cashManagerService:          cashManagerService,
		samuraiService:              samuraiService,
		referalService:              referalService,
		scheduler:                   scheduler,
	}
}

func (h *DaimyoHandler) MainMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	//if err := h.schedulerService.Add(ctx, domain.TaskDTO{
	//	Alias:           "notify_samurai",
	//	Name:            fmt.Sprintf("notify_samurai %s", msg.From.Username),
	//	Arguments:       domain.FuncArgs{"chatId": msg.Chat.ID},
	//	IntervalMinutes: 0,
	//	RunAt:           time.Now().Add(time.Second * 10),
	//}); err != nil {
	//	return err
	//}

	switch msg.Text {
	case domain.DaimyoMainMenu.MakeReplenishmentRequest:
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

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoChooseReplenishmentRequestBank

		return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)

	case domain.DaimyoMainMenu.Requests:
		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.DaimyoRepRequestsMenu.Active),
				tg.NewKeyboardButton(domain.DaimyoRepRequestsMenu.Objectionable),
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoRepReqMenuHandler

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	case domain.DaimyoMainMenu.CardLimit:
		limits, err := h.cardService.GetLimits(ctx, string(msg.From.Username))
		if err != nil {
			return err
		}

		for _, limit := range limits {
			return msg.Answer(limit).DoVoid(ctx)
		}

		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	case domain.DaimyoMainMenu.Report:
		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.DaimyoReportMenu.EnterShiftData),
				tg.NewKeyboardButton(domain.DaimyoReportMenu.ReportRequest),
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoReportMenuHandler

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	case domain.DaimyoMainMenu.Hierarchy:
		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.DaimyoHierarchyMenu.CreateSamurai),
				tg.NewKeyboardButton(domain.DaimyoHierarchyMenu.InSubordination),
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoHierarchyMenuHandler

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

//
//func (h *DaimyoHandler) CreateEntityMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
//	switch msg.Text {
//	case domain.AdminCreateEntityMenu.CreateShogun:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunNickname
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEntityMenu.CreateDaimyoHandler:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateDaimyoUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEntityMenu.CreateSamurai:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEntityMenu.CreateCashManagerHandler:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEntityMenu.CreateCard:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCardName
//
//		return msg.Answer(fmt.Sprintf("Введите название карты")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//	//case domain.AdminCreateEntityMenu.Back:
//	//	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
//	//	return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
//	default:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
//		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
//	}
//}
