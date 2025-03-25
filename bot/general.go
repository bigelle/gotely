package bot

import (
	"net/http"

	"github.com/bigelle/gotely/api/objects"
)

type Bot interface {
	Token() string
	ApiUrl() string
	Client() *http.Client
	OnUpdate(objects.Update) error
}
