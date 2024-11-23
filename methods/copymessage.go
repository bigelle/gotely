package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type CopyMessage[T string | int] struct {
	ChatId                T
	FromChatId            T
	MessageId             int
	MessageThreadId       *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	AllowPaidBroadcast    *bool
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyKeyboard
	DisableNotification   *bool
	ProtectContent        *bool
}

func (c CopyMessage[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c.MessageId < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessage[T]) MarshalJSON() ([]byte, error) {
	type alias CopyMessage[T]
	return json.Marshal(alias(c))
}

func (c CopyMessage[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessage", c)
}

type CopyMessages[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageIds          []int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
	RemoveCaption       *bool
}

func (c CopyMessages[T]) MarshalJSON() ([]byte, error) {
	type alias CopyMessages[T]
	return json.Marshal(alias(c))
}

func (c CopyMessages[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(c.MessageIds) < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessages[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessages", c)
}
