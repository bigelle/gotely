package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

// MessageOrBool represents the result of "edit-" methods in the Telegram Bot API.
// These methods can return either a Message or a boolean value depending on
// whether the method was used to edit a regular message or an inline message.
//
// - If the method edits a regular message, a Message object is returned.
//
// - If the method edits an inline message, a boolean value is returned to indicate success.
//
// This structure encapsulates both possible return types for easier handling in Go.
type MessageOrBool struct {
	Message *types.Message
	Bool    *bool
}

type EditMessageText[T int | string] struct {
	ChatId               *T
	Text                 string
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ParseMode            *string
	Entities             *[]types.MessageEntity
	LinkPreviewOptions   *types.LinkPreviewOptions
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (e EditMessageText[T]) Validate() error {
	if len(e.Text) < 1 || len(e.Text) > 4096 {
		return errors.ErrInvalidParam("text parameter must be between 1 and 4096 characters")
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.Entities != nil && e.ParseMode != nil {
		return errors.ErrInvalidParam("entities can't be used if parse_mode is provided")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "editMessageText", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageText", e)
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
	Entities              *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	ReplyMarkup           *types.InlineKeyboardMarkup
}

func (e EditMessageCaption[T]) Validate() error {
	if e.Caption != nil {
		if len(*e.Caption) < 1 || len(*e.Caption) > 4096 {
			return errors.ErrInvalidParam("caption parameter must be between 1 and 4096 characters")
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
			}
		}
	}
	if e.Entities != nil && e.ParseMode != nil {
		return errors.ErrInvalidParam("entities can't be used if parse_mode is provided")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "editMessageCaption", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageCaption", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageMedia[T int | string] struct {
	Media                 types.InputMedia
	ChatId                *T
	BusinessConnectionId  *string
	MessageId             *int
	InlineMessageId       *string
	ShowCaptionAboveMedia *bool
	ReplyMarkup           *types.InlineKeyboardMarkup
}

func (e EditMessageMedia[T]) Validate() error {
	if err := e.Media.Validate(); err != nil {
		return err
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "editMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageMedia", e)
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
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (e EditMessageLiveLocation[T]) Validate() error {
	if e.Latitude == nil {
		return errors.ErrInvalidParam("latitude parameter can't be empty")
	}
	if e.Longtitude == nil {
		return errors.ErrInvalidParam("longtitude parameter can't be empty")
	}
	if e.HorizontalAccuracy != nil {
		if *e.HorizontalAccuracy < 0 || *e.HorizontalAccuracy > 1500 {
			return errors.ErrInvalidParam("horizontal_accuracy parameter must be between 0 and 1500 meetrs")
		}
	}
	if e.Heading != nil {
		if *e.Heading < 1 || *e.Heading > 360 {
			return errors.ErrInvalidParam("heading parameter must be between 1 and 360 degrees")
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "editMessageLiveLocation", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageLiveLocation", e)
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
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (e StopMessageLiveLocation[T]) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "stopMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageMedia", e)
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
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (e EditMessageReplyMarkup[T]) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if c, ok := any(*e.ChatId).(string); ok {
				if strings.TrimSpace(c) == "" {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
			if c, ok := any(*e.ChatId).(int); ok {
				if c < 1 {
					return errors.ErrInvalidParam("chat_id parameter can't be empty")
				}
			}
		}
		if e.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		} else {
			if *e.MessageId < 1 {
				return errors.ErrInvalidParam("message_id parameter can't be empty")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}
