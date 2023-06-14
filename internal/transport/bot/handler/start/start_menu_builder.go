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
			tg.NewKeyboardButton(domain.AdminMainMenu.CreateGARA),
			tg.NewKeyboardButton(domain.AdminMainMenu.UploadData),
			tg.NewKeyboardButton(domain.AdminMainMenu.Deposits),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildShogunStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.ShogunMainMenu.Requests),
			tg.NewKeyboardButton(domain.ShogunMainMenu.Cards),
			tg.NewKeyboardButton(domain.ShogunMainMenu.Report),
			tg.NewKeyboardButton(domain.ShogunMainMenu.Hierarchy),
			tg.NewKeyboardButton(domain.ShogunMainMenu.GARA),
			tg.NewKeyboardButton(domain.ShogunMainMenu.UploadData),
			tg.NewKeyboardButton(domain.ShogunMainMenu.Deposits),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildCashManagerStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.CashManagerMainMenu.Requests),
			tg.NewKeyboardButton(domain.CashManagerMainMenu.WithdrawalRequests),
			tg.NewKeyboardButton(domain.CashManagerMainMenu.RemainingFunds),
			tg.NewKeyboardButton(domain.CashManagerMainMenu.CurrentBalance),
			tg.NewKeyboardButton(domain.CashManagerMainMenu.ReplenishmentList),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildDaimyoStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.DaimyoMainMenu.MakeReplenishmentRequest),
			tg.NewKeyboardButton(domain.DaimyoMainMenu.Requests),
			tg.NewKeyboardButton(domain.DaimyoMainMenu.CardLimit),
			tg.NewKeyboardButton(domain.DaimyoMainMenu.Report),
			tg.NewKeyboardButton(domain.DaimyoMainMenu.Hierarchy),
		)...,
	).WithResizeKeyboardMarkup()
}

func buildSamuraiStartMenu() *tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboardMarkup(
		tg.NewButtonColumn(
			tg.NewKeyboardButton(domain.SamuraiMainMenu.EnterData),
		)...,
	).WithResizeKeyboardMarkup()
}
