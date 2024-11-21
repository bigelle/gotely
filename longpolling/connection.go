package longpolling

import (
	"fmt"
	"sync"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/types"
)

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
		offset:         &default_offset,
		limit:          &default_limit,
		timeout:        &default_timeout,
		allowedUpdates: nil,
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
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		pollUpdates()
	}()
	go func() {
		defer longPollingBotInstance.waitgroup.Done()
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
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
