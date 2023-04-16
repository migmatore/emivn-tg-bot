package app

import (
	"context"
	"emivn-tg-bot/internal/config"
	"emivn-tg-bot/internal/transport/bot"
	"emivn-tg-bot/pkg/logging"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
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
	botClient, err := newBotClient(ctx, a.config)
	if err != nil {
		return fmt.Errorf("Init bot client failed with %w", err)
	}

	bot, err := bot.New(ctx)
	if err != nil {
		return fmt.Errorf("Init bot failed with %w", err)
	}

	if a.config.AppConfig.BaseURL != "" {
		fullURL := a.config.AppConfig.BaseURL + "/webhook"
		logging.GetLogger(ctx).Info("Start webhook server...")

		return tgb.NewWebhook(
			bot,
			botClient,
			fullURL,
			tgb.WithWebhookLogger(logging.GetLogger(ctx)),
		).Run(
			ctx,
			":8080",
		)
	} else {
		logging.GetLogger(ctx).Info("start polling...")

		return tgb.NewPoller(
			bot,
			botClient,
		).Run(
			ctx,
		)
	}
}

func newBotClient(ctx context.Context, config *config.Config) (*tg.Client, error) {
	client := tg.New(config.AppConfig.Token)

	me, err := client.Me(ctx)
	if err != nil {
		return nil, err
	}

	logging.GetLogger(ctx).Infof("Authorized as bot https://t.me/%s", me.Username)

	return client, nil
}
