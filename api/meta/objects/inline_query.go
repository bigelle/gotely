package objects

import (
	"encoding/json"
	"fmt"

	"github.com/bigelle/utils.go/ensure"

	"github.com/bigelle/tele.go/api/meta/interfaces"
)

type InlineQuery struct {
	Id       string
	From     User
	Location Location
	Query    string
	Offset   string
	ChatType string
}

type ChosenInlineQuery struct {
	ResultId        string
	From            User
	Location        Location
	InlineMessageId string
	Query           string
}

var ValidThumbTypes = []string{"image/jpeg", "image/gif", "video/mp4"}

type InlineQueryResult interface {
	interfaces.Validable
	interfaces.BotApiObject
}

type InlineQueryResultArticle struct {
	Type                string
	Id                  string
	Title               string
	InputMessageContent *InputMessageContent
	ReplyMarkup         *InlineKeyboardMarkup
	Url                 string
	HideUrl             bool
	Description         string
	ThumbnailUrl        string
	ThumbnailWidth      int
	ThumbnailHeight     int
}

func (i InlineQueryResultArticle) MarshalJSON() ([]byte, error) {
	type Alias InlineQueryResultArticle
	return json.Marshal(&struct {
		Type string
		*Alias
	}{
		Type:  "article",
		Alias: (*Alias)(&i),
	})
}

func (i InlineQueryResultArticle) Validate() error {
	if !ensure.NotEmpty(i.Id) {
		return fmt.Errorf("id parameter can't be empty")
	}
	if !ensure.NotEmpty(i.Title) {
		return fmt.Errorf("title field can't be empty")
	}
	if err := i.InputMessageContent.Validate(); err != nil {
		return err
	}
	if i.ReplyMarkup != nil {
		if err := i.InputMessageContent.Validate(); err != nil {
			return err
		}
	}

	return nil
}
