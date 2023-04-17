package app

import (
	"context"
	"emivn-tg-bot/internal/config"
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

	//storages := storage.New(pool)

	//services := service.New(service.Deps{
	//	Transactor: storages.Transactor,
	//})

	handlers := handler.New(handler.Deps{})

	router := handlers.Init(ctx)

	bot := bot.New(router, pool, a.config)

	if err := bot.Run(ctx); err != nil {
		return fmt.Errorf("Init bot failed with %w", err)
	}

	return nil
}
