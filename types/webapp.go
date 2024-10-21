package types

import (
	"fmt"

	"github.com/bigelle/utils.go/ensure"
)

type WebAppInfo struct {
	Url string `json:"url"`
}

func (w WebAppInfo) Validate() error {
	if !ensure.NotEmpty(w.Url) {
		return fmt.Errorf("url can't be empty")
	}
	return nil
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type SentWebAppMessage struct {
	InlineMessageId *string `json:"inline_message_id,omitempty"`
}
