package longpolling

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/tgbot"
)

type LongPollingBot struct {
	Bot tgbot.Bot

	// for getting updates
	offset         *int
	limit          int
	timeout        int
	allowedUpdates *[]string

	// service
	chUpdate    chan objects.Update
	ctx         context.Context
	cancel      context.CancelFunc
	workingPool uint
	logger      slog.Logger
}

func (l *LongPollingBot) Start() {
	l.logger.Info("validating...")
	if err := l.Validate(); err != nil {
		l.logger.Error("bot failed validation;", "err", err.Error())
		os.Exit(1)
	}

	l.logger.Info("initializing...")
	l.chUpdate = make(chan objects.Update)
	l.ctx, l.cancel = context.WithCancel(context.Background())

	l.logger.Info("launching goroutines...")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.poll()
	}()

	l.logger.Info("preparing to work with", "working pool size", l.workingPool)
	wg.Add(int(l.workingPool))
	for range l.workingPool {
		go func() {
			defer wg.Done()
			l.answer()
		}()
	}
	l.logger.Info("bot is online")
	wg.Wait()
}

func (l LongPollingBot) Stop() {
	if l.cancel != nil {
		l.cancel()
	}
	close(l.chUpdate)
	l.logger.Info("bot is offline")
}

func (l LongPollingBot) Validate() error {
	var err gotely.ErrFailedValidation
	if l.Bot.Token() == "" {
		err = append(err, fmt.Errorf("API token can't be empty"))
	}
	if l.limit < 1 || l.limit > 100 {
		err = append(err, fmt.Errorf("limit must be between 1 and 100"))
	}
	if l.timeout < 0 {
		err = append(err, fmt.Errorf("timeout must be positive"))
	}
	allowed := map[string]struct{}{
		"message":                   {},
		"edited_message":            {},
		"channel_post":              {},
		"edited_channel_post":       {},
		"business_connection":       {},
		"business_message":          {},
		"edited_business_message":   {},
		"deleted_business_messages": {},
		"message_reaction":          {},
		"message_reaction_count":    {},
		"inline_query":              {},
		"chosen_inline_result":      {},
		"callback_query":            {},
		"shipping_query":            {},
		"pre_checkout_query":        {},
		"purchased_paid_media":      {},
		"poll":                      {},
		"poll_answer":               {},
		"my_chat_member":            {},
		"chat_member":               {},
		"chat_join_request":         {},
		"chat_boost":                {},
		"removed_chat_boost":        {},
	}
	if l.allowedUpdates != nil {
		for _, upd := range *l.allowedUpdates {
			if _, ok := allowed[upd]; !ok {
				err = append(err, fmt.Errorf("unknown update type: %s", upd))
			}
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func New(bot tgbot.Bot, opts ...Option) LongPollingBot {
	lpb := LongPollingBot{
		Bot: bot,

		limit:          100,
		timeout:        30,
		allowedUpdates: nil,

		workingPool: 1,
		logger:      *slog.Default(),
	}
	for _, opt := range opts {
		opt(&lpb)
	}
	return lpb
}

type Option func(*LongPollingBot)

func WithTimeout(t int) Option {
	return func(lpb *LongPollingBot) {
		lpb.timeout = t
	}
}

func WithLimit(l int) Option {
	return func(lpb *LongPollingBot) {
		lpb.limit = l
	}
}

func WithAllowedUpdates(u *[]string) Option {
	return func(lpb *LongPollingBot) {
		lpb.allowedUpdates = u
	}
}

func WithWorkingPool(p uint) Option {
	return func(lpb *LongPollingBot) {
		if p == 0 {
			lpb.logger.Warn("attempted to set working pool size to", "pool", 0, "falling back to default size", 1)
			lpb.workingPool = 1
		}
		lpb.workingPool = p
	}
}

func (l LongPollingBot) poll() {
	for {
		select {
		case <-l.ctx.Done():
			l.logger.Info("exiting polling loop")
			return

		default:
			g := GetUpdates{
				Offset:         l.offset,
				Limit:          &l.limit,
				Timeout:        &l.timeout,
				AllowedUpdates: l.allowedUpdates,
			}
			var upds []objects.Update
			err := gotely.SendRequestWith(
				g,
				&upds,
				l.Bot.Token(),
				gotely.WithClient(l.Bot.Client()),
				gotely.WithContext(l.ctx),
				gotely.WithUrl(l.Bot.ApiUrl()),
			)
			if err != nil {
				l.logger.Error("error while requesting for new updates;",
					"err", err.Error(),
					"offset", g.Offset,
					"limit", g.Limit,
					"timeout", g.Timeout,
					"allowed_updates", g.AllowedUpdates,
				)
				continue
			}

			if len(upds) > 0 {
				for _, upd := range upds {
					select {
					case l.chUpdate <- upd:
						l.logger.Info("new incoming update;", "update_id", upd.UpdateId)
						offset := upd.UpdateId + 1
						l.offset = &offset
					case <-l.ctx.Done():
						l.logger.Info("exiting polling loop")
						return
					}
				}
			}
		}
	}
}

func (l *LongPollingBot) answer() {
	for {
		select {
		case <-l.ctx.Done():
			l.logger.Info("exiting answering loop")
			return

		case upd := <-l.chUpdate:
			err := l.Bot.OnUpdate(upd)
			if err != nil {
				l.logger.Error("error while answering to an update;", "update_id", upd.UpdateId, "err", err.Error())
				continue
			}
			l.logger.Info("done answering to update", "update_id", upd.UpdateId)
		}
	}
}
