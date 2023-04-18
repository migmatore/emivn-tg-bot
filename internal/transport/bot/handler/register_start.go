package handler

import "github.com/mr-linch/go-tg/tgb"

func (h *Handler) registerStartHandler() {
	h.Message(h.StartHandler.Start, tgb.Command("start"))
}
