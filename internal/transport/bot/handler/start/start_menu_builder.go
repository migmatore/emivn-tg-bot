package start

import (
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
)

func buildAdminStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.AdminMainMenu.RequestsList),
			tg.NewKeyboardButton(domain.AdminMainMenu.Cards),
			tg.NewKeyboardButton(domain.AdminMainMenu.Report),
			tg.NewKeyboardButton(domain.AdminMainMenu.Hierarchy),
			tg.NewKeyboardButton(domain.AdminMainMenu.UploadData),
			tg.NewKeyboardButton(domain.AdminMainMenu.Deposits),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildDaimyoStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.DaimyoMainMenu.MakeReplenishmentRequest),
		)...,
	).WithResizeKeyboardMarkup()
}
