package bot

import (
	"context"
	"emivn-tg-bot/internal/config"
	"emivn-tg-bot/internal/storage/psql"
	"emivn-tg-bot/internal/transport/bot/handler"
	"emivn-tg-bot/pkg/logging"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"net/http"
	"os"
	"os/signal"
)

type Bot struct {
	router    *tgb.Router
	pool      psql.AtomicPoolClient
	config    *config.Config
	scheduler *handler.Scheduler
}

func New(router *tgb.Router, pool psql.AtomicPoolClient, config *config.Config, scheduler *handler.Scheduler) *Bot {
	return &Bot{
		router:    router,
		pool:      pool,
		config:    config,
		scheduler: scheduler,
	}
}

func (b *Bot) Run(ctx context.Context) error {
	botClient, err := newBotClient(ctx, b.config)
	if err != nil {
		return fmt.Errorf("Init bot client failed with %w", err)
	}

	if b.config.AppConfig.BaseURL != "" {
		fullURL := b.config.AppConfig.BaseURL + "/webhook"
		logging.GetLogger(ctx).Info("Start webhook server...")

		webhook := tgb.NewWebhook(
			b.router,
			botClient,
			fullURL,
			tgb.WithWebhookLogger(logging.GetLogger(ctx)),
		)

		if err := webhook.Setup(ctx); err != nil {
			return err
		}

		mux := http.NewServeMux()
		mux.Handle("/webhook", webhook)

		return b.runServer(ctx, mux)
	} else {
		logging.GetLogger(ctx).Info("start polling...")

		//go b.scheduler.Run(context.WithValue(ctx, domain.TaskKey{}, botClient))

		err := tgb.NewPoller(
			b.router,
			botClient,
		).Run(
			ctx,
		)
		if err != nil {
			return err
		}

		return nil
	}
}

func (b *Bot) runServer(ctx context.Context, handler http.Handler) error {
	server := &http.Server{
		Addr:    ":" + b.config.Listen.Port,
		Handler: handler,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logging.GetLogger(ctx).Info("Close all database connections...")
		b.pool.Close()
		logging.GetLogger(ctx).Info("All database connections have been closed!")

		if err := server.Shutdown(ctx); err != nil {
			logging.GetLogger(ctx).Errorf("Server is not shutting down! Reason: %v", err)
		}

		logging.GetLogger(ctx).Info("Server has successfully shut down!")

		close(idleConnsClosed)
	}()

	logging.GetLogger(ctx).Info("listening on %s", ":"+b.config.Listen.Port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}

	return nil
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
