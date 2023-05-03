package start

import (
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
)

func buildAdminStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.AdminMenu.CreateEntity),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildDaimyoStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.DaimyoMenu.MakeReplenishmentRequest),
		)...,
	).WithResizeKeyboardMarkup()
}
