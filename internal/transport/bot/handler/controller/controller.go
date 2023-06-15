package controller

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"strconv"
)

type CardService interface {
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
}

type DaimyoService interface {
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
}

type SamuraiService interface {
	GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.SamuraiDTO, error)
}

type ControllerService interface {
	CreateTurnover(ctx context.Context, dto domain.ControllerTurnoverDTO) error
}

type ControllerHandler struct {
	sessionManager *session.Manager[domain.Session]

	controllerService ControllerService
	cardService       CardService
	daimyoService     DaimyoService
	samuraiService    SamuraiService
}

func New(
	sm *session.Manager[domain.Session],

	controllerService ControllerService,
	cardService CardService,
	daimyoService DaimyoService,
	samuraiService SamuraiService,
) *ControllerHandler {
	return &ControllerHandler{
		sessionManager:    sm,
		controllerService: controllerService,
		cardService:       cardService,
		daimyoService:     daimyoService,
		samuraiService:    samuraiService,
	}
}

func (h *ControllerHandler) EnterDataMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	if msg.Text != domain.SamuraiMainMenu.EnterData {
		sessionManager.Step = domain.SessionStepInit
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}

	sessionManager.ControllerTurnover.ControllerUsername = string(msg.From.Username)

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range daimyos {
		buttons = append(buttons, tg.NewKeyboardButton(item.Username))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepControllerChooseDaimyoMenuHandler

	return msg.Answer("Выберите дайме").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ControllerHandler) ChooseDaimyoMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	samurais, err := h.samuraiService.GetAllByDaimyo(ctx, msg.Text)
	if err != nil {
		return err
	}

	buttons := make([]tg.KeyboardButton, 0)

	for _, item := range samurais {
		buttons = append(buttons, tg.NewKeyboardButton(item.Username))
	}

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			buttons...,
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepControllerChooseSamuraiMenuHandler

	return msg.Answer("Выберите самурая").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ControllerHandler) ChooseSamuraiMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.ControllerTurnover.SamuraiUsername = msg.Text

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

	sessionManager.Step = domain.SessionStepControllerChooseBankMenuHandler

	return msg.Answer("Выберите банк").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ControllerHandler) ChooseBankMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.ControllerTurnover.BankTypeName = msg.Text

	sessionManager.Step = domain.SessionStepControllerCreateTurnoverHandler

	return msg.Answer("ведите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").DoVoid(ctx)
}

func (h *ControllerHandler) CreateTurnoverMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	finalAmount, err := strconv.ParseFloat(msg.Text, 64)
	if err != nil {
		sessionManager.Step = domain.SessionStepControllerChooseBankMenuHandler

		return msg.Answer("Введите данные на конец смены с 8 до 12 часов дня. Без пробелов, точек и иных знаков.").
			DoVoid(ctx)
	}

	sessionManager.ControllerTurnover.FinalAmount = finalAmount

	if err := h.controllerService.CreateTurnover(ctx, sessionManager.ControllerTurnover); err != nil {
		return err
	}

	sessionManager.Step = domain.SessionStepInit

	return msg.Answer("Данные записаны. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
