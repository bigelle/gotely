package telego

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/bigelle/tele.go/objects"
)

type ErrBadBot struct {
	BadField string
	Message  string
}

func (e ErrBadBot) Error() string {
	return fmt.Sprintf("an error at %s field: %s", e.BadField, e.Message)
}

var global_token string

func GetToken() string {
	return global_token
}

// LongPolingBot is a struct that is used to set up long-polling bot
type LongPollingBot struct {
	//Telegram bot API access token
	Token string
	// A function which is called with every incoming update
	OnUpdate func(objects.Update) error
	// A function which is called whenever an error occurs
	OnError func(error)
	//Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit int
	//Timeout in seconds for long polling. Should be positive, short polling should be used for testing purposes only.
	//Defaults to 30
	Timeout int
	//A list of the update types you want your bot to receive.
	// See Update for a complete list of available update types.
	//Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	//If not specified, the previous setting will be used.
	AllowedUpdates *[]string
	errChan        chan error
	updChan        chan objects.Update
	ctx            context.Context
	cancel         context.CancelFunc
	offset         *int
}

type LongPollingOption func(*LongPollingBot)

func NewDefaultLongPollingBot(tkn string, onUpd func(objects.Update) error, opts ...LongPollingOption) (*LongPollingBot, error) {
	l := LongPollingBot{
		Token:          tkn,
		OnUpdate:       onUpd,
		OnError:        defaultOnError,
		Limit:          100,
		Timeout:        30,
		AllowedUpdates: nil,
	}
	for _, opt := range opts {
		opt(&l)
	}
	if err := l.Validate(); err != nil {
		return nil, err
	}
	return &l, nil
}

func WithLimit(l int) LongPollingOption {
	return func(lpb *LongPollingBot) {
		lpb.Limit = l
	}
}

func WithAllowedUpdates(au []string) LongPollingOption {
	return func(lpb *LongPollingBot) {
		lpb.AllowedUpdates = &au
	}
}

func WithTimeout(t int) LongPollingOption {
	return func(lpb *LongPollingBot) {
		lpb.Timeout = t
	}
}

func WithOnErrFunc(onE func(error)) LongPollingOption {
	return func(lpb *LongPollingBot) {
		lpb.OnError = onE
	}
}

func defaultOnError(e error) {
	fmt.Println(e.Error())
}

func (l *LongPollingBot) Start() {
	global_token = l.Token

	l.updChan = make(chan objects.Update)
	l.errChan = make(chan error)

	ctx, cancel := context.WithCancel(context.Background())
	l.ctx = ctx
	l.cancel = cancel

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		l.handle_errors()
	}()
	go func() {
		defer wg.Done()
		l.poll()
	}()
	go func() {
		defer wg.Done()
		l.handle_updates()
	}()
	wg.Wait()
}

func (l *LongPollingBot) Stop() {
	l.cancel()
	close(l.errChan)
	close(l.updChan)
}

func (l LongPollingBot) Validate() error {
	if l.Token == "" {
		return ErrBadBot{
			BadField: "Token",
			Message:  "it should not be empty",
		}
	}
	if l.OnUpdate == nil {
		return ErrBadBot{
			BadField: "OnUpdate",
			Message:  "this function is necessary for bot to operate",
		}
	}
	if l.OnError == nil {
		return ErrBadBot{
			BadField: "OnError",
			Message:  "it can't be nil. if you want to disable logging, please provide a function that doesn`t do anything",
		}
	}
	if l.Limit < 1 || l.Limit > 100 {
		return ErrBadBot{
			BadField: "Limit",
			Message:  "it should be between 1 and 100",
		}
	}
	if l.Timeout < 0 {
		return ErrBadBot{
			BadField: "Timeout",
			Message:  "it should be positive",
		}
	}
	if l.AllowedUpdates != nil {
		if len(*l.AllowedUpdates) != 0 {
			valid := map[string]struct{}{
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
			for _, upd := range *l.AllowedUpdates {
				if _, ok := valid[upd]; !ok {
					return ErrBadBot{
						BadField: "AllowedUpdates",
						Message:  fmt.Sprintf("unknown update type `%s`", upd),
					}
				}
			}
		}
	}
	return nil
}

type getUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

type apiResponse[T any] struct {
	Ok          bool
	Description *string
	Result      T
	ErrorCode   *int
	Parameters  *objects.ResponseParameters
}

func (l *LongPollingBot) poll() {
	for {
		select {
		case <-l.ctx.Done():
			return
		default:
			url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates", l.Token)

			payload := getUpdates{
				Offset:         l.offset,
				Limit:          &l.Limit,
				Timeout:        &l.Timeout,
				AllowedUpdates: l.AllowedUpdates,
			}
			b, err := json.Marshal(payload)
			if err != nil {
				l.errChan <- err
				continue
			}

			req, err := http.NewRequestWithContext(l.ctx, "GET", url, bytes.NewReader(b))
			if err != nil {
				l.errChan <- err
				continue
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				l.errChan <- err
				continue
			}

			respb, err := io.ReadAll(resp.Body)
			if err != nil {
				l.errChan <- err
				continue
			}

			var result apiResponse[[]objects.Update]
			if err := json.Unmarshal(respb, &result); err != nil {
				l.errChan <- err
				continue
			}

			if !result.Ok {
				l.errChan <- fmt.Errorf("%s", *result.Description)
				continue
			}

			upds := result.Result
			if len(upds) > 0 {
				for _, upd := range upds {
					l.updChan <- upd
				}
				newoffset := upds[len(upds)-1].UpdateId + 1
				l.offset = &newoffset
			}

		}
	}
}

func (l *LongPollingBot) handle_updates() {
	for {
		select {
		case <-l.ctx.Done():
			return
		case upd := <-l.updChan:
			if l.offset != nil {
				fmt.Printf("Processing update with UpdateId: %d, current offset: %d\n", upd.UpdateId, *l.offset)
			}
			err := l.OnUpdate(upd)
			if err != nil {
				l.errChan <- err
			}
		}
	}
}

func (l *LongPollingBot) handle_errors() {
	for {
		select {
		case <-l.ctx.Done():
			return
		case err := <-l.errChan:
			l.OnError(err)
		}
	}
}
