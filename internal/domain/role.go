package domain

type Role int

const (
	AdminRole Role = iota + 1
	ShogunRole
	DaimyoRole
	SamuraiRole
	CashManagerRole
)

func (r Role) String() string {
	switch r {
	case AdminRole:
		return "Администратор"
	case ShogunRole:
		return "Сёгун"
	case DaimyoRole:
		return "Даймё"
	case SamuraiRole:
		return "Самурай"
	case CashManagerRole:
		return "Инкассатор"
	default:
		return ""
	}
}
