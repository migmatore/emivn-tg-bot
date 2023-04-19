package domain

type UserRole struct {
	UserRoleId int    `json:"user_role_id"`
	Username   string `json:"username"`
	RoleId     int    `json:"role_id"`
	//Role string
}
