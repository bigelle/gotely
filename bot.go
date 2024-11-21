package telego

import (
	"io"
	"os"

	"github.com/bigelle/tele.go/types"
)

type Bot struct {
	Token    string
	OnUpdate func(types.Update) error
	Writer   io.StringWriter
}

var bot Bot

func NewBot(t string, onupd func(types.Update) error, opts ...BotOption) (Bot, error) {
	b := Bot{
		Token:    t,
		OnUpdate: onupd,
		Writer:   os.Stdout,
	}
	for _, opt := range opts {
		opt(&b)
	}
	bot = b
	return bot, nil
}

type BotOption func(*Bot)

func WithWriter(w io.StringWriter) BotOption {
	return func(b *Bot) {
		b.Writer = w
	}
}

func GetToken() string {
	return bot.Token
}
