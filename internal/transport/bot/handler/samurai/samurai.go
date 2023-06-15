package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"log"
	"strconv"
	"time"
)

type CardService interface {
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
}

type SamuraiService interface {
	GetByUsername(ctx context.Context, username string) (domain.SamuraiDTO, error)
	CreateTurnover(ctx context.Context, dto domain.SamuraiTurnoverDTO) error
}

type SamuraiHandler struct {
	sessionManager *session.Manager[domain.Session]

	cardService    CardService
	samuraiService SamuraiService
}

func NewSamuraiHandler(
	sm *session.Manager[domain.Session],
	cardService CardService,
	samuraiService SamuraiService,
) *SamuraiHandler {
	return &SamuraiHandler{
		sessionManager: sm,
		cardService:    cardService,
		samuraiService: samuraiService,
	}
}

func (h *SamuraiHandler) EnterDataMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	if msg.Text != domain.SamuraiMainMenu.EnterData {
		sessionManager.Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}

	sessionManager.SamuraiTurnover.SamuraiUsername = string(msg.From.Username)

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

	sessionManager.Step = domain.SessionStepSamuraiChooseBankMenuHandler

	return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *SamuraiHandler) ChooseBankMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.SamuraiTurnover.BankTypeName = msg.Text

	sessionManager.Step = domain.SessionStepSamuraiCreateTurnoverHandler

	return msg.Answer("Введите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").DoVoid(ctx)
}

func (h *SamuraiHandler) CreateTurnoverMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	finalAmount, err := strconv.ParseFloat(msg.Text, 64)
	if err != nil {
		sessionManager.Step = domain.SessionStepSamuraiChooseBankMenuHandler

		return msg.Answer("Введите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").
			DoVoid(ctx)
	}

	sessionManager.SamuraiTurnover.FinalAmount = finalAmount

	if err := h.samuraiService.CreateTurnover(ctx, sessionManager.SamuraiTurnover); err != nil {
		return err
	}

	sessionManager.Step = domain.SessionStepInit

	return msg.Answer("Данные записаны. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}

func (h *SamuraiHandler) Notify(ctx context.Context, args domain.FuncArgs) (status domain.TaskStatus, when interface{}) {
	if id, ok := args["id"]; ok {

		var id tg.ChatID = tg.ChatID(id.(float64))

		if client, ok := ctx.Value(domain.TaskKey{}).(*tg.Client); ok {
			client.SendMessage(id, "hello").DoVoid(ctx)
		}

		return domain.TaskStatusWait, time.Now().Add(time.Second * 10)
	}
	//else {
	//	if client, ok := ctx.Value(domain.TaskKey{}).(*tg.Client); ok {
	//		var id tg.ChatID = 6109520093
	//
	//		client.SendMessage(id, "hello").DoVoid(ctx)
	//	}
	//}

	log.Print("Not found name arg in func args")

	return domain.TaskStatusWait, time.Now().Add(time.Second * 30)
}
