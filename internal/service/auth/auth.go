package auth

import (
	"context"
	"emivn-tg-bot/internal/domain"
)

type AuthStorage interface {
	UserRole(ctx context.Context, username string) (string, error)
}

type EntityStorage interface {
	UpdateUsername(ctx context.Context, old string, new string) error
}

type ReferalStorage interface {
	GetByLink(ctx context.Context, link string) (domain.Referal, error)
	Delete(ctx context.Context, link string) error
}

type UserRoleStorage interface {
	UpdateUsername(ctx context.Context, old string, new string) error
}

type RoleStorage interface {
	GetById(ctx context.Context, roleId int) (string, error)
}

type AuthService struct {
	storage             AuthStorage
	shogunStorage       EntityStorage
	daimyoStorage       EntityStorage
	samuraiStorage      EntityStorage
	cashManagerStorage  EntityStorage
	controllerStorage   EntityStorage
	mainOperatorStorage EntityStorage
	referalStorage      ReferalStorage
	userRoleStorage     UserRoleStorage
	roleStorage         RoleStorage
}

func NewAuthService(
	s AuthStorage,
	shogunStorage EntityStorage,
	daimyoStorage EntityStorage,
	samuraiStorage EntityStorage,
	cashManagerStorage EntityStorage,
	controllerStorage EntityStorage,
	mainOperatorStorage EntityStorage,
	referalStorage ReferalStorage,
	userRoleStorage UserRoleStorage,
	roleStorage RoleStorage,
) *AuthService {
	return &AuthService{
		storage:             s,
		shogunStorage:       shogunStorage,
		daimyoStorage:       daimyoStorage,
		samuraiStorage:      samuraiStorage,
		cashManagerStorage:  cashManagerStorage,
		controllerStorage:   controllerStorage,
		mainOperatorStorage: mainOperatorStorage,
		referalStorage:      referalStorage,
		userRoleStorage:     userRoleStorage,
		roleStorage:         roleStorage,
	}
}

func (s *AuthService) CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error) {
	role, err := s.storage.UserRole(ctx, username)
	if err != nil {
		return false, err
	}

	if role == requiredRole.String() {
		return true, nil
	}

	return false, nil
}

// GetRole returns user role
func (s *AuthService) GetRole(ctx context.Context, username string) (string, error) {
	role, err := s.storage.UserRole(ctx, username)
	if err != nil {
		return "", err
	}

	return role, nil
}

func (s *AuthService) Auth(ctx context.Context, link string, username string) (string, error) {
	if link == "" {
		role, err := s.storage.UserRole(ctx, username)
		if err != nil {
			return "", err
		}

		return role, nil
	}

	role, _ := s.storage.UserRole(ctx, username)
	if role != "" {
		return role, nil
	}

	referal, err := s.referalStorage.GetByLink(ctx, link)
	if err != nil {
		return "", err
	}

	if err := s.userRoleStorage.UpdateUsername(ctx, link, username); err != nil {
		return "", err
	}

	roleName, err := s.roleStorage.GetById(ctx, referal.RoleId)
	if err != nil {
		return "", err
	}

	switch roleName {
	case domain.ShogunRole.String():
		if err := s.changeUsername(ctx, s.shogunStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.ShogunRole.String(), nil
	case domain.DaimyoRole.String():
		if err := s.changeUsername(ctx, s.daimyoStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.DaimyoRole.String(), nil
	case domain.SamuraiRole.String():
		if err := s.changeUsername(ctx, s.samuraiStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.SamuraiRole.String(), nil
	case domain.CashManagerRole.String():
		if err := s.changeUsername(ctx, s.cashManagerStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.CashManagerRole.String(), nil
	case domain.ControllerRole.String():
		if err := s.changeUsername(ctx, s.controllerStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.ControllerRole.String(), nil
	case domain.MainOperatorRole.String():
		if err := s.changeUsername(ctx, s.mainOperatorStorage, referal.Link, username); err != nil {
			return "", err
		}

		return domain.MainOperatorRole.String(), nil
	}

	return "", nil
}

func (s *AuthService) changeUsername(ctx context.Context, storage EntityStorage, link string, username string) error {
	if err := storage.UpdateUsername(ctx, link, username); err != nil {
		return err
	}

	if err := s.referalStorage.Delete(ctx, link); err != nil {
		return err
	}

	return nil
}
