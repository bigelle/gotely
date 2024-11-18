package telego

import (
	"errors"
	"io"
	"os"
	"sync"

	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/types"
)

type Bot struct {
	// API access token
	Token string
	// a function to process incoming updates
	OnUpdate func(types.Update) error
	// used for displayng warnings, errors and useful information.
	// PANICS if something is gone wrong while writing to io.Writer
	// default: stdout
	Logger io.Writer
}

func (b Bot) Validate() error {
	if err := assertions.ParamNotEmpty(b.Token, "Token"); err != nil {
		return err
	}
	if b.OnUpdate == nil {
		return errors.New("function OnUpdate can't be nil")
	}
	if b.Logger == nil {
		return errors.New(
			"logger can't be nil. if you want to disable console messages, consider using io.Discard",
		)
	}
	return nil
}

type BotOption func(*Bot)

var bot Bot
var once *sync.Once

// if bot already exists, returning existing instance and nil error
func NewBot(t string, u func(types.Update) error, opts ...BotOption) (Bot, error) {
	var err error
	once.Do(func() {
		b := Bot{
			Token:    t,
			OnUpdate: u,
			Logger:   os.Stdout,
		}
		for _, opt := range opts {
			opt(&b)
		}
		if e := b.Validate(); e != nil {
			err = e
			return
		}
		bot = b
	})
	return bot, err
}

func WithWriter(w io.Writer) BotOption {
	return func(b *Bot) {
		b.Logger = w
	}
}

func GetToken() string {
	return bot.Token
}
