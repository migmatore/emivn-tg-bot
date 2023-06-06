package domain

type Samurai struct {
	Username         string
	Nickname         string
	DaimyoUsername   string
	TurnoverPerShift float32
	ChatId           *int
}

type SamuraiDTO struct {
	SamuraiId        int
	Username         string
	Nickname         string
	DaimyoUsername   string
	TurnoverPerShift float32
	ChatId           *int
}
