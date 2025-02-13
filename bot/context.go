package bot

import (
	"net/http"

	"github.com/bigelle/gotely/api/objects"
)

// Context is a structured set of settings that the bot currently uses to send requests and responses
// that will be used to work with this update
type Context struct {
	//Update that bot is currently responding to
	Update objects.Update
	//Client that bot is used
	Client *http.Client
	//Telegram Bot API URL to which the response will be sent
	ApiUrl string
}
