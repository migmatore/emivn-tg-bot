package domain

type Card struct {
	CardId         int
	Name           string
	LastDigits     int
	DailyLimit     int
	DaimyoUsername string
	BankTypeId     int
}

type CardDTO struct {
	Name           string
	LastDigits     int
	DailyLimit     int
	DaimyoUsername string
	BankType       string
}
