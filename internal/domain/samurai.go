package domain

type Samurai struct {
	SamuraiId        int     `json:"samurai_id"`
	Username         string  `json:"username"`
	Nickname         string  `json:"nickname"`
	DaimyoId         int     `json:"daimyo_id"`
	TurnoverPerShift float32 `json:"turnover_per_shift"`
}
