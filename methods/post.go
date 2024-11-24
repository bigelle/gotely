package methods

import (
	"encoding/json"
	"fmt"
	"strings"

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
	ReplyMarkup          *types.ReplyMarkup        `json:"reply_markup,omitempty"`
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
	ReplyMarkup           *types.ReplyMarkup
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

type SendPhoto[T string | int, B types.InputFile | string] struct {
	ChatId                T
	Photo                 B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
}

func (s SendPhoto[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Photo).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Photo).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendPhoto[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendPhoto[T, B]
	return json.Marshal(alias(s))
}

func (s SendPhoto[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPhoto", s)
}

type SendAudio[T string | int, B types.InputFile | string] struct {
	ChatId               T
	Photo                B
	BusinessConnectionId *string
	MessageThreadId      *int
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]types.MessageEntity
	Duration             *int
	Performer            *string
	Title                *string
	Thumbnail            *B
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendAudio[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Photo).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Photo).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendAudio[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendAudio[T, B]
	return json.Marshal(alias(s))
}

func (s SendAudio[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendAudio", s)
}
