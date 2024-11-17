package longpolling

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

var longPollingBotInstance LongPollingBot

type LongPollingBot struct {
	OnUpdate func(types.Update) error
	// a channel that will store all of the updates that are
	// waiting to be processed
	updates chan types.Update
	// a context channel for stopping bot
	stopChan chan struct{}
	// a waiting group to sync getUpdates and handleUpdates
	waitgroup *sync.WaitGroup
	// Optional: Identifier of the first update to ben returned.
	// Must be greater by one than the highest among the identifiers of
	// previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned. An update is
	// Considered confirmed as soon as getUpdates is called
	// with an offset higher than its update_id.
	// The negative offset can be specified to retrieve
	// updates starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	offset *int
	// Optional: Limits the number of updates to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	limit *int
	// Optional: Timeout in seconds for long polling.
	//Defaults to 30. Should be positive,
	//short polling should be used for testing purposes only.
	timeout *int
	// Optional: list of the update types you want your bot to
	// receive. For example, specify "message",
	// "edited_channel_post", "callback_query" to only receive
	// updates of these types.
	// Specify an empty list to receive all update types
	// except chat_member, message_reaction, and
	// message_reaction_count (default). If not specified, the
	// previous setting will be used.
	allowedUpdates *[]string
}

func (l LongPollingBot) Validate() error {
	if l.OnUpdate == nil {
		return errors.New("function OnUpdate can't be nil")
	}
	if *l.limit < 1 || *l.limit > 100 {
		return errors.New("limit parameter should be between 1 and 100")
	}
	if *l.timeout < 0 {
		return errors.New("timeout parameter should be positive")
	}
	allowed := []string{
		"message",
		"edited_message",
		"channel_post",
		"edited_channel_post",
		"inline_query",
		"chosen_inline_result",
		"callback_query",
		"shipping_query",
		"pre_checkout_query",
		"poll",
		"poll_answer",
		"my_chat_member",
		"chat_member",
		"chat_join_request",
	}
	containsAll := func(slice1, slice2 []string) bool {
		elements := make(map[string]struct{})
		for _, item := range slice2 {
			elements[item] = struct{}{}
		}
		for _, item := range slice1 {
			if _, exists := elements[item]; !exists {
				return false
			}
		}
		return true
	}
	if l.allowedUpdates != nil {
		if !containsAll(*l.allowedUpdates, allowed) {
			return fmt.Errorf("unknown allowed_updates parameter: %v", *l.allowedUpdates)
		}
	}
	return nil
}

type LongPollingOption func(*LongPollingBot)

var (
	default_offset  = 0
	default_limit   = 100
	default_timeout = 30
)

func Connect(b telego.Bot, opts ...LongPollingOption) error {
	// creating an instance
	lpb := LongPollingBot{
		OnUpdate:       b.OnUpdate,
		updates:        make(chan types.Update),
		stopChan:       make(chan struct{}),
		waitgroup:      &sync.WaitGroup{},
		offset:         &default_offset,
		limit:          &default_limit,
		timeout:        &default_timeout,
		allowedUpdates: nil,
	}
	for _, opt := range opts {
		opt(&lpb)
	}

	// validation
	if err := lpb.Validate(); err != nil {
		return err
	}
	if _, err := getMe(); err != nil {
		return err
	}
	longPollingBotInstance = lpb

	// launching goroutines
	longPollingBotInstance.waitgroup.Add(2)
	go getUpdates()
	go handleUpdates()
	return nil
}

func Disconnect() {
	longPollingBotInstance.stopChan <- struct{}{}
	close(longPollingBotInstance.stopChan)
	longPollingBotInstance.waitgroup.Wait()
}

func getMe() (types.User, error) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", telego.GetToken())
	resp, err := http.Get(url)
	if err != nil {
		return types.User{}, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.User{}, err
	}

	var respObj types.ApiResponse[types.User]
	if err := json.Unmarshal(b, &respObj); err != nil {
		return types.User{}, err
	}

	return respObj.Result, nil
}

func getUpdates() {
	gu := GetUpdates{
		AllowedUpdates: longPollingBotInstance.allowedUpdates,
		Limit:          longPollingBotInstance.limit,
		Timeout:        longPollingBotInstance.timeout,
		Offset:         longPollingBotInstance.offset,
	}
	for {
		select {
		case <-longPollingBotInstance.stopChan:
			return
		default:
			upds, err := gu.Execute()
			if err != nil {
				// logging error
				// if error is critical, panic
				// TODO: error types
				continue
			}
			for _, upd := range upds {
				longPollingBotInstance.updates <- upd
			}
		}
	}
}

type GetUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g *GetUpdates) SetOffset(o int) *GetUpdates {
	g.Offset = &o
	return g
}

func (g *GetUpdates) SetLimit(l int) *GetUpdates {
	g.Limit = &l
	return g
}

func (g *GetUpdates) SetTimeout(t int) *GetUpdates {
	g.Timeout = &t
	return g
}

func (g *GetUpdates) SetAllowedUpdates(s []string) *GetUpdates {
	g.AllowedUpdates = &s
	return g
}

func (g GetUpdates) Execute() ([]types.Update, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	b, err := internal.MakeGetRequest(telego.GetToken(), "getUpdates", data)
	if err != nil {
		return nil, err
	}

	var resp types.ApiResponse[[]types.Update]
	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d: %s", resp.ErrorCode, *resp.Description)
	}
	return resp.Result, nil
}

func handleUpdates() {
	for upd := range longPollingBotInstance.updates {
		err := longPollingBotInstance.OnUpdate(upd)
		if err != nil {
			// logging and panic if an error is critical
		}
	}
}
