package telego

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"
	"sync"

	"github.com/bigelle/tele.go/types"
)

type LongPollingOption func(*LongPollingBot)

var (
	default_longpolling_offset          = 0
	default_longpolling_limit           = 100
	default_longpolling_timeout         = 30
	default_longpolling_allowed_updates = []string{}
)

func Connect(b Bot, opts ...LongPollingOption) error {
	// creating an instance
	lpb := LongPollingBot{
		OnUpdate:       b.OnUpdate,
		Offset:         &default_longpolling_offset,
		Limit:          &default_longpolling_limit,
		Timeout:        &default_longpolling_timeout,
		AllowedUpdates: &default_longpolling_allowed_updates,
		updates:        make(chan types.Update),
		stopChan:       make(chan struct{}),
		waitgroup:      &sync.WaitGroup{},
		writer:         b.Writer,
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
	longPollingBotInstance.writer.WriteString("INFO: bot is now online!\n")
	longPollingBotInstance.waitgroup.Add(2)
	go func() {
		defer longPollingBotInstance.waitgroup.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic while polling updates: %v\n", r)
			}
		}()
		pollUpdates()
	}()
	go func() {
		defer longPollingBotInstance.waitgroup.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic while handling updates: %v\n", r)
			}
		}()
		handleUpdates()
	}()
	longPollingBotInstance.waitgroup.Wait()

	return nil
}

func Disconnect() {
	longPollingBotInstance.stopChan <- struct{}{}
	close(longPollingBotInstance.stopChan)
}

func getMe() (*types.User, error) {
	token := GetToken()
	if token == "" {
		return nil, types.ErrInvalidParam("token can't be empty. did you connect the bot?")
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/getMe", token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respObj apiResponse[types.User]
	if err := json.Unmarshal(b, &respObj); err != nil {
		return nil, err
	}

	return &respObj.Result, nil
}

type apiResponse[T any] struct {
	Ok          bool                      `json:"ok"`
	ErrorCode   int                       `json:"error_code"`
	Description *string                   `json:"description,omitempty"`
	Parameters  *types.ResponseParameters `json:"parameters,omitempty"`
	Result      T                         `json:"result"`
}

type ErrNoUpdates struct {
	Cause   error
	Message *string
	Code    *int
}

func (e ErrNoUpdates) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("failed to get new updates: %s", e.Cause.Error())
	}
	if e.Code != nil && e.Message != nil {
		return fmt.Sprintf("failed to get new updates: Telegram server has responded with code %d and message `%s`", *e.Code, *e.Message)
	}
	return fmt.Sprintf("failed to get new updates")
}

func pollUpdates() {
	gu := getUpdates{
		AllowedUpdates: longPollingBotInstance.AllowedUpdates,
		Limit:          longPollingBotInstance.Limit,
		Timeout:        longPollingBotInstance.Timeout,
		Offset:         longPollingBotInstance.Offset,
	}
	for {
		select {
		case <-longPollingBotInstance.stopChan:
			return
		default:
			upds, err := gu.Execute()
			if err != nil {
				longPollingBotInstance.writer.WriteString(err.Error())
				continue
			}
			updsUnpacked := *upds
			if len(*upds) > 0 {
				for _, upd := range updsUnpacked {
					longPollingBotInstance.updates <- upd
				}
				lastUpdate := updsUnpacked[len(updsUnpacked)-1]
				*longPollingBotInstance.Offset = lastUpdate.UpdateId + 1
			}
		}
	}
}

type getUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g getUpdates) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return types.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	if g.Timeout != nil {
		if *g.Timeout < 0 {
			return types.ErrInvalidParam("timeout parameter must be positive")
		}
	}
	if g.AllowedUpdates != nil && len(*g.AllowedUpdates) > 0 {
		allowedUpdates := []string{
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
		for _, p := range *g.AllowedUpdates {
			if !slices.Contains(allowedUpdates, p) {
				return types.ErrInvalidParam(fmt.Sprintf("invalid param: %s. allowed parameters: %s", p, strings.Join(allowedUpdates, ", ")))
			}
		}
	}
	return nil
}

