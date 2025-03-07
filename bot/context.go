package bot

import (
	"net/http"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/methods"
	"github.com/bigelle/gotely/api/objects"
)

// Context is a structured set of settings that the bot currently uses
// to send requests and responses for handling this update.
type Context struct {
	// Telegram Bot API token
	Token string
	// Update that the bot is currently responding to
	Update objects.Update
	// HTTP client used by the bot
	Client *http.Client
	// Telegram Bot API URL to which the response will be sent
	ApiUrl string
}

func (c *Context) SendMessage(m methods.SendMessage) (*objects.Message, error) {
	return gotely.SendPostRequestWith[objects.Message](
		m,
		c.Token,
		gotely.WithClient(c.Client),
		gotely.WithUrl(c.ApiUrl),
	)
}

func (c *Context) SendPhoto(m methods.SendPhoto) (*objects.Message, error) {
	return gotely.SendPostRequestWith[objects.Message](
		&m,
		c.Token,
		gotely.WithClient(c.Client),
		gotely.WithUrl(c.ApiUrl),
	)
}
