package methods

import (
	"encoding/json"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type SendMessage[T int | string] struct {
	ChatId               T                         `json:"chat_id"`
	Text                 string                    `json:"text"`
	BusinessConnectionId *string                   `json:"business_connection_id,omitempty"`
	MessageThreadId      *int                      `json:"message_thread_id,omitempty"`
	ParseMode            *string                   `json:"parse_mode,omitempty"`
	Entities             *[]types.MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  *bool                     `json:"disable_notification,omitempty"`
	ProtectContent       *bool                     `json:"protect_content,omitempty"`
	MessageEffectId      *string                   `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          *types.ReplyKeyboard      `json:"reply_markup,omitempty"`
}

func (s SendMessage[T]) Validate() error {
	if err := assertions.ParamNotEmpty(s.Text, "Text"); err != nil {
		return err
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c == 0 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(string); ok {
		if err := assertions.ParamNotEmpty(c, "ChatId"); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMessage[T]) MarshalJSON() ([]byte, error) {
	type alias SendMessage[T]
	return json.Marshal(alias(s))
}

func (s SendMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendMessage", s)
}
