package longpolling

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
	"github.com/bigelle/gotely/bot"
)

type LongPollingBot struct {
	Bot bot.Bot

	// for getting updates
	offset         *int
	limit          int
	timeout        int
	allowedUpdates *[]string

	// service
	chUpdate    chan objects.Update
	ctx         context.Context
	cancel      context.CancelFunc
	workingPool int
	// TODO maybe logger
}

func (l *LongPollingBot) Start() {
	if err := l.Validate(); err != nil {
		log.Fatal(err)
	}

	l.chUpdate = make(chan objects.Update)

	l.ctx, l.cancel = context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		l.poll()
	}()
	wg.Add(l.workingPool)
	for range l.workingPool {
		go func() {
			defer wg.Done()
			l.answer()
		}()
	}
	log.Println("bot is online")
	wg.Wait()
}

func (l LongPollingBot) Stop() {
	if l.cancel != nil {
		l.cancel()
	}
	close(l.chUpdate)
}

func (l LongPollingBot) Validate() error {
	if l.Bot.Token() == "" {
		return fmt.Errorf("API token can't be empty")
	}
	if l.limit < 1 || l.limit > 100 {
		return fmt.Errorf("limit must be between 1 and 100")
	}
	if l.timeout < 0 {
		return fmt.Errorf("timeout must be positive")
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
				return fmt.Errorf("unknown update type: %s", upd)
			}
		}
	}
	return nil
}

func New(bot bot.Bot, opts ...Option) LongPollingBot {
	lpb := LongPollingBot{
		Bot:            bot,
		limit:          100,
		timeout:        30,
		allowedUpdates: nil,

		workingPool: 1,
	}
	for _, opt := range opts {
		opt(&lpb)
	}
	return lpb
}

type Option func(*LongPollingBot)

// TODO opts

func (l LongPollingBot) poll() {
	for {
		select {
		case <-l.ctx.Done():
			log.Println("exiting polling loop")
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
				log.Printf("error: %s", err.Error())
				continue
			}

			if len(upds) > 0 {
				for _, upd := range upds {
					select {
					case l.chUpdate <- upd:
						log.Println("new incoming update; id = ", upd.UpdateId)
						offset := upd.UpdateId + 1
						l.offset = &offset
					case <-l.ctx.Done():
						log.Println("exiting polling loop")
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
			log.Println("exiting answering loop")
			return

		case upd := <-l.chUpdate:
			err := l.Bot.OnUpdate(upd)
			if err != nil {
				log.Printf("error: %s", err.Error())
				continue
			}
		}
	}
}
