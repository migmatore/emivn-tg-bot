package domain

type SamuraiTurnover struct {
	TurnoverId      int
	SamuraiUsername string
	StartDate       string
	InitialAmount   float64
	FinalAmount     float64
	Turnover        float64
	BankTypeId      int
}

type SamuraiTurnoverDTO struct {
	SamuraiUsername string
	StartDate       string
	InitialAmount   float64
	FinalAmount     float64
	Turnover        float64
	BankTypeName    string
}
