package app

import (
	"context"
	"emivn-tg-bot/internal/config"
	"emivn-tg-bot/internal/service"
	"emivn-tg-bot/internal/storage"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/internal/transport/bot"
	"emivn-tg-bot/internal/transport/bot/handler"
	"emivn-tg-bot/pkg/logging"
	"fmt"
)

type App struct {
	config *config.Config
}

func NewApp(config *config.Config) App {
	return App{
		config: config,
	}
}

func (a *App) Run(ctx context.Context) error {
	logging.GetLogger(ctx).Info("Start app initializing...")

	logging.GetLogger(ctx).Info("Database connection initializing...")
	pool, err := psql.NewPostgres(ctx, 3, a.config)
	if err != nil {
		logging.GetLogger(ctx).Fatalf("Failed to initialize db connection: %s", err.Error())
		return err
	}

	logging.GetLogger(ctx).Info("Database reconnection goroutine initializing...")
	go pool.Reconnect(ctx, a.config)

	storages := storage.New(pool)

	services := service.New(service.Deps{
		Transactor:                        storages.Transactor,
		AuthStorage:                       storages.Auth,
		ShogunStorage:                     storages.Shogun,
		DaimyoStorage:                     storages.Daimyo,
		SamuraiStorage:                    storages.Samurai,
		CashManagerStorage:                storages.CashManager,
		UserRoleStorage:                   storages.UserRole,
		RoleStorage:                       storages.Role,
		CardStorage:                       storages.Card,
		ReplenishmentRequestStorage:       storages.ReplenishmentRequest,
		ReplenishmentRequestStatusStorage: storages.ReplenishmentRequestStatusStorage,
	})

	handlers := handler.New(handler.Deps{
		TransactorService:           services.Transactor,
		AuthService:                 services.Auth,
		ShogunService:               services.Shogun,
		DaimyoService:               services.Daimyo,
		SamuraiService:              services.Samurai,
		CashManagerService:          services.CashManager,
		CardService:                 services.Card,
		ReplenishmentRequestService: services.ReplenishmentRequest,
	})

	router := handlers.Init(ctx)

	bot := bot.New(router, pool, a.config)

	if err := bot.Run(ctx); err != nil {
		return fmt.Errorf("Init bot failed with %w", err)
	}

	return nil
}
