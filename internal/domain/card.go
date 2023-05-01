package domain

type Card struct {
	CardId         int
	Name           string
	LastDigits     int
	DailyLimit     int
	DaimyoUsername string
}

type CardDTO struct {
	Name           string
	LastDigits     int
	DailyLimit     int
	DaimyoUsername string
}
