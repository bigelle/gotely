package methods

import (
	"strings"

	"github.com/bigelle/gotely/api/objects"
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

// Use this method to edit text and game messages.
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageText struct {
	//Optional:
	//Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Optional:
	//Required if inline_message_id is not specified.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	//Optional:
	//Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	//Optional:
	//Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	//Required:
	//New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`
	//Optional:
	//Mode for parsing entities in the message text.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional:
	//A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities *[]objects.MessageEntity `json:"entities,omitempty"`
	//Optional:
	//Link preview generation options for the message
	LinkPreviewOptions *objects.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	//Optional:
	//A JSON-serialized object for an inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageText) Validate() error {
	if len(e.Text) < 1 || len(e.Text) > 4096 {
		return objects.ErrInvalidParam("text parameter must be between 1 and 4096 characters")
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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

// Use this method to edit captions of messages.
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageCaption struct {
	//Optional:
	//Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Optional:
	//Required if inline_message_id is not specified.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	//Optional:
	//Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	//Optional:
	//Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	//Optional:
	//New caption of the message, 0-1024 characters after entities parsing
	Caption *string `json:"caption"`
	//Optional:
	//Mode for parsing entities in the message caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional:
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Optional:
	//Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional:
	//A JSON-serialized object for an inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageCaption) Validate() error {
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
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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
	if e.CaptionEntities != nil && e.ParseMode != nil {
		return objects.ErrInvalidParam("entities can't be used if parse_mode is provided")
	}
	if e.ReplyMarkup != nil {
		if err := e.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise. When an inline message is edited,
// a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL.
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageMedia struct {
	Media                 objects.InputMedia            `json:"media"`
	ChatId                *string                       `json:"chat_id,omitempty"`
	BusinessConnectionId  *string                       `json:"business_connection_id,omitempty"`
	MessageId             *int                          `json:"message_id,omitempty"`
	InlineMessageId       *string                       `json:"inline_message_id,omitempty"`
	ShowCaptionAboveMedia *bool                         `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageMedia) Validate() error {
	if err := e.Media.Validate(); err != nil {
		return err
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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

type EditMessageLiveLocation struct {
	Latitude             *float64                      `json:"latitude"`
	Longitude            *float64                      `json:"longitude"`
	LivePeriod           *int                          `json:"live_period,omitempty"`
	HorizontalAccuracy   *float64                      `json:"horizontal_accuracy,omitempty"`
	Heading              *int                          `json:"heading,omitempty"`
	ProximityAlertRadius *int                          `json:"proximity_alert_radius,omitempty"`
	ChatId               *string                       `json:"chat_id,omitempty"`
	BusinessConnectionId *string                       `json:"business_connection_id,omitempty"`
	MessageId            *int                          `json:"message_id,omitempty"`
	InlineMessageId      *string                       `json:"inline_message_id,omitempty"`
	ReplyMarkup          *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageLiveLocation) Validate() error {
	if e.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if e.Longitude == nil {
		return objects.ErrInvalidParam("longitude parameter can't be empty")
	}
	if e.HorizontalAccuracy != nil {
		if *e.HorizontalAccuracy < 0 || *e.HorizontalAccuracy > 1500 {
			return objects.ErrInvalidParam("horizontal_accuracy parameter must be between 0 and 1500 meters")
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
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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

type StopMessageLiveLocation struct {
	ChatId               *string                       `json:"chat_id,omitempty"`
	BusinessConnectionId *string                       `json:"business_connection_id,omitempty"`
	MessageId            *int                          `json:"message_id,omitempty"`
	InlineMessageId      *string                       `json:"inline_message_id,omitempty"`
	ReplyMarkup          *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e StopMessageLiveLocation) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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

type EditMessageReplyMarkup struct {
	ChatId               *string                       `json:"chat_id,omitempty"`
	BusinessConnectionId *string                       `json:"business_connection_id,omitempty"`
	MessageId            *int                          `json:"message_id,omitempty"`
	InlineMessageId      *string                       `json:"inline_message_id,omitempty"`
	ReplyMarkup          *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageReplyMarkup) Validate() error {
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id parameter can'be empty if chat_id and message_id are not specified")
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil && len(*e.ChatId) == 0 {
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

type StopPoll struct {
	ChatId               string                        `json:"chat_id"`
	MessageId            int                           `json:"message_id"`
	BusinessConnectionId *string                       `json:"business_connection_id,omitempty"`
	ReplyMarkup          *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s StopPoll) Validate() error {
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

type DeleteMessage struct {
	ChatId    string `json:"chat_id"`
	MessageId int    `json:"message_id"`
}

func (d DeleteMessage) Validate() error {
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

type DeleteMessages struct {
	ChatId     string `json:"chat_id"`
	MessageIds []int  `json:"message_ids"`
}

func (d DeleteMessages) Validate() error {
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
