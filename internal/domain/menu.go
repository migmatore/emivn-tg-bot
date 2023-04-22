package domain

import "github.com/mr-linch/go-tg"

var AdminMenu = struct {
	CreateEntity string
}{
	CreateEntity: "Создать сущность",
}

var AdminCreateEnityMenu = struct {
	CreateShogun string
	Back         string
}{
	CreateShogun: "Создать сёгуна",
	Back:         "Назад",
}

var Menu = tg.NewReplyKeyboardMarkup(
	tg.NewButtonColumn(
		tg.NewKeyboardButton(AdminMenu.CreateEntity),
	)...,
).WithResizeKeyboardMarkup()
