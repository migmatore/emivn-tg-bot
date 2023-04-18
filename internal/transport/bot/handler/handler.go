package handler

import (
	"context"
	"emivn-tg-bot/internal/transport/bot/handler/db_write"
	"emivn-tg-bot/internal/transport/bot/handler/start"
	"fmt"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
	"regexp"
	"strconv"
)

type Deps struct {
	session *session.Manager[Session]
}

type Handler struct {
	*tgb.Router

	StartHandler   *start.StartHandler
	DbWriteHandler *db_write.DbWriteHandler
}

func New(deps Deps) *Handler {
	return &Handler{
		Router:         tgb.NewRouter(),
		StartHandler:   start.NewStartHandler(),
		DbWriteHandler: db_write.NewDbWriteHandler(),
	}
}

type SessionStep int8

const (
	SessionStepInit = iota
	SessionStepName
	SessionStepAge
	SessionStepGender
)

var (
	genders = []string{
		"Male",
		"Female",
		"Attack Helicopter",
		"Other",
	}
)

type Session struct {
	Step SessionStep

	Name   string
	Age    int
	Gender string
}

func (h *Handler) Init(ctx context.Context) *tgb.Router {
	//bot.Use(tgb.MiddlewareFunc(func(next tgb.Handler) tgb.Handler {
	//	return tgb.HandlerFunc(func(c context.Context, update *tgb.Update) error {
	//		ctx = logging.ContextWithLogger(ctx)
	//		return next.Handle(ctx, update)
	//	})
	//}))

	//h.Router.Message(h.StartHandler.Start, tgb.Command("start"), tgb.ChatType(tg.ChatTypePrivate))

	//h.registerStartHandler()
	h.registerDbWriteHandler()

	sessionManager := session.NewManager(Session{
		Step: SessionStepInit,
	})

	isSessionStep := func(state SessionStep) tgb.Filter {
		return sessionManager.Filter(func(session *Session) bool {
			return session.Step == state
		})
	}

	isDigit := tgb.Regexp(regexp.MustCompile(`^\d+$`))

	h.Router.
		Use(sessionManager).
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			// handle /start command
			sessionManager.Get(ctx).Step = SessionStepName
			return msg.Update.Reply(ctx, msg.Answer("Hi, what is your name?"))
		}, tgb.Command("start")).
		Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
			// handle no command with SessionStepInitial
			return mu.Update.Reply(ctx, mu.Answer("Press /start to fill the form"))
		}, isSessionStep(SessionStepInit)).
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			// handle name input
			session := sessionManager.Get(ctx)

			session.Name = msg.Text
			session.Step = SessionStepAge

			return msg.Update.Reply(ctx, msg.Answer("What is your age?"))
		}, isSessionStep(SessionStepName)).
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			// handle no digit input when state is SessionStepAge
			return msg.Update.Reply(ctx, msg.Answer("Please, send me just number"))
		}, isSessionStep(SessionStepAge), tgb.Not(isDigit)).
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			// handle correct age input
			age, err := strconv.Atoi(msg.Text)
			if err != nil {
				return fmt.Errorf("parse age: %w", err)
			}

			session := sessionManager.Get(ctx)
			session.Age = int(age)
			session.Step = SessionStepGender

			buttonLayout := tg.NewButtonLayout[tg.KeyboardButton](1)
			for _, gender := range genders {
				buttonLayout.Insert(tg.NewKeyboardButton(gender))
			}

			return msg.Update.Reply(ctx, msg.Answer("What is your gender?").ReplyMarkup(
				tg.NewReplyKeyboardMarkup(
					buttonLayout.Keyboard()...,
				).WithResizeKeyboardMarkup(),
			))
		}, isSessionStep(SessionStepAge), isDigit).
		Message(func(ctx context.Context, mu *tgb.MessageUpdate) error {
			// handle gender input and display results
			session := sessionManager.Get(ctx)

			session.Gender = mu.Text

			answer := mu.Answer(tg.HTML.Text(
				tg.HTML.Line(tg.HTML.Underline(tg.HTML.Text("Your profile:"))),
				tg.HTML.Line(tg.HTML.Bold("├ Your name:"), tg.HTML.Code(session.Name)),
				tg.HTML.Line(tg.HTML.Bold("├ Your age:"), tg.HTML.Code(strconv.Itoa(session.Age))),
				tg.HTML.Line(tg.HTML.Bold("└ Your gender:"), tg.HTML.Code(session.Gender)),
				"",
				tg.HTML.Line(tg.HTML.Italic("press /start to fill again")),
			)).ReplyMarkup(tg.NewReplyKeyboardRemove()).ParseMode(tg.HTML)

			sessionManager.Reset(session)

			return mu.Update.Reply(ctx, answer)
		}, isSessionStep(SessionStepGender), tgb.TextIn(genders)).
		Message(func(ctx context.Context, msg *tgb.MessageUpdate) error {
			return msg.Update.Reply(ctx, msg.Answer("Please, choose one of the buttons below 👇"))
		}, isSessionStep(SessionStepGender), tgb.Not(tgb.TextIn(genders)))

	return h.Router
}
