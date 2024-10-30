package types

import (
	"github.com/bigelle/tele.go/internal/assertions"
)

type WebAppInfo struct {
	Url string `json:"url"`
}

func (w WebAppInfo) Validate() error {
	if err := assertions.ParamNotEmpty(w.Url, "Url"); err != nil {
		return err
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
