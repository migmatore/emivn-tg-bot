package domain

// TODO: menu code refactoring(with back button)

// Admin

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

var AdminCardsMenu = struct {
	CreateCard string
	CardsList  string
	Limit      string
	Balance    string
}{
	CreateCard: "Создать карту",
	CardsList:  "Список карт",
	Limit:      "Лимит",
	Balance:    "Баланс",
}

var AdminHierarchyMenu = struct {
	CreateEntity    string
	InSubordination string
}{
	CreateEntity:    "Создать",
	InSubordination: "В подчинении",
}

var AdminCreateEntityMenu = struct {
	CreateShogun       string
	CreateDaimyo       string
	CreateSamurai      string
	CreateCashManager  string
	CreateController   string
	CreateMainOperator string
	CreateCard         string
	Back               string
}{
	CreateShogun:       "Сёгун",
	CreateDaimyo:       "Дайме",
	CreateSamurai:      "Самурай",
	CreateCashManager:  "Инкассатор",
	CreateController:   "Контролёр",
	CreateMainOperator: "Главный оператор",
}

// Shogun

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

var ShogunCardsMenu = struct {
	CreateCard string
	CardsList  string
	Limit      string
	Balance    string
}{
	CreateCard: "Создать карту",
	CardsList:  "Список карт",
	Limit:      "Лимит",
	Balance:    "Баланс",
}

var ShogunHierarchyMenu = struct {
	CreateEntity    string
	InSubordination string
}{
	CreateEntity:    "Создать",
	InSubordination: "В подчинении",
}

var ShogunCreateEntityMenu = struct {
	CreateDaimyo       string
	CreateSamurai      string
	CreateCashManager  string
	CreateMainOperator string
}{
	CreateDaimyo:       "Дайме",
	CreateSamurai:      "Самурай",
	CreateCashManager:  "Инкассатор",
	CreateMainOperator: "Главный оператор",
}

var ShogunSubordinationMenu = struct {
	Daimyo       string
	Samurai      string
	MainOperator string
	CashManager  string
}{
	Daimyo:       "Дайме",
	Samurai:      "Самураи",
	MainOperator: "Главные операторы",
	CashManager:  "Инкассаторы",
}

// Daimyo

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

var DaimyoReportMenu = struct {
	EnterShiftData string
	ReportRequest  string
}{
	EnterShiftData: "Ввести данные за смену",
	ReportRequest:  "Запросить отчет",
}

var DaimyoReportPeriodMenu = struct {
	ForShift  string
	ForPeriod string
}{
	ForShift:  "За смену с 8 до 12",
	ForPeriod: "За период",
}

var DaimyoHierarchyMenu = struct {
	CreateSamurai   string
	InSubordination string
}{
	CreateSamurai:   "Создать Самурая",
	InSubordination: "В подчинении",
}

// Samurai

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

// Cash manager

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

// Controller

var ControllerMainMenu = struct {
	EnterData string
}{
	EnterData: "Ввести данные на конец смены",
}
