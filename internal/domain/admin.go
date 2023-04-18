package domain

type Admin struct {
	AdminId  int    `json:"admin_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}
