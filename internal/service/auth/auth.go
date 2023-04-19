package auth

import (
	"context"
)

type AuthStorage interface {
	CheckAuth(ctx context.Context, username string) (bool, error)
}

type AuthService struct {
	storage AuthStorage
}

func NewAuthService(s AuthStorage) *AuthService {
	return &AuthService{storage: s}
}

func (s *AuthService) Auth(ctx context.Context, username string) (bool, error) {
	return s.storage.CheckAuth(ctx, username)
}
