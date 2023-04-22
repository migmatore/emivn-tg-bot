package domain

// TODO: menu code refactoring(with back button)

var AdminMenu = struct {
	CreateEntity string
}{
	CreateEntity: "Создать сущность",
}

var AdminCreateEnityMenu = struct {
	CreateShogun      string
	CreateDaimyo      string
	CreateSamurai     string
	CreateCashManager string
	CreateCard        string
	Back              string
}{
	CreateShogun:      "Создать сёгуна",
	CreateDaimyo:      "Создать даймё",
	CreateSamurai:     "Создать самурая",
	CreateCashManager: "Создать инкассатора",
	CreateCard:        "Создать карту",
	//Back:              "Назад",
}

//var Menu = tg.NewReplyKeyboardMarkup(
//	tg.NewButtonColumn(
//		tg.NewKeyboardButton(AdminMenu.CreateEntity),
//	)...,
//).WithResizeKeyboardMarkup()
