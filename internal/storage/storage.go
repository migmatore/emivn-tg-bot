package storage

import (
	"emivn-tg-bot/internal/storage/auth"
	"emivn-tg-bot/internal/storage/card"
	"emivn-tg-bot/internal/storage/cash_manager"
	"emivn-tg-bot/internal/storage/daimyo"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/internal/storage/replenishment_request"
	"emivn-tg-bot/internal/storage/replenishment_request_status"
	"emivn-tg-bot/internal/storage/role"
	"emivn-tg-bot/internal/storage/samurai"
	"emivn-tg-bot/internal/storage/shogun"
	"emivn-tg-bot/internal/storage/user_role"
)

type Storage struct {
	Transactor *Transact

	Auth                              *auth.AuthStorage
	Shogun                            *shogun.ShogunStorage
	Daimyo                            *daimyo.DaimyoStorage
	Samurai                           *samurai.SamuraiStorage
	CashManager                       *cash_manager.CashManagerStorage
	Card                              *card.CardStorage
	ReplenishmentRequest              *replenishment_request.ReplenishmentRequestStorage
	ReplenishmentRequestStatusStorage *replenishment_request_status.ReplenishmentRequestStatusStorage
	UserRole                          *user_role.UserRoleStorage
	Role                              *role.RoleStorage
}

func New(pool psql.AtomicPoolClient) *Storage {
	return &Storage{
		Transactor:                        NewTransactor(pool),
		Auth:                              auth.NewAuthStorage(pool),
		Shogun:                            shogun.NewShogunStorage(pool),
		Daimyo:                            daimyo.NewDaimyoStorage(pool),
		Samurai:                           samurai.NewSamuraiStorage(pool),
		CashManager:                       cash_manager.NewCashManagerStorage(pool),
		Card:                              card.NewCardStorage(pool),
		ReplenishmentRequest:              replenishment_request.NewReplenishmentRequestStorage(pool),
		ReplenishmentRequestStatusStorage: replenishment_request_status.NewReplenishmentRequestStatusStorage(pool),
		UserRole:                          user_role.NewUserRoleStorage(pool),
		Role:                              role.NewRoleStorage(pool),
	}
}
