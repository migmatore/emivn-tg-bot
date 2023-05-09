package domain

type ReplenishmentRequestStatus int

const (
	ActiveRequest ReplenishmentRequestStatus = iota
	ObjectionableRequest
)

func (r ReplenishmentRequestStatus) String() string {
	switch r {
	case ActiveRequest:
		return "Активный"
	case ObjectionableRequest:
		return "Спорный"
	default:
		return ""
	}
}
