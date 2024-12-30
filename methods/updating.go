package methods

import (
	"encoding/json"
	"strings"

	"github.com/bigelle/tele.go/objects"
)

// MessageOrBool represents the result of "edit-" methods in the Telegram Bot API.
// These methods can return either a Message or a boolean value depending on
// whether the method was used to edit a regular message or an inline message.
//
// - If the method edits a regular message, a Message object is returned.
//
// - If the method edits an inline message, a boolean value is returned to indicate success.
//
// This structure encapsulates both possible return objects for easier handling in Go.
type MessageOrBool struct {
	Message *objects.Message
	Bool    *bool
}

type EditMessageText[T int | string] struct {
	ChatId               *T
	Text                 string
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ParseMode            *string
	Entities             *[]objects.MessageEntity
	LinkPreviewOptions   *objects.LinkPreviewOptions
	ReplyMarkup          *objects.InlineKeyboardMarkup
}

func (e EditMessageText[T]) Validate() error {
	if len(e.Text) < 1 || len(e.Text) > 4096 {
		return objects.ErrInvalidParam("text parameter must be between 1 and 4096 characters")
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.Entities != nil && e.ParseMode != nil {
		return objects.ErrInvalidParam("entities can't be used if parse_mode is provided")
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e EditMessageText[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageText[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("editMessageText", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageText", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageCaption[T int | string] struct {
	ChatId                *T
	Caption               *string
	BusinessConnectionId  *string
	MessageId             *int
	InlineMessageId       *string
	ParseMode             *string
	Entities              *[]objects.MessageEntity
	ShowCaptionAboveMedia *bool
	ReplyMarkup           *objects.InlineKeyboardMarkup
}

func (e EditMessageCaption[T]) Validate() error {
	if e.Caption != nil {
		if len(*e.Caption) < 1 || len(*e.Caption) > 4096 {
			return objects.ErrInvalidParam("caption parameter must be between 1 and 4096 characters")
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.Entities != nil && e.ParseMode != nil {
		return objects.ErrInvalidParam("entities can't be used if parse_mode is provided")
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e EditMessageCaption[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageCaption[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("editMessageCaption", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageCaption", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageMedia[T int | string] struct {
	Media                 objects.InputMedia
	ChatId                *T
	BusinessConnectionId  *string
	MessageId             *int
	InlineMessageId       *string
	ShowCaptionAboveMedia *bool
	ReplyMarkup           *objects.InlineKeyboardMarkup
}

func (e EditMessageMedia[T]) Validate() error {
	if err := e.Media.Validate(); err != nil {
		return err
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e EditMessageMedia[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageMedia[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("editMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageMedia", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageLiveLocation[T int | string] struct {
	Latitude             *float64
	Longtitude           *float64
	LivePeriod           *int
	HorizontalAccuracy   *float64
	Heading              *int
	ProximityAlertRadius *int
	ChatId               *T
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
}

func (e EditMessageLiveLocation[T]) Validate() error {
	if e.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if e.Longtitude == nil {
		return objects.ErrInvalidParam("longtitude parameter can't be empty")
	}
	if e.HorizontalAccuracy != nil {
		if *e.HorizontalAccuracy < 0 || *e.HorizontalAccuracy > 1500 {
			return objects.ErrInvalidParam("horizontal_accuracy parameter must be between 0 and 1500 meetrs")
		}
	}
	if e.Heading != nil {
		if *e.Heading < 1 || *e.Heading > 360 {
			return objects.ErrInvalidParam("heading parameter must be between 1 and 360 degrees")
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e EditMessageLiveLocation[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageLiveLocation[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("editMessageLiveLocation", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageLiveLocation", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type StopMessageLiveLocation[T int | string] struct {
	ChatId               *T
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
}

func (e StopMessageLiveLocation[T]) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e StopMessageLiveLocation[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e StopMessageLiveLocation[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("stopMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageMedia", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageReplyMarkup[T int | string] struct {
	ChatId               *T
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
}

func (e EditMessageReplyMarkup[T]) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return objects.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return objects.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (e EditMessageReplyMarkup[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageReplyMarkup[T]) Execute() (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type StopPoll[T int | string] struct {
	ChatId               T
	MessageId            int
	BusinessConnectionId *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
}

func (s StopPoll[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	if s.ReplyMarkup != nil {
		if err := s.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s StopPoll[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s StopPoll[T]) Execute() (*objects.Poll, error) {
	return MakePostRequest[objects.Poll]("stopPoll", s)
}

type DeleteMessage[T int | string] struct {
	ChatId    T
	MessageId int
}

func (d DeleteMessage[T]) Validate() error {
	if c, ok := any(d.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(d.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if d.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (d DeleteMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessage[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteMessage", d)
}

type DeleteMessages[T int | string] struct {
	ChatId     T
	MessageIds []int
}

func (d DeleteMessages[T]) Validate() error {
	if c, ok := any(d.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(d.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if len(d.MessageIds) < 1 || len(d.MessageIds) > 100 {
		return objects.ErrInvalidParam("message_ids parameter must be between 1 and 100")
	}

	return nil
}

func (d DeleteMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessages[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteMessages", d)
}
