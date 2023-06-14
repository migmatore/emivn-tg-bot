package samurai

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"log"
	"time"
)

type SamuraiHandler struct {
	sessionManager *session.Manager[domain.Session]
}

func NewSamuraiHandler(sm *session.Manager[domain.Session]) *SamuraiHandler {
	return &SamuraiHandler{sessionManager: sm}
}

func (h *SamuraiHandler) MenuSelectionHandler(ctx context.Context, msg *tgb.MessageUpdate) error {

	return nil
}

func (h *SamuraiHandler) Notify(ctx context.Context, args domain.FuncArgs) (status domain.TaskStatus, when interface{}) {
	if id, ok := args["id"]; ok {

		var id tg.ChatID = tg.ChatID(id.(float64))

		if client, ok := ctx.Value(domain.TaskKey{}).(*tg.Client); ok {
			client.SendMessage(id, "hello").DoVoid(ctx)
		}

		return domain.TaskStatusWait, time.Now().Add(time.Second * 10)
	}
	//else {
	//	if client, ok := ctx.Value(domain.TaskKey{}).(*tg.Client); ok {
	//		var id tg.ChatID = 6109520093
	//
	//		client.SendMessage(id, "hello").DoVoid(ctx)
	//	}
	//}

	log.Print("Not found name arg in func args")

	return domain.TaskStatusWait, time.Now().Add(time.Second * 30)
}
