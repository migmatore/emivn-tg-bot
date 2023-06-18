package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/card"
	"emivn-tg-bot/internal/service/cash_manager"
	"emivn-tg-bot/internal/service/controller"
	"emivn-tg-bot/internal/service/daimyo"
	"emivn-tg-bot/internal/service/main_operator"
	"emivn-tg-bot/internal/service/replenishment_request"
	"emivn-tg-bot/internal/service/replenishment_request_status"
	"emivn-tg-bot/internal/service/role"
	"emivn-tg-bot/internal/service/samurai"
	"emivn-tg-bot/internal/service/scheduler"
	"emivn-tg-bot/internal/service/shogun"
	"emivn-tg-bot/internal/service/user_role"
	"emivn-tg-bot/internal/storage"
)

type Deps struct {
	Transactor storage.Transactor

	AuthStorage                       auth.AuthStorage
	ShogunStorage                     shogun.ShogunStorage
	DaimyoStorage                     daimyo.DaimyoStorage
	SamuraiStorage                    samurai.SamuraiStorage
	SamuraiTurnoverStorage            samurai.SamuraiTurnoverStorage
	CashManagerStorage                cash_manager.CashManagerStorage
	ControllerStorage                 controller.ControllerStorage
	ControllerTurnoverStorage         controller.ControllerTurnoverStorage
	MainOperatorStorage               main_operator.MainOperatorStorage
	CardStorage                       card.CardStorage
	ReplenishmentRequestStorage       replenishment_request.ReplenishmentRequestStorage
	ReplenishmentRequestStatusStorage replenishment_request_status.ReplenishmentRequestStatusStorage
	UserRoleStorage                   user_role.UserRoleStorage
	RoleStorage                       role.RoleStorage
	SchedulerStorage                  scheduler.SchedulerStorage
}

type Service struct {
	Transactor *TransactorService

	Auth                 *auth.AuthService
	Shogun               *shogun.ShogunService
	Daimyo               *daimyo.DaimyoService
	Samurai              *samurai.SamuraiService
	CashManager          *cash_manager.CashManagerService
	Controller           *controller.ControllerService
	MainOperator         *main_operator.MainOperatorService
	Card                 *card.CardService
	ReplenishmentRequest *replenishment_request.ReplenishmentRequestService
	SchedulerService     *scheduler.SchedulerService
}

func New(deps Deps) *Service {
	return &Service{
		Transactor: NewTransactorService(deps.Transactor),

		Auth:   auth.NewAuthService(deps.AuthStorage),
		Shogun: shogun.NewShogunService(deps.Transactor, deps.ShogunStorage, deps.UserRoleStorage, deps.RoleStorage),
		Daimyo: daimyo.NewDaimyoService(
			deps.Transactor,
			deps.DaimyoStorage,
			deps.SamuraiTurnoverStorage,
			deps.ControllerTurnoverStorage,
			deps.SamuraiStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		Samurai: samurai.NewSamuraiService(
			deps.Transactor,
			deps.SamuraiStorage,
			deps.SamuraiTurnoverStorage,
			deps.CardStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		CashManager: cash_manager.NewCashManagerService(
			deps.Transactor,
			deps.CashManagerStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		Controller: controller.New(
			deps.Transactor,
			deps.ControllerStorage,
			deps.ControllerTurnoverStorage,
			deps.CardStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		MainOperator: main_operator.New(
			deps.Transactor,
			deps.MainOperatorStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		Card: card.NewCardService(deps.Transactor, deps.CardStorage),
		ReplenishmentRequest: replenishment_request.NewReplenishmentRequestService(
			deps.ReplenishmentRequestStorage,
			deps.CashManagerStorage,
			deps.DaimyoStorage,
			deps.CardStorage,
			deps.ReplenishmentRequestStatusStorage,
		),
		SchedulerService: scheduler.New(deps.SchedulerStorage),
	}
}
