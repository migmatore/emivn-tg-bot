package domain

type Card struct {
	CardId        int
	Name          string
	OwnerUsername string
	LastDigits    int
	DailyLimit    int
	Balance       float64
	BankTypeId    int
}

type CardDTO struct {
	Name          string
	OwnerUsername string
	LastDigits    int
	DailyLimit    int
	Balance       float64
	BankType      string
}
