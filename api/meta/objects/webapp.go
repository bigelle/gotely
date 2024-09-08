package objects

import (
	"fmt"

	"github.com/bigelle/utils.go/ensure"
)

type WebAppInfo struct {
	Url string
}

func (w WebAppInfo) Validate() error {
	if !ensure.NotEmpty(w.Url) {
		return fmt.Errorf("url can't be empty")
	}
	return nil
}

type WebAppData struct{
	Data string
	ButtonText string
}

type SentWebAppMessage struct{
	InlineMessageId string
}
