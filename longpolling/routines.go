package longpolling

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

func pollUpdates() {
	gu := GetUpdates{
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

type GetUpdates struct {
	Offset         *int      `json:"offset,omitempty"`
	Limit          *int      `json:"limit,omitempty"`
	Timeout        *int      `json:"timeout,omitempty"`
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g GetUpdates) MarshalJSON() ([]byte, error) {
	type alias GetUpdates
	return json.Marshal(alias(g))
}

func (g GetUpdates) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return errors.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	if g.Timeout != nil {
		if *g.Timeout < 0 {
			return errors.ErrInvalidParam("timeout parameter must be positive")
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
				return errors.ErrInvalidParam(fmt.Sprintf("invalid param: %s. allowed parameters: %s", p, strings.Join(allowedUpdates, ", ")))
			}
		}
	}
	return nil
}

func (g GetUpdates) Execute() (*[]types.Update, error) {
	return internal.MakeGetRequest[[]types.Update](telego.GetToken(), "getUpdates", g)
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
