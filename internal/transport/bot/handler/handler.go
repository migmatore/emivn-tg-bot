package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/transport/bot/handler/admin"
	"emivn-tg-bot/internal/transport/bot/handler/daimyo"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"log"
	"time"
)

type AuthService interface {
	//CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	GetRole(ctx context.Context, username string) (string, error)
}

type ShogunService interface {
	Create(ctx context.Context, dto domain.ShogunDTO) error
	GetAll(ctx context.Context) ([]*domain.ShogunDTO, error)
}

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
	//Notify(args domain.FuncArgs) (status domain.TaskStatus, when interface{})
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetByUsername(ctx context.Context, daimyoUsername string) ([]*domain.CardDTO, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
}

type SchedulerService interface {
	Configure(listeners domain.TaskFuncsMap, sleepDuration time.Duration)
	Add(ctx context.Context, dto domain.TaskDTO) error
	Run(ctx context.Context) error
}

// TODO: Refactor DI
type Deps struct {
	sessionManager *session.Manager[domain.Session]

	AuthService                 AuthService
	ShogunService               ShogunService
	DaimyoService               DaimyoService
	SamuraiService              SamuraiService
	CashManagerService          CashManagerService
	CardService                 CardService
	ReplenishmentRequestService ReplenishmentRequestService
	SchedulerService            SchedulerService
}

type Handler struct {
	*tgb.Router
	sessionManager *session.Manager[domain.Session]

	StartHandler  *start.StartHandler
	AdminHandler  *admin.AdminHandler
	DaimyoHandler *daimyo.DaimyoHandler

	schedulerService SchedulerService
}

func New(deps Deps) *Handler {
	sm := NewSessionManager()

	return &Handler{
		Router:         tgb.NewRouter(),
		sessionManager: sm.Manager,
		StartHandler:   start.NewStartHandler(sm.Manager, deps.AuthService),
		AdminHandler: admin.NewAdminHandler(
			sm.Manager,
			deps.ShogunService,
			deps.DaimyoService,
			deps.SamuraiService,
			deps.CashManagerService,
			deps.CardService,
		),
		DaimyoHandler: daimyo.NewDaimyoHandler(
			sm.Manager,
			deps.CardService,
			deps.ReplenishmentRequestService,
			deps.CashManagerService,
			deps.SchedulerService,
		),
		schedulerService: deps.SchedulerService,
	}
}

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	h.Router.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
		return tgb.HandlerFunc(func(ctx context.Context, update *tgb.Update) error {
			defer func(started time.Time) {
				log.Printf("%#v [%s]", update, time.Since(started))
			}(time.Now())

			return next.Handle(ctx, update)
		})
	}))
	h.registerSession()
	h.registerStartHandlers()
	h.registerAdminHandlers()
	h.registerDaimyoHandler()

	//listenersMap := domain.TaskFuncsMap{
	//	"notify_samurai": h.DaimyoHandler.Notify,
	//}
	//
	//h.schedulerService.Configure(listenersMap, time.Second*1)
	//
	//if err := h.schedulerService.Run(context.Background()); err != nil {
	//	log.Printf("scheduler error %v", err)
	//}

	return h.Router
}
