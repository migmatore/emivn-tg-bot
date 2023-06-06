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
	if client, ok := ctx.Value(domain.TaskKey{}).(*tg.Client); ok {
		var id tg.ChatID = 6109520093

		client.SendMessage(id, "hello").DoVoid(ctx)
	}

	//if name, ok := args["name"]; ok {
	//	log.Println("PrintWithArgs:", time.Now(), name)
	//	return domain.TaskStatusDeferred, time.Now().Add(time.Second * 10)
	//}

	log.Print("Not found name arg in func args")

	return domain.TaskStatusDeferred, time.Now().Add(time.Second * 5)
}
