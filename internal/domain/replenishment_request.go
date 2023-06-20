package domain

type ReplenishmentRequest struct {
	ReplenishmentRequestId int
	CashManagerUsername    string
	OwnerUsername          string
	CardId                 int
	Amount                 float32
	StatusId               int
}

type ReplenishmentRequestDTO struct {
	CashManagerUsername string
	OwnerUsername       string
	CardName            string
	Amount              float32
	Status              string
}
