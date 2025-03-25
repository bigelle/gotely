package bot

import (
	"net/http"

	"github.com/bigelle/gotely/api/objects"
)

const (
	// Default Telegram Bot API URL template.
	// <token> will be replaced with the bot token, and <method> with the API method name.
	DEFAULT_URL = "https://api.telegram.org/bot<token>/<method>"
)

// Bot is an interface which is used to create and configure
// both LongPolling and Webhook bots.
type Bot interface {
	// Returns the Telegram Bot API token.
	Token() string
	// Returns the Telegram Bot API URL template in the format "https://api.telegram.org/bot<token>/<method>".
	// Can be overridden to use a custom API URL, for example, when running a local Bot API instance.
	// If unsure, just return [DEFAULT_URL].
	ApiUrl() string
	// Returns a *[http.Client] used for making API requests.
	Client() *http.Client
	// A function called on every incoming update from the Telegram Bot API.
	OnUpdate(objects.Update) error
}
