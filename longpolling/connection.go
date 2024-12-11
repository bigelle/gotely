package longpolling

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type LongPollingOption func(*LongPollingBot)

var (
	default_offset          = 0
	default_limit           = 100
	default_timeout         = 30
	default_allowed_updates = []string{}
)

func Connect(b telego.Bot, opts ...LongPollingOption) error {
	// creating an instance
	lpb := LongPollingBot{
		OnUpdate:       b.OnUpdate,
		Offset:         &default_offset,
		Limit:          &default_limit,
		Timeout:        &default_timeout,
		AllowedUpdates: &default_allowed_updates,
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

	var respObj internal.ApiResponse[types.User]
	if err := json.Unmarshal(b, &respObj); err != nil {
		return types.User{}, err
	}

	return respObj.Result, nil
}
