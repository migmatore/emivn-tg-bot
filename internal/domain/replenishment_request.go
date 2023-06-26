package domain

type ReplenishmentRequest struct {
	ReplenishmentRequestId int
	CashManagerUsername    string
	OwnerUsername          string
	CardId                 int
	RequiredAmount         float32
	ActualAmount           float32
	StatusId               int
	CreationDate           string
	CreationTime           string
}

type ReplenishmentRequestDTO struct {
	CashManagerUsername string
	OwnerUsername       string
	CardName            string
	RequiredAmount      float32
	ActualAmount        float32
	Status              string
}
