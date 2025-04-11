// This package provides an interface for creating [longpolling.LongPollingBot] or [webhook.WebhookBot].
// It is used to store the bot's configuration as well as the logic for handling received updates.
//
// Licensed under the MIT License. See LICENSE file for details.
package tgbot

import (
	"net/http"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
)

// Bot is an interface which is used to create and configure
// both LongPolling and Webhook bots.
type Bot interface {
	// Returns the Telegram Bot API token.
	Token() string
	// Returns the Telegram Bot API URL template in the format "https://api.telegram.org/bot<token>/<method>".
	// Can be overridden to use a custom API URL, for example, when running a local Bot API instance.
	// If unsure, just return [gotely.DEFAULT_URL].
	ApiURLTemplate() string
	// Returns a *[http.Client] used for making API requests.
	Client() *http.Client
	// A function called on every incoming update from the Telegram Bot API.
	OnUpdate(objects.Update) error
}

type DefaultBot struct{}

func (b DefaultBot) Client() *http.Client {
	return http.DefaultClient
}

func (b DefaultBot) ApiURLTemplate() string {
	return gotely.DEFAULT_URL_TEMPLATE
}
