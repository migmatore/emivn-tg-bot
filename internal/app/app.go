package app

import (
	"context"
	"emivn-tg-bot/internal/config"
	"emivn-tg-bot/internal/transport/bot"
	"emivn-tg-bot/pkg/logging"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"log"
	"net/http"
	"time"
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

		webhook := tgb.NewWebhook(
			bot,
			botClient,
			fullURL,
			tgb.WithWebhookLogger(logging.GetLogger(ctx)),
		)

		if err := webhook.Setup(ctx); err != nil {
			return err
		}

		mux := http.NewServeMux()
		mux.Handle("/webhook", webhook)

		return runServer(ctx, mux, "")
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

func runServer(ctx context.Context, handler http.Handler, listen string) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		<-ctx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("shutdown: %v", err)
		}
	}()

	log.Printf("listening on %s", ":8080")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
}
