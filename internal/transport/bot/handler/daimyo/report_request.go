package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
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
		reports, err := h.daimyoService.CreateSamuraiReport(ctx, "2023-06-16")
		if err != nil {
			return err
		}

		for _, item := range reports {
			msg.Answer(item).DoVoid(ctx)
		}

		h.sessionManager.Reset(h.sessionManager.Get(ctx))

		return msg.Answer("Данные за смену").DoVoid(ctx)
	case domain.DaimyoReportPeriodMenu.ForPeriod:
		return nil
	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}
