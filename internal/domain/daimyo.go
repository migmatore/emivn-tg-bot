package domain

type Daimyo struct {
	DaimyoId     int
	Username     string
	Nickname     string
	CardsBalance float32
	ShogunId     int
}

type DaimyoDTO struct {
	Username       string
	Nickname       string
	CardsBalance   float32
	ShogunUsername string
}
