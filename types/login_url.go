package types

import (
	"github.com/bigelle/tele.go/internal/assertions"
)

type LoginUrl struct {
	Url               string  `json:"url"`
	ForwardText       *string `json:"forward_text,omitempty"`
	BotUsername       *string `json:"bot_username,omitempty"`
	RequestWriteAcess *bool   `json:"request_write_acess,omitempty"`
}

// TODO: maybe validating the hash of the received data
func (l LoginUrl) Validate() error {
	if err := assertions.ParamNotEmpty(l.Url, "Url"); err != nil {
		return err
	}
	return nil
}
