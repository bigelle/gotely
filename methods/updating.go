package methods

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bigelle/gotely/objects"
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
	client      *http.Client
	baseUrl     string
}

func (s *EditMessageText) WithClient(c *http.Client) *EditMessageText {
	s.client = c
	return s
}

func (s EditMessageText) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditMessageText) WithApiBaseUrl(u string) *EditMessageText {
	s.baseUrl = u
	return s
}

func (s EditMessageText) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditMessageText) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageText) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "editMessageText", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageText", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
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
	client      *http.Client
	baseUrl     string
}

func (s *EditMessageCaption) WithClient(c *http.Client) *EditMessageCaption {
	s.client = c
	return s
}

func (s EditMessageCaption) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditMessageCaption) WithApiBaseUrl(u string) *EditMessageCaption {
	s.baseUrl = u
	return s
}

func (s EditMessageCaption) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditMessageCaption) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageCaption) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "editMessageCaption", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageCaption", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
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
	client                *http.Client
	baseUrl               string
}

func (s *EditMessageMedia) WithClient(c *http.Client) *EditMessageMedia {
	s.client = c
	return s
}

func (s EditMessageMedia) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditMessageMedia) WithApiBaseUrl(u string) *EditMessageMedia {
	s.baseUrl = u
	return s
}

func (s EditMessageMedia) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditMessageMedia) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageMedia) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "editMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageMedia", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageLiveLocation struct {
	Latitude             *float64
	Longtitude           *float64
	LivePeriod           *int
	HorizontalAccuracy   *float64
	Heading              *int
	ProximityAlertRadius *int
	ChatId               *string
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
	client               *http.Client
	baseUrl              string
}

func (s *EditMessageLiveLocation) WithClient(c *http.Client) *EditMessageLiveLocation {
	s.client = c
	return s
}

func (s EditMessageLiveLocation) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditMessageLiveLocation) WithApiBaseUrl(u string) *EditMessageLiveLocation {
	s.baseUrl = u
	return s
}

func (s EditMessageLiveLocation) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (e EditMessageLiveLocation) Validate() error {
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

func (e EditMessageLiveLocation) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageLiveLocation) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "editMessageLiveLocation", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageLiveLocation", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type StopMessageLiveLocation struct {
	ChatId               *string
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
	client               *http.Client
	baseUrl              string
}

func (s *StopMessageLiveLocation) WithClient(c *http.Client) *StopMessageLiveLocation {
	s.client = c
	return s
}

func (s StopMessageLiveLocation) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *StopMessageLiveLocation) WithApiBaseUrl(u string) *StopMessageLiveLocation {
	s.baseUrl = u
	return s
}

func (s StopMessageLiveLocation) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e StopMessageLiveLocation) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e StopMessageLiveLocation) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "stopMessageMedia", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageMedia", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type EditMessageReplyMarkup struct {
	ChatId               *string
	BusinessConnectionId *string
	MessageId            *int
	InlineMessageId      *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
	client               *http.Client
	baseUrl              string
}

func (s *EditMessageReplyMarkup) WithClient(c *http.Client) *EditMessageReplyMarkup {
	s.client = c
	return s
}

func (s EditMessageReplyMarkup) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditMessageReplyMarkup) WithApiBaseUrl(u string) *EditMessageReplyMarkup {
	s.baseUrl = u
	return s
}

func (s EditMessageReplyMarkup) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditMessageReplyMarkup) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditMessageReplyMarkup) Execute(token string) (MessageOrBool, error) {
	if e.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "editMessageReplyMarkup", e)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type StopPoll struct {
	ChatId               string
	MessageId            int
	BusinessConnectionId *string
	ReplyMarkup          *objects.InlineKeyboardMarkup
	client               *http.Client
	baseUrl              string
}

func (s *StopPoll) WithClient(c *http.Client) *StopPoll {
	s.client = c
	return s
}

func (s StopPoll) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *StopPoll) WithApiBaseUrl(u string) *StopPoll {
	s.baseUrl = u
	return s
}

func (s StopPoll) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s StopPoll) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s StopPoll) Execute(token string) (*objects.Poll, error) {
	return SendTelegramPostRequest[objects.Poll](token, "stopPoll", s)
}

type DeleteMessage struct {
	ChatId    string
	MessageId int
	client    *http.Client
	baseUrl   string
}

func (s *DeleteMessage) WithClient(c *http.Client) *DeleteMessage {
	s.client = c
	return s
}

func (s DeleteMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteMessage) WithApiBaseUrl(u string) *DeleteMessage {
	s.baseUrl = u
	return s
}

func (s DeleteMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (d DeleteMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessage) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteMessage", d)
}

type DeleteMessages struct {
	ChatId     string
	MessageIds []int
	client     *http.Client
	baseUrl    string
}

func (s *DeleteMessages) WithClient(c *http.Client) *DeleteMessages {
	s.client = c
	return s
}

func (s DeleteMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteMessages) WithApiBaseUrl(u string) *DeleteMessages {
	s.baseUrl = u
	return s
}

func (s DeleteMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (d DeleteMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessages) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteMessages", d)
}
