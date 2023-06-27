package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"time"
)

func (h *DaimyoHandler) ReportMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoReportMenu.EnterShiftData:
		return nil
	case domain.DaimyoReportMenu.ReportRequest:
		kb := tg.NewReplyKeyboardMarkup(
			tg.NewButtonColumn(
				tg.NewKeyboardButton(domain.DaimyoReportPeriodMenu.ForShift),
				tg.NewKeyboardButton(domain.DaimyoReportPeriodMenu.ForPeriod),
			)...,
		).WithResizeKeyboardMarkup()

		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoReportPeriodMenuHandler

		return msg.Answer("Выберите действие").ReplyMarkup(kb).DoVoid(ctx)
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) ReportPeriodMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoReportPeriodMenu.ForShift:
		reports, err := h.daimyoService.CreateSamuraiReport(
			ctx,
			string(msg.From.Username),
			time.Now().Format("2006-01-02"),
		)
		if err != nil {
			return err
		}

		for _, item := range reports {
			msg.Answer(item).DoVoid(ctx)
		}

		h.sessionManager.Reset(h.sessionManager.Get(ctx))

		return msg.Answer("Данные за смену").DoVoid(ctx)
	case domain.DaimyoReportPeriodMenu.ForPeriod:
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoReportStartPeriod

		return msg.Answer("Введите дату начала периода.\nПример 2023-06-15").
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) ReportStartPeriodHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	_, err := time.Parse("2006-01-02", msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepDaimyoReportStartPeriod
		return msg.Answer("Введите дату начала периода.\nПример 2023-06-15").DoVoid(ctx)
	}

	sessionManager.ReportPeriod.StartDate = msg.Text

	sessionManager.Step = domain.SessionStepDaimyoReportEndPeriod

	return msg.Answer("Введите дату конца периода включительно.\nПример 2023-06-15").DoVoid(ctx)
}

func (h *DaimyoHandler) ReportEndPeriodHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	_, err := time.Parse("2006-01-02", msg.Text)
	if err != nil {
		sessionManager.Step = domain.SessionStepDaimyoReportEndPeriod
		return msg.Answer("Введите дату конца периода включительно.\nПример 2023-06-15").DoVoid(ctx)
	}

	sessionManager.ReportPeriod.EndDate = msg.Text

	reports, err := h.daimyoService.CreateSamuraiReportWithPeriod(
		ctx,
		string(msg.From.Username),
		sessionManager.ReportPeriod.StartDate,
		sessionManager.ReportPeriod.EndDate,
	)
	if err != nil {
		return err
	}

	for _, item := range reports {
		msg.Answer(item).DoVoid(ctx)
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
