package objects

import (
	"fmt"

	"github.com/bigelle/utils.go/ensure"
)

type LoginUrl struct {
	Url               string
	ForwardText       string
	BotUsername       string
	RequestWriteAcess bool
}

func (l LoginUrl) Validate() error {
	if !ensure.NotEmpty(l.Url) {
		return fmt.Errorf("url can't be empty")
	}
	return nil
}
