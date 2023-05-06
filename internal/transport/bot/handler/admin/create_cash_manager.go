package admin

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"strings"
)

func (h *AdminHandler) EnterCashManagerUsername(ctx context.Context, msg *tgb.MessageUpdate) error {
	// TODO: create regular expression to check the username is correct
	h.cashManager.Username = strings.ReplaceAll(msg.Text, "@", "")

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManagerNickname
	return msg.Answer("Введите nickname").DoVoid(ctx)
}

func (h *AdminHandler) EnterCashManagerNickname(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.cashManager.Nickname = msg.Text

	shoguns, err := h.shogunService.GetAll(ctx)
	if err != nil {
		return err
	}

	var str string

	for _, shogun := range shoguns {
		str += "@" + shogun.Username + "\n"
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepCreateCashManager

	return msg.Answer(fmt.Sprintf("Введите username сёгуна, к которому будет привязан инкассатор. \nСписок сёгунов: \n%s", str)).
		ReplyMarkup(tg.NewReplyKeyboardRemove()).
		DoVoid(ctx)
}

func (h *AdminHandler) CreateCashManager(ctx context.Context, msg *tgb.MessageUpdate) error {
	h.cashManager.ShogunUsername = strings.ReplaceAll(msg.Text, "@", "")

	if err := h.cashManagerService.Create(ctx, h.cashManager); err != nil {
		return err
	}

	h.sessionManager.Get(ctx).Step = domain.SessionStepInit
	return msg.Answer("Инкассатор успешно создан. Напишите /start").ReplyMarkup(tg.NewReplyKeyboardRemove()).DoVoid(ctx)
}
