package longpolling

import (
	"encoding/json"
	"fmt"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

func pollUpdates() {
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
				longPollingBotInstance.writer.Write([]byte(err.Error()))
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