func (g getUpdates) Execute() (*[]types.Update, error) {
	token := GetToken()
	if token == "" {
		return nil, ErrNoUpdates{
			Cause: fmt.Errorf("token can't be empty. did you connect the bot?"),
		}
	}

	if err := g.Validate(); err != nil {
		return nil, ErrNoUpdates{
			Cause: fmt.Errorf("can't make request: %w", err),
		}
	}

	data, err := json.Marshal(g)
	if err != nil {
		return nil, ErrNoUpdates{
			Cause: fmt.Errorf("can't marshal request body: %w", err),
		}
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, "getUpdates")

	req, err := http.NewRequest("GET", url, bytes.NewReader(data))
	if err != nil {
		return nil, ErrNoUpdates{
			Cause: fmt.Errorf("can't form a request: %w", err),
		}
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		msg := err.Error()
		return nil, ErrNoUpdates{
			Code:    &resp.StatusCode,
			Message: &msg,
		}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't read response: %w", err)
	}

	var apiResp apiResponse[[]types.Update]
	if err := json.Unmarshal(b, &apiResp); err != nil {
		return nil, fmt.Errorf("can't unmarshal response body: %w", err)
	}
	if !apiResp.Ok {
		return nil, ErrNoUpdates{
			Code:    &apiResp.ErrorCode,
			Message: apiResp.Description,
		}
	}
	return &apiResp.Result, nil
}

func handleUpdates() {
	for {
		select {
		case upd := <-longPollingBotInstance.updates:
			err := longPollingBotInstance.OnUpdate(upd)
			if err != nil {
				// logging and panic if an error is critical
			}
		case <-longPollingBotInstance.stopChan:
			return
		}
	}

}

var longPollingBotInstance LongPollingBot

type LongPollingBot struct {
	OnUpdate func(types.Update) error
	// Optional: Identifier of the first update to ben returned.
	// Must be greater by one than the highest among the identifiers of
	// previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned. An update is
	// Considered confirmed as soon as getUpdates is called
	// with an offset higher than its update_id.
	// The negative offset can be specified to retrieve
	// updates starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	Offset *int
	// Optional: Limits the number of updates to be retrieved.
	// Values between 1-100 are accepted. Defaults to 100.
	Limit *int
	// Optional: Timeout in seconds for long polling.
	//Defaults to 30. Should be positive,
	//short polling should be used for testing purposes only.
	Timeout *int
	// Optional: list of the update types you want your bot to
	// receive. For example, specify "message",
	// "edited_channel_post", "callback_query" to only receive
	// updates of these types.
	// Specify an empty list to receive all update types
	// except chat_member, message_reaction, and
	// message_reaction_count (default). If not specified, the
	// previous setting will be used.
	AllowedUpdates *[]string
	// a channel that will store all of the updates that are
	// waiting to be processed
	updates chan types.Update
	// a context channel for stopping bot
	stopChan chan struct{}
	// a waiting group to sync getUpdates and handleUpdates
	waitgroup *sync.WaitGroup
	// used to displaying messages about any warnings, errors, etc
	writer io.StringWriter
}

func (l LongPollingBot) Validate() error {
	if l.OnUpdate == nil {
		return errors.New("function OnUpdate can't be nil")
	}
	if *l.Limit < 1 || *l.Limit > 100 {
		return errors.New("limit parameter should be between 1 and 100")
	}
	if *l.Timeout < 0 {
		return errors.New("timeout parameter should be positive")
	}
	if l.AllowedUpdates != nil && len(*l.AllowedUpdates) > 0 {
		allowedUpdates := []string{
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
		for _, p := range *l.AllowedUpdates {
			if !slices.Contains(allowedUpdates, p) {
				return types.ErrInvalidParam(fmt.Sprintf("invalid param: %s. allowed parameters: %s", p, strings.Join(allowedUpdates, ", ")))
			}
		}
	}
	return nil
}
