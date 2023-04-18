package domain

type Daimyo struct {
	DaimyoId     int     `json:"daimyo_id"`
	Username     string  `json:"username"`
	Nickname     string  `json:"nickname"`
	CardsBalance float32 `json:"cards_balance"`
	ShogunId     int     `json:"shogun_id"`
}
