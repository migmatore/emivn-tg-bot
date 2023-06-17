package domain

type Daimyo struct {
	Username       string
	Nickname       string
	CardsBalance   float32
	ShogunUsername string
}

type DaimyoDTO struct {
	Username       string
	Nickname       string
	CardsBalance   float32
	ShogunUsername string
}

type SamuraiReport struct {
	SamuraiUsername    string
	ControllerTurnover float64
	SamuraiTurnover    float64
	BankTypeName       string
}

type SamuraiReportDTO struct {
	SamuraiUsername    string
	ControllerTurnover float64
	SamuraiTurnover    float64
	BankTypeName       string
}
