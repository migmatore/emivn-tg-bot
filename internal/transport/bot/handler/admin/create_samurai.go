package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterSamuraiUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	// TODO: create regular expression to check the username is correct
	h.samurai.Username = strings.ReplaceAll(msg.Text, "@", "")

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamuraiNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterSamuraiNickname(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.samurai.Nickname = msg.Text

	daimyos, err := h.daimyoService.GetAll(ctx)
	if err != nil {
		return err
	}

	var str string

	for _, daimyo := range daimyos {
		str += "@" + daimyo.Username + "\n"
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateSamurai

	return msg.Answer(fmt.Sprintf("Введите username даймё, к которому будет привязан самурай. \nСписок даёме: \n%s", str)).
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateSamurai(ctx context.Context, msg *tgb.MessageUpdate) error {
	// TODO: create regular expression to check the username is correct
	h.samurai.DaimyoUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.samuraiService.Create(ctx, h.samurai); err != nil {
		return err
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	return msg.Answer("Самурай успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
