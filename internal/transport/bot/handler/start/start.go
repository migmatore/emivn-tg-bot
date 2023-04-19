package start

import (
	"context"
	"github.com/mr-linch/go-tg/tgb"
)

type AuthService interface {
	Auth(ctx context.Context, username string) (bool, error)
}

type StartHandler struct {
	AuthService AuthService
}

func NewStartHandler(s AuthService) *StartHandler {
	return &StartHandler{AuthService: s}
}

func (s *StartHandler) Start(ctx context.Context, msg *tgb.MessageUpdate) error {
	auth, _ := s.AuthService.Auth(ctx, string(msg.From.Username))

	if auth {
		return msg.Answer("You are welcume!!!!!").DoVoid(ctx)
	}

	return msg.Answer("You are not allowed!!!!").DoVoid(ctx)
}
