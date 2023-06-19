package domain

type Card struct {
	CardId         int
	Name           string
	DaimyoUsername string
	LastDigits     int
	DailyLimit     int
	Balance        float64
	BankTypeId     int
}

type CardDTO struct {
	Name           string
	DaimyoUsername string
	LastDigits     int
	DailyLimit     int
	Balance        float64
	BankType       string
}
