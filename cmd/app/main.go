package main

import (
	"context"
	"emivn-tg-bot/internal/app"
	"emivn-tg-bot/internal/config"
	"emivn-tg-bot/pkg/logging"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, cancel = signal.NotifyContext(ctx, os.Interrupt, os.Kill)

	logger := logging.GetLogger(ctx)
	ctx = logging.ContextWithLogger(ctx, logger)
	logger.Info("Logger initializing")

	logger.Info("Config initializing")
	cfg := config.GetConfig(ctx)

	logger.SetLoggingLevel(cfg.AppConfig.LogLevel)

	a := app.NewApp(cfg)
	if err := a.Run(ctx); err != nil {
		logger.Error("Run failed")
		logger.Fatalln(err)
	}
}
