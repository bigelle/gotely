package tgbot

import (
	"net/http"

	"github.com/bigelle/gotely/api/objects"
)

// Bot is an interface which is used to create and configure
// both LongPolling and Webhook bots.
type Bot interface {
	// Returns the Telegram Bot API token.
	Token() string
	// Returns the Telegram Bot API URL template in the format "https://api.telegram.org/bot<token>/<method>".
	// Can be overridden to use a custom API URL, for example, when running a local Bot API instance.
	// If unsure, just return [gotely.DEFAULT_URL].
	ApiUrl() string
	// Returns a *[http.Client] used for making API requests.
	Client() *http.Client
	// A function called on every incoming update from the Telegram Bot API.
	OnUpdate(objects.Update) error
}
