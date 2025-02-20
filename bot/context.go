package bot

import (
	"net/http"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/methods"
	"github.com/bigelle/gotely/api/objects"
)

// Context is a structured set of settings that the bot currently uses to send requests and responses
// that will be used to work with this update
type Context struct {
	//Telegram Bot API token
	Token string
	//Update that bot is currently responding to
	Update objects.Update
	//Client that bot is used
	Client *http.Client
	//Telegram Bot API URL to which the response will be sent
	ApiUrl string
}

func (c *Context) SendMessage(m methods.SendMessage) (*objects.Message, error) {
	return gotely.SendPostRequestWith[objects.Message](
		m,
		c.Token,
		gotely.WithConfig(gotely.RequestConfig{
			Client:         c.Client,
			RequestBaseUrl: c.ApiUrl,
		}),
	)
}

func (c *Context) SendPhoto(m methods.SendPhoto) (*objects.Message, error) {
	return gotely.SendPostRequestWith[objects.Message](
		&m,
		c.Token,
		gotely.WithConfig(gotely.RequestConfig{
			Client:         c.Client,
			RequestBaseUrl: c.ApiUrl,
		}),
	)
}
