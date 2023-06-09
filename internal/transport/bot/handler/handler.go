package handler

import (
	"context"
	"emivn-tg-bot/internal/domain"
	"emivn-tg-bot/internal/transport/bot/handler/admin"
	"emivn-tg-bot/internal/transport/bot/handler/cash_manager"
	"emivn-tg-bot/internal/transport/bot/handler/controller"
	"emivn-tg-bot/internal/transport/bot/handler/daimyo"
	"emivn-tg-bot/internal/transport/bot/handler/main_operator"
	"emivn-tg-bot/internal/transport/bot/handler/samurai"
	"emivn-tg-bot/internal/transport/bot/handler/shogun"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"time"
)

type TransactorService interface {
	WithinTransaction(ctx context.Context, txFunc func(context context.Context) error) error
}

type AuthService interface {
	//CheckAuthRole(ctx context.Context, username string, requiredRole domain.Role) (bool, error)
	Auth(ctx context.Context, link string, username string) (string, error)
}

type ShogunService interface {
	Create(ctx context.Context, dto domain.ShogunDTO) error
	GetAll(ctx context.Context) ([]*domain.ShogunDTO, error)
	GetByNickname(ctx context.Context, nickname string) (domain.ShogunDTO, error)
}

type DaimyoService interface {
	Create(ctx context.Context, dto domain.DaimyoDTO) error
	GetAll(ctx context.Context) ([]*domain.DaimyoDTO, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.DaimyoDTO, error)
	GetByUsername(ctx context.Context, username string) (domain.DaimyoDTO, error)
	GetByNickname(ctx context.Context, nickname string) (domain.DaimyoDTO, error)
	CreateSamuraiReport(ctx context.Context, daimyoUsername string, date string) ([]string, error)
	CreateSamuraiReportWithPeriod(ctx context.Context, daimyoUsername string, startDate string,
		endDate string) ([]string, error)
}

type SamuraiService interface {
	Create(ctx context.Context, dto domain.SamuraiDTO) error
	SetChatId(ctx context.Context, username string, id tg.ChatID) error
	GetByUsername(ctx context.Context, username string) (domain.SamuraiDTO, error)
	GetByNickname(ctx context.Context, nickname string) (domain.SamuraiDTO, error)
	CreateTurnover(ctx context.Context, dto domain.SamuraiTurnoverDTO) error
	GetAllByDaimyo(ctx context.Context, daimyoUsername string) ([]*domain.SamuraiDTO, error)
	GetAllByDaimyoNickname(ctx context.Context, nickname string) ([]*domain.SamuraiDTO, error)
}

type CashManagerService interface {
	Create(ctx context.Context, dto domain.CashManagerDTO) error
	SetChatId(ctx context.Context, username string, id tg.ChatID) error
}

type ControllerService interface {
	Create(ctx context.Context, dto domain.ControllerDTO) error
	CreateTurnover(ctx context.Context, dto domain.ControllerTurnoverDTO) error
}

type MainOperatorService interface {
	Create(ctx context.Context, dto domain.MainOperatorDTO) error
	GetByUsername(ctx context.Context, username string) (domain.MainOperatorDTO, error)
}

type CardService interface {
	Create(ctx context.Context, dto domain.CardDTO) error
	GetAllByUsername(ctx context.Context, bankName string, ownerUsername string) ([]*domain.CardDTO, error)
	GetAllByShogun(ctx context.Context, shogunUsername string) ([]*domain.CardDTO, error)
	GetByUsername(ctx context.Context, ownerUsername string) (domain.CardDTO, error)
	GetByName(ctx context.Context, name string) (domain.CardDTO, error)
	ChangeLimit(ctx context.Context, name string, limit int) error
	GetBankNames(ctx context.Context) ([]*domain.BankDTO, error)
	GetCardsBalancesByShogun(ctx context.Context, shogunUsername string) ([]string, error)
	GetLimits(ctx context.Context, owner string) ([]string, error)
}

type ReplenishmentRequestService interface {
	Create(ctx context.Context, dto domain.ReplenishmentRequestDTO) (tg.ChatID, error)
	CheckIfExists(ctx context.Context, cardName string) (bool, error)
	GetAllByCashManager(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequestDTO, error)
	GetAllByOwner(ctx context.Context, username string, status string) ([]*domain.ReplenishmentRequestDTO, error)
	GetByCardName(ctx context.Context, name string) (domain.ReplenishmentRequestDTO, error)
	ChangeStatus(ctx context.Context, cardName string, status string) error
	ConfirmRequest(ctx context.Context, dto domain.ReplenishmentRequestDTO) error
}

