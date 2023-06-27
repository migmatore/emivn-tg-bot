package role

import "context"

type RoleStorage interface {
	GetIdByName(ctx context.Context, role string) (int, error)
	GetById(ctx context.Context, roleId int) (string, error)
}

type RoleService struct {
	storage RoleStorage
}

func NewRoleService(s RoleStorage) *RoleService {
	return &RoleService{storage: s}
}
