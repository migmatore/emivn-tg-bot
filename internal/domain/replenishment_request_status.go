package domain

type ReplenishmentRequestStatus int

const (
	ActiveRequests ReplenishmentRequestStatus = iota
	ObjectionableRequests
	CompletedRequests
)

func (r ReplenishmentRequestStatus) String() string {
	switch r {
	case ActiveRequests:
		return "Активные"
	case ObjectionableRequests:
		return "Спорные"
	case CompletedRequests:
		return "Выполненные"
	default:
		return ""
	}
}
