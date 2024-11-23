package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type ForwardMessage[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageId           int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
}

func (f ForwardMessage[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if f.MessageId < 1 {
		return assertions.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (f ForwardMessage[T]) MarshalJSON() ([]byte, error) {
	type alias ForwardMessage[T]
	return json.Marshal(alias(f))
}

func (f ForwardMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessage", f)
}

type ForwardMessages[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageIds          []int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
}

func (f ForwardMessages[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(f.MessageIds) < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (f ForwardMessages[T]) MarshalJSON() ([]byte, error) {
	type alias ForwardMessages[T]
	return json.Marshal(alias(f))
}

func (f ForwardMessages[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessages", f)
}
