package daimyo

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/utils"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *DaimyoHandler) HierarchyMenuHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.DaimyoHierarchyMenu.CreateSamurai:
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoCreateSamuraiNickname

		return msg.Answer(fmt.Sprintf("Введите имя")).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	case domain.DaimyoHierarchyMenu.InSubordination:
		samurais, err := h.samuraiService.GetAllByDaimyo(ctx, string(msg.From.Username))
		if err != nil {
			return err
		}

		buttons := make([]tg.KeyboardButton, 0)

		for _, samurai := range samurais {
			buttons = append(buttons, tg.NewKeyboardButton(samurai.Nickname))
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

func (h *DaimyoHandler) EnterSamuraiNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Nickname = msg.Text

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Tag),
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Link),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepDaimyoSamuraiCreationMethod
	return msg.Answer("Введите telegram username").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *DaimyoHandler) SamuraiCreationHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.EntityCreationMethodMenu.Tag:
		h.sessionManager.Get(ctx).Step = domain.SessionStepDaimyoCreateSamurai
		return msg.Answer("Введите telegram username").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	case domain.EntityCreationMethodMenu.Link:
		sessionManager := h.sessionManager.Get(ctx)

		link := utils.GenerateLink()

		sessionManager.Samurai.Username = link
		sessionManager.Samurai.DaimyoUsername = string(msg.From.Username)

		if err := h.referalService.Create(ctx, link, domain.SamuraiRole.String()); err != nil {
			return err
		}

		if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
			return err
		}

		h.sessionManager.Reset(sessionManager)

		me, err := msg.Client.Me(ctx)
		if err != nil {
			return err
		}

		return msg.Answer(fmt.Sprintf("https://t.me/%s?start=%s", me.Username, link)).
			ReplyMarkup(tg.NewReplyKeyboardRemove()).
			DoVoid(ctx)

	default:
		h.sessionManager.Reset(h.sessionManager.Get(ctx))
		return msg.Answer("Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
	}
}

func (h *DaimyoHandler) CreateSamuraiHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Samurai.Username = strings.ReplaceAll(msg.Text, "@", "")
	sessionManager.Samurai.DaimyoUsername = string(msg.From.Username)

	if err := h.samuraiService.Create(ctx, sessionManager.Samurai); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Самурай успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
