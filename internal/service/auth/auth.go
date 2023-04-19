package auth

import "context"

type AuthStorage interface {
}

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Auth(ctx context.Context, username string) (bool, error) {
	if username == "migmatore" {
		return true, nil
	} else {
		return false, nil
	}
}
