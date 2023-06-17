package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *DaimyoHandler) HierarchyMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoHierarchyMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoCreateSamuraiUsername

		return msg.Answer(fmt.Sprintf("Введите telegram username")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.DaimyoHierarchyMenu.InSubordination:
		samurais, err := h.samuraiService.GetAllByDaimyo(ctx, string(msg.From.Username))
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

		h.sessionManager.Reset(h.sessionManager.Get(ctx))

		return msg.Answer("Напишите /start").ReplyMarkup(kb).DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) EnterSamuraiUsernameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepDaimyoCreateSamuraiNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *DaimyoHandler) CreateSamuraiHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Nickname = msg.Text
	sessionManager.Samurai.DaimyoUsername = string(msg.From.Username)

	if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Самурай успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
