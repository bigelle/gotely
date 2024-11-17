package telego

import (
	"errors"
	"fmt"
	"slices"

	"github.com/bigelle/tele.go/internal/assertions"
	"github.com/bigelle/tele.go/types"
)

type Bot struct {
	// API access token
	Token string
	// a function to process incoming updates
	OnUpdate func(types.Update) error
	// Optional: enable or disable logger
	// Enabled by default
	enableLogger bool
	// Optional: a logger for printing information about any
	// incoming errors, warnings, info, etc.
	// if no logger provided, will be used default one
	logger ILogger
	//Optional: a minimal logging level.
	//Anything below than this level will be ignored.
	//Allowed: "INFO", "WARN", "ERROR", "FATAL"
	level string
}

func (b Bot) Validate() error {
	if err := assertions.ParamNotEmpty(b.Token, "Token"); err != nil {
		return err
	}
	if b.OnUpdate == nil {
		return errors.New("function OnUpdate can't be nil")
	}
	if !slices.Contains([]string{"INFO", "WARN", "ERROR", "FATAL"}, b.level) {
		return errors.New("logging level should be INFO, WARN, ERROR or FATAL")
	}
	return nil
}

type ILogger interface {
	Info(...fmt.Stringer)
	Warn(...fmt.Stringer)
	Error(...fmt.Stringer)
	Fatal(...fmt.Stringer)
}

type defaultLoggerImpl struct {
	// FIXME: there's nothing currently
}

func (d defaultLoggerImpl) Info(s ...fmt.Stringer) {
	fmt.Println(s)
}

func (d defaultLoggerImpl) Warn(s ...fmt.Stringer) {
	fmt.Println(s)
}

func (d defaultLoggerImpl) Error(s ...fmt.Stringer) {
	fmt.Println(s)
}

func (d defaultLoggerImpl) Fatal(s ...fmt.Stringer) {
	fmt.Println(s)
}

type BotOption func(*Bot)

var bot Bot

func NewBot(t string, u func(types.Update) error, opts ...BotOption) (Bot, error) {
	b := Bot{
		Token:        t,
		OnUpdate:     u,
		enableLogger: true,
		logger:       defaultLoggerImpl{},
		level:        "INFO",
	}
	for _, opt := range opts {
		opt(&b)
	}
	if err := b.Validate(); err != nil {
		return Bot{}, err
	}
	bot = b
	return b, nil
}

func GetToken() string {
	return bot.Token
}

func WithEnableLogger(b bool) BotOption {
	return func(bot *Bot) {
		bot.enableLogger = b
	}
}

func WithLogger(l ILogger) BotOption {
	return func(b *Bot) {
		b.logger = l
	}
}

func WithLoggingLevel(lvl string) BotOption {
	return func(b *Bot) {
		b.level = lvl
	}
}
