package domain

type ControllerTurnover struct {
	TurnoverId         int
	ControllerUsername string
	SamuraiUsername    string
	StartDate          string
	InitialAmount      float64
	FinalAmount        float64
	Turnover           float64
	BankTypeId         int
}

type ControllerTurnoverDTO struct {
	ControllerUsername string
	SamuraiUsername    string
	StartDate          string
	InitialAmount      float64
	FinalAmount        float64
	Turnover           float64
	BankTypeName       string
}
