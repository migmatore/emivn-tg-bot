package domain

// TODO: menu code refactoring(with back button)

var AdminMainMenu = struct {
	Cards        string
	RequestsList string
	Report       string
	Hierarchy    string
	CreateGARA   string
	UploadData   string
	Deposits     string
}{
	Cards:        "Карты",
	RequestsList: "Список заявок",
	Report:       "Запросить отчет",
	Hierarchy:    "Иерархия",
	CreateGARA:   "Создать ГАРА",
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

var ShogunMainMenu = struct {
	Requests   string
	Cards      string
	Report     string
	Hierarchy  string
	GARA       string
	UploadData string
	Deposits   string
}{
	Requests:   "Список заявок",
	Cards:      "Карты",
	Report:     "Запросить отчет",
	Hierarchy:  "Иерархия",
	GARA:       "ГАРА",
	UploadData: "Загрузить данные",
	Deposits:   "Пополнения",
}

var DaimyoMainMenu = struct {
	MakeReplenishmentRequest string
	Requests                 string
	CardLimit                string
	Report                   string
	Hierarchy                string
}{
	MakeReplenishmentRequest: "Запросить пополнение",
	Requests:                 "Заявки",
	CardLimit:                "Лимит по карте",
	Report:                   "Отчет",
	Hierarchy:                "Иерархия",
}

var SamuraiMainMenu = struct {
	EnterData string
}{
	EnterData: "Ввести данные на конец смены",
}

var SamuraiEnterDataMenu = struct {
	ChooseBank string
}{
	ChooseBank: "Выбрать банк",
}

var CashManagerMainMenu = struct {
	Requests           string
	WithdrawalRequests string
	RemainingFunds     string
	CurrentBalance     string
	ReplenishmentList  string
}{
	Requests:           "Заявки",
	WithdrawalRequests: "Заявки на вывод",
	RemainingFunds:     "Остаток средств",
	CurrentBalance:     "Текущий остаток",
	ReplenishmentList:  "Список пополнений",
}

var ControllerMainMenu = struct {
	EnterData string
}{
	EnterData: "Ввести данные на конец смены",
}
