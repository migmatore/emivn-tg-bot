package auth

import (
	"context"
	"emivn-tg-bot/internal/domain"
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

func (s *AuthService) Auth(ctx context.Context, username string, requiredRole domain.Role) (bool, error) {
	role, err := s.storage.UserRole(ctx, username)
	if err != nil {
		return false, err
	}

	if role == requiredRole.String() {
		return true, nil
	}

	return false, nil
}

func (s *AuthService) Redirect(ctx context.Context, username string) domain.SessionStep {
	role, err := s.storage.UserRole(ctx, username)
	if err != nil {
		return domain.SessionStepStart
	}
	switch role {
	case domain.AdminRole.String():
		return domain.SessionStepAdminRole
	case domain.ShogunRole.String():
		//return domain.SessionStepReadData
	}

	return domain.SessionStepStart
}
