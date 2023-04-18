package domain

type Card struct {
	CardId          int    `json:"card_id"`
	IssuingBankInfo string `json:"issuing_bank_info"`
	DailyLimit      int    `json:"daily_limit"`
	//Username string
	//Nickname string
}
