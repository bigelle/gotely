package longpolling

import (
	"sync"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/types"
)

type LongPollingOption func(*longPollingBot)

var (
	default_offset  = 0
	default_limit   = 100
	default_timeout = 30
)

func Connect(b telego.Bot, opts ...LongPollingOption) error {
	// creating an instance
	lpb := longPollingBot{
		OnUpdate:       b.OnUpdate,
		offset:         &default_offset,
		limit:          &default_limit,
		timeout:        &default_timeout,
		allowedUpdates: nil,
		updates:        make(chan types.Update),
		stopChan:       make(chan struct{}),
		waitgroup:      &sync.WaitGroup{},
		writer:         b.Logger,
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
	go pollUpdates()
	go handleUpdates()
	return nil
}

func Disconnect() {
	longPollingBotInstance.stopChan <- struct{}{}
	close(longPollingBotInstance.stopChan)
	longPollingBotInstance.waitgroup.Wait()
}
