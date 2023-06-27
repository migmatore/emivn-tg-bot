package shogun

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/pkg/utils"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *ShogunHandler) EnterDaimyoNicknameHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)

	sessionManager.Daimyo.Nickname = msg.Text

	kb := tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Tag),
			tg.NewKeyboardButton(domain.EntityCreationMethodMenu.Link),
		)...,
	).WithResizeKeyboardMarkup()

	sessionManager.Step = domain.SessionStepShogunDaimyoCreationMethod

	return msg.Answer("Выберите способ").ReplyMarkup(kb).DoVoid(ctx)
}

func (h *ShogunHandler) DaimyoCreationHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	switch msg.Text {
	case domain.EntityCreationMethodMenu.Tag:
		h.sessionManager.Get(ctx).Step = domain.SessionStepShogunCreateDaimyo
		return msg.Answer("Введите telegram username").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)

	case domain.EntityCreationMethodMenu.Link:
		sessionManager := h.sessionManager.Get(ctx)

		link := utils.GenerateLink(sessionManager.Daimyo.Nickname)

		sessionManager.Daimyo.Username = link
		sessionManager.Daimyo.ShogunUsername = string(msg.From.Username)

		if err := h.referalService.Create(ctx, link, domain.DaimyoRole.String()); err != nil {
			return err
		}

		if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
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

func (h *ShogunHandler) CreateDaimyoHandler(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Username = strings.ReplaceAll(msg.Text, "@", "")
	sessionManager.Daimyo.ShogunUsername = string(msg.From.Username)

	if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)

	return msg.Answer("Даймё успешно создан.\nНапишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