type ReferalService interface {
	Create(ctx context.Context, link string, role string) error
}

// TODO: Refactor DI
type Deps struct {
	sessionManager *session.Manager[domain.Session]

	AuthService                 AuthService
	ShogunService               ShogunService
	DaimyoService               DaimyoService
	SamuraiService              SamuraiService
	CashManagerService          CashManagerService
	ControllerService           ControllerService
	MainOperatorService         MainOperatorService
	CardService                 CardService
	ReplenishmentRequestService ReplenishmentRequestService
	ReferalService              ReferalService

	TransactorService TransactorService
	SchedulerService  SchedulerService
}

type Handler struct {
	*tgb.Router
	sessionManager *session.Manager[domain.Session]

	StartHandler        *start.StartHandler
	AdminHandler        *admin.AdminHandler
	DaimyoHandler       *daimyo.DaimyoHandler
	CashManagerHandler  *cash_manager.CashManagerHandler
	SamuraiHandler      *samurai.SamuraiHandler
	ShogunHandler       *shogun.ShogunHandler
	ControllerHandler   *controller.ControllerHandler
	MainOperatorHandler *main_operator.MainOperatorHandler

	scheduler *Scheduler
}

func New(deps Deps) *Handler {
	sm := NewSessionManager()
	scheduler := NewScheduler(deps.TransactorService, deps.SchedulerService)

	return &Handler{
		Router:         tgb.NewRouter(),
		sessionManager: sm.Manager,
		scheduler:      scheduler,
		StartHandler: start.NewStartHandler(
			sm.Manager,
			deps.AuthService,
			deps.SamuraiService,
			deps.CashManagerService,
			scheduler,
		),
		AdminHandler: admin.NewAdminHandler(
			sm.Manager,
			deps.ShogunService,
			deps.DaimyoService,
			deps.SamuraiService,
			deps.CashManagerService,
			deps.ControllerService,
			deps.MainOperatorService,
			deps.CardService,
			deps.ReferalService,
		),
		DaimyoHandler: daimyo.NewDaimyoHandler(
			sm.Manager,
			deps.CardService,
			deps.DaimyoService,
			deps.ReplenishmentRequestService,
			deps.CashManagerService,
			deps.SamuraiService,
			deps.ReferalService,
			scheduler,
		),
		CashManagerHandler: cash_manager.New(
			sm.Manager,
			deps.ReplenishmentRequestService,
			deps.CardService,
			deps.DaimyoService,
			deps.MainOperatorService,
		),
		SamuraiHandler: samurai.NewSamuraiHandler(
			sm.Manager,
			deps.CardService,
			deps.SamuraiService,
		),
		ShogunHandler: shogun.NewShogunHandler(
			sm.Manager,
			deps.DaimyoService,
			deps.SamuraiService,
			deps.CashManagerService,
			deps.MainOperatorService,
			deps.CardService,
			deps.ReferalService,
		),
		ControllerHandler: controller.New(
			sm.Manager,
			deps.ControllerService,
			deps.CardService,
			deps.DaimyoService,
			deps.SamuraiService,
		),
		MainOperatorHandler: main_operator.New(
			sm.Manager,
			deps.CardService,
			deps.ReplenishmentRequestService,
			deps.MainOperatorService,
		),
	}
}

func (h *Handler) Init(ctx context.Context) (*tgb.Router, *Scheduler) {
	listenersMap := domain.TaskFuncsMap{
		"notify_samurai":    h.SamuraiHandler.Notify,
		"change_card_limit": h.DaimyoHandler.NotifyCardLimitChange,
	}

	h.scheduler.Configure(listenersMap, time.Second*1)
	//if err := h.scheduler.Run(ctx); err != nil {
	//	logging.GetLogger(ctx).Errorf("scheduler error %v", err)
	//}

	//h.Router.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
	//	return tgb.HandlerFunc(func(ctx context.Context, update *tgb.Update) error {
	//		defer func(started time.Time) {
	//			log.Printf("%#v [%s]", update, time.Since(started))
	//		}(time.Now())
	//
	//		return next.Handle(ctx, update)
	//	})
	//}))
	h.registerSession()
	h.registerStartHandlers()
	h.registerAdminHandlers()
	h.registerDaimyoHandler()
	h.registerSamuraiHandler()
	h.registerControllerHandler()
	h.registerShogunHandler()
	h.registerCashManagerHandlers()
	h.registerMainOperatorHandlers()

	return h.Router, h.scheduler
}
