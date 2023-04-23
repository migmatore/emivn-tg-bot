package domain

type Samurai struct {
	SamuraiId        int
	Username         string
	Nickname         string
	DaimyoId         int
	TurnoverPerShift float32
}

type SamuraiDTO struct {
	SamuraiId        int
	Username         string
	Nickname         string
	DaimyoUsername   string
	TurnoverPerShift float32
}
