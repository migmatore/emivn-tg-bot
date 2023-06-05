package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"log"
	"time"
)

type CardService interface {
	GetByUsername(ctx context.Context, bankName string, daimyoUsername string) ([]*domain.CardDTO, error)
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
}

type CashManagerService interface {
	//GetByDaimyoUsername(ctx context.Context, username string) (domain.CashManagerDTO, error)
}

type SchedulerService interface {
	Add(ctx context.Context, dto domain.TaskDTO) error
}

type DaimyoHandler struct {
	sessionManager *session.Manager[domain.Session]

	cardService                 CardService
	replenishmentRequestService ReplenishmentRequestService
	cashManagerService          CashManagerService

	schedulerService SchedulerService
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
	}
}

func (h *DaimyoHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
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
	case domain.DaimyoMenu.MakeReplenishmentRequest:
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

		h.sessionManager.Get(ctx).Step = domain.SessionStepEnterReplenishmentRequestCardName

		return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)
	default:
		h.sessionManager.Get(ctx).Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) Notify(args domain.FuncArgs) (status domain.TaskStatus, when interface{}) {
	if name, ok := args["name"]; ok {
		log.Println("PrintWithArgs:", time.Now(), name)
		return domain.TaskStatusDeferred, time.Now().Add(time.Second * 10)
	}

	log.Print("Not found name arg in func args")

	return domain.TaskStatusDeferred, time.Now().Add(time.Second * 10)
}

//
//func (h *DaimyoHandler) CreateEntityMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
//	switch msg.Text {
//	case domain.AdminCreateEntityMenu.CreateShogun:
//		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogunUsername
//
//		return msg.Answer(fmt.Sprintf("Введите telegram username")).
//			ReplyMarkup(tg.NewReplyKeyboardRemove()).
//			DoVoid(ctx)
//
//	case domain.AdminCreateEntityMenu.CreateDaimyo:
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
//	case domain.AdminCreateEntityMenu.CreateCashManager:
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
