package longpolling

import (
	"fmt"
	"io"

	"github.com/bigelle/gotely"
)

// Use this method to receive incoming updates using long polling
// (https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// Returns an Array of [objects.Update] objects.
type GetUpdates struct {
	// Identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned.
	// An update is considered confirmed as soon as [GetUpdates] is called with an offset higher than its update_id.
	// The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
	// All previous updates will be forgotten.
	Offset *int `json:"offset,omitempty"`
	// Limits the number of updates to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int `json:"limit,omitempty"`
	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling.
	// Should be positive, short polling should be used for testing purposes only.
	Timeout *int `json:"timeout,omitempty"`
	// A JSON-serialized list of the update types you want your bot to receive.
	// For example, specify ["message", "edited_channel_post", "callback_query"] to only receive updates of these types.
	// See [objects.Update] for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member, message_reaction, and message_reaction_count (default).
	// If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to [GetUpdates],
	// so unwanted updates may be received for a short period of time.
	AllowedUpdates *[]string `json:"allowed_updates,omitempty"`
}

func (g GetUpdates) Endpoint() string {
	return "getUpdates"
}

func (g GetUpdates) Validate() error {
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return fmt.Errorf("limit must be between 1 and 100")
		}
	}
	if g.Timeout != nil {
		if *g.Timeout < 0 {
			return fmt.Errorf("timeout must be positive")
		}
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
	if g.AllowedUpdates != nil {
		for _, upd := range *g.AllowedUpdates {
			if _, ok := allowed[upd]; !ok {
				return fmt.Errorf("unknown update type: %s", upd)
			}
		}
	}
	return nil
}

func (g GetUpdates) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g GetUpdates) ContentType() string {
	return "application/json"
}
