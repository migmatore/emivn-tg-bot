package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/utils"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterShogunNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Shogun.Nickname = msg.Text

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Tag),
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Link),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepShogunCreationMethod
	return msg.Answer("Выберите способ").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *AdminHandler) ShogunCreationHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.EntityCreationMethodMenu.Tag:
		h.sessionManager.Get(ctx).Step = domain.SessionStepCreateShogun
		return msg.Answer("Введите telegram username").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	case domain.EntityCreationMethodMenu.Link:
		sessionManager := h.sessionManager.Get(ctx)

		link := utils.GenerateLink(sessionManager.Shogun.Nickname)

		sessionManager.Shogun.Username = link

		if err := h.referalService.Create(ctx, link, domain.ShogunRole.String()); err != nil {
			return err
		}

		if err := h.shogunService.Create(ctx, sessionManager.Shogun); err != nil {
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

func (h *AdminHandler) EnterShogunUsernameAndCreateHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Shogun.Username = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.shogunService.Create(ctx, sessionManager.Shogun); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Сёгун успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
