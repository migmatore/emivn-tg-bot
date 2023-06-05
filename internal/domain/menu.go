package domain

// TODO: menu code refactoring(with back button)

var AdminMainMenu = struct {
	Cards        string
	RequestsList string
	Report       string
	Hierarchy    string
	UploadData   string
	Deposits     string
}{
	Cards:        "Карты",
	RequestsList: "Список заявок",
	Report:       "Запросить отчет",
	Hierarchy:    "Иерархия",
	UploadData:   "Выгрузить данные",
	Deposits:     "Пополнения",
}

var AdminHierarchyMenu = struct {
	CreateEntity    string
	InSubordination string
}{
	CreateEntity:    "Создать",
	InSubordination: "В подчинении",
}

var AdminCreateEntityMenu = struct {
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

var DaimyoMenu = struct {
	MakeReplenishmentRequest string
}{
	MakeReplenishmentRequest: "Сделать запрос на пополнение карты",
}
