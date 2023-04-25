package auth

import (
	"context"
)

type AuthStorage interface {
	UserRole(ctx context.Context, username string) (string, error)
}

type AuthService struct {
	storage AuthStorage
}

func NewAuthService(s AuthStorage) *AuthService {
	return &AuthService{storage: s}
}

//func (s *AuthService) CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error) {
//	role, err := s.storage.UserRole(ctx, username)
//	if err != nil {
//		return false, err
//	}
//
//	if role == requiredRole.String() {
//		return true, nil
//	}
//
//	return false, nil
//}

// GetRole returns user role
func (s *AuthService) GetRole(ctx context.Context, username string) (string, error) {
	role, err := s.storage.UserRole(ctx, username)
	if err != nil {
		return "", err
	}

	return role, nil
}
