package handler

import (
	"context"
	"github.com/mr-linch/go-tg/tgb"
)

type StartHandler struct {
}

func NewStartHandler() *StartHandler {
	return &StartHandler{}
}

func (s *StartHandler) Start(ctx context.Context, msg *tgb.MessageUpdate) error {
	return msg.Answer("hello").DoVoid(ctx)
}
