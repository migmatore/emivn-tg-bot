package domain

type CashManager struct {
	CashManagerId          int    `json:"cash_manager_id"`
	Username               string `json:"username"`
	Nickname               string `json:"nickname"`
	ReplenishmentRequestId int    `json:"replenishment_request_id"`
}
