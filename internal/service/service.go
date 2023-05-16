package service

import (
	"emivn-tg-bot/internal/service/auth"
	"emivn-tg-bot/internal/service/card"
	"emivn-tg-bot/internal/service/cash_manager"
	"emivn-tg-bot/internal/service/daimyo"
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
	CashManagerStorage                cash_manager.CashManagerStorage
	CardStorage                       card.CardStorage
	ReplenishmentRequestStorage       replenishment_request.ReplenishmentRequestStorage
	ReplenishmentRequestStatusStorage replenishment_request_status.ReplenishmentRequestStatusStorage
	UserRoleStorage                   user_role.UserRoleStorage
	RoleStorage                       role.RoleStorage
	SchedulerStorage                  scheduler.SchedulerStorage
}

type Service struct {
	Auth                 *auth.AuthService
	Shogun               *shogun.ShogunService
	Daimyo               *daimyo.DaimyoService
	Samurai              *samurai.SamuraiService
	CashManager          *cash_manager.CashManagerService
	Card                 *card.CardService
	ReplenishmentRequest *replenishment_request.ReplenishmentRequestService
	SchedulerService     *scheduler.SchedulerService
}

func New(deps Deps) *Service {
	return &Service{
		Auth:   auth.NewAuthService(deps.AuthStorage),
		Shogun: shogun.NewShogunService(deps.Transactor, deps.ShogunStorage, deps.UserRoleStorage, deps.RoleStorage),
		Daimyo: daimyo.NewDaimyoService(
			deps.Transactor,
			deps.DaimyoStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		Samurai: samurai.NewSamuraiService(
			deps.Transactor,
			deps.SamuraiStorage,
			deps.UserRoleStorage,
			deps.RoleStorage,
		),
		CashManager: cash_manager.NewCashManagerService(
			deps.Transactor,
			deps.CashManagerStorage,
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
		SchedulerService: scheduler.New(deps.Transactor, deps.SchedulerStorage),
	}
}
