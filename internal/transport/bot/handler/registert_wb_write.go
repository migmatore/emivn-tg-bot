package handler

import (
	"github.com/mr-linch/go-tg/tgb"
)

func (h *Handler) registerDbWriteHandler() {
	h.Message(h.DbWriteHandler.Menu, tgb.Command("db_menu")).
		Message(h.DbWriteHandler.Write, tgb.TextEqual("Write")).
		Message(h.DbWriteHandler.Read, tgb.TextEqual("Read"))
}
