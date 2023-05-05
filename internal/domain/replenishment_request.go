package domain

type ReplenishmentRequest struct {
	ReplenishmentRequestId int
	CashManagerUsername    string
	DaimyoUsername         string
	CardId                 int
	Amount                 float32
	StatusId               int
}

type ReplenishmentRequestDTO struct {
	CashManagerUsername string
	DaimyoUsername      string
	CardName            string
	Amount              float32
	Status              string
}
