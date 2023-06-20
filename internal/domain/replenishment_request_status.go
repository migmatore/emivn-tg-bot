package domain

type ReplenishmentRequestStatus int

const (
	ActiveRequest ReplenishmentRequestStatus = iota
	ObjectionableRequest
	CompletedRequests
)

func (r ReplenishmentRequestStatus) String() string {
	switch r {
	case ActiveRequest:
		return "Активные"
	case ObjectionableRequest:
		return "Спорные"
	case CompletedRequests:
		return "Выполненные"
	default:
		return ""
	}
}
