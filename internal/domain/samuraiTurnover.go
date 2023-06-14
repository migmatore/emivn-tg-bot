package domain

import "time"

type SamuraiTurnover struct {
	TurnoverId      int
	SamuraiUsername string
	StartDate       time.Time
	InitialAmount   float64
	FinalAmount     float64
	Turnover        float64
	BankTypeId      int
}

type SamuraiTurnoverDTO struct {
	SamuraiUsername string
	StartDate       time.Time
	InitialAmount   float64
	FinalAmount     float64
	Turnover        float64
	BankTypeName    string
}
