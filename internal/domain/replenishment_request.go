package domain

type ReplenishmentRequest struct {
	ReplenishmentRequestId int `json:"replenishment_request_id"`
	DaimyoId               int `json:"daimyo_id"`
	CardId                 int `json:"card_id"`
}
