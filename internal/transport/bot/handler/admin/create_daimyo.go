package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterDaimyoUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.Username = strings.ReplaceAll(msg.Text, "@", "")

	sessionManager.Step = domain.SessionStepCreateDaimyoNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterDaimyoNickname(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	sessionManager.Daimyo.Nickname = msg.Text

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	var str string

	for _, shogun := range shoguns {
		str += "@" + shogun.Username + "\n"
	}

	sessionManager.Step = domain.SessionStepCreateDaimyo

	return msg.Answer(fmt.Sprintf("Введите username сёгуна, к которому будет привязан даймё. \nСписок сёгунов: \n%s", str)).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateDaimyo(ctx context.Context, msg *tgb.MessageUpdate) error {
	sessionManager := h.sessionManager.Get(ctx)
	// TODO: create regular expression to check the username is correct
	sessionManager.Daimyo.ShogunUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.daimyoService.Create(ctx, sessionManager.Daimyo); err != nil {
		return err
	}

	h.sessionManager.Reset(sessionManager)
	return msg.Answer("Даймё успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
