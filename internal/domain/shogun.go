package domain

type Shogun struct {
	ShogunId int    `json:"shogun_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type ShogunDTO struct {
	Username string
	Nickname string
}
