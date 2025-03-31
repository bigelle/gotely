package methods

import (
	"fmt"
	"io"
	"mime/multipart"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/objects"
)

// Use this method to edit text and game messages.
// On success, if the edited message is not an inline message, the edited [objects.Message] is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageText struct {
	// REQUIRED:
	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// Mode for parsing entities in the message text.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities *[]objects.MessageEntity `json:"entities,omitempty"`
	// Link preview generation options for the message
	LinkPreviewOptions *objects.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageText) Validate() error {
	var err gotely.ErrFailedValidation
	if len(e.Text) < 1 || len(e.Text) > 4096 {
		err = append(err, fmt.Errorf("text parameter must be between 1 and 4096 characters"))
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.Entities != nil && e.ParseMode != nil {
		err = append(err, fmt.Errorf("entities can't be used if parse_mode is provided"))
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if e.LinkPreviewOptions != nil {
		if er := e.LinkPreviewOptions.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s EditMessageText) Endpoint() string {
	return "editMessageText"
}

func (s EditMessageText) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditMessageText) ContentType() string {
	return "application/json"
}

// Use this method to edit captions of messages.
// On success, if the edited message is not an inline message, the edited [objects.Message] is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageCaption struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// New caption of the message, 0-1024 characters after entities parsing
	Caption *string `json:"caption"`
	// Mode for parsing entities in the message caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media. Supported only for animation, photo and video messages.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageCaption) Validate() error {
	var err gotely.ErrFailedValidation
	if e.Caption != nil {
		if len(*e.Caption) < 1 || len(*e.Caption) > 4096 {
			err = append(err, fmt.Errorf("caption parameter must be between 1 and 4096 characters"))
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.CaptionEntities != nil && e.ParseMode != nil {
		err = append(err, fmt.Errorf("entities can't be used if parse_mode is provided"))
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s EditMessageCaption) Endpoint() string {
	return "editMessageCaption"
}

func (s EditMessageCaption) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditMessageCaption) ContentType() string {
	return "application/json"
}

// Use this method to edit animation, audio, document, photo, or video messages, or to add media to text messages.
// If a message is part of a message album, then it can be edited only to an audio for audio albums,
// only to a document for document albums and to a photo or a video otherwise. When an inline message is edited,
// a new file can't be uploaded; use a previously uploaded file via its file_id or specify a URL.
// On success, if the edited message is not an inline message, the edited [objects.Message] is returned, otherwise True is returned.
// Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageMedia struct {
	// REQUIRED:
	// A JSON-serialized object for a new media content of the message
	Media objects.InputMedia `json:"media"`

	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	contentType string
}

func (e EditMessageMedia) Validate() error {
	var err gotely.ErrFailedValidation
	if er := e.Media.Validate(); er != nil {
		err = append(err, er)
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s EditMessageMedia) Endpoint() string {
	return "editMessageMedia"
}

func (s *EditMessageMedia) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := s.Media.WriteTo(mw); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ChatId != nil {
			if err := mw.WriteField("chat_id", *s.ChatId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageId != nil {
			if err := mw.WriteField("message_id", fmt.Sprint(*s.MessageId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.InlineMessageId != nil {
			if err := mw.WriteField("inline_message_id", *s.InlineMessageId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyMarkup != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_markup", *s.ReplyMarkup); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s EditMessageMedia) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to edit live location messages.
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to [StopMessageLiveLocation].
// On success, if the edited message is not an inline message, the edited [objects.Message] is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	// REQUIRED:
	// Latitude of new location
	Latitude *float64 `json:"latitude"`
	// REQUIRED:
	// Longitude of new location
	Longitude *float64 `json:"longitude"`

	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// New period in seconds during which the location can be updated, starting from the message send date.
	// If 0x7FFFFFFF is specified, then the location can be updated forever.
	// Otherwise, the new value must not exceed the current live_period by more than a day,
	// and the live location expiration date must remain within the next 90 days.
	// If not specified, then live_period remains unchanged
	LivePeriod *int `json:"live_period,omitempty"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
	// Direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int `json:"heading,omitempty"`
	// The maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int `json:"proximity_alert_radius,omitempty"`
	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageLiveLocation) Validate() error {
	var err gotely.ErrFailedValidation
	if e.Latitude == nil {
		err = append(err, fmt.Errorf("latitude parameter can't be empty"))
	}
	if e.Longitude == nil {
		err = append(err, fmt.Errorf("longitude parameter can't be empty"))
	}
	if e.HorizontalAccuracy != nil {
		if *e.HorizontalAccuracy < 0 || *e.HorizontalAccuracy > 1500 {
			err = append(err, fmt.Errorf("horizontal_accuracy parameter must be between 0 and 1500 meters"))
		}
	}
	if e.Heading != nil {
		if *e.Heading < 1 || *e.Heading > 360 {
			err = append(err, fmt.Errorf("heading parameter must be between 1 and 360 degrees"))
		}
	}
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s EditMessageLiveLocation) Endpoint() string {
	return "editMessageLiveLocation"
}

func (s EditMessageLiveLocation) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditMessageLiveLocation) ContentType() string {
	return "application/json"
}

// Use this method to stop updating a live location message before live_period expires.
// On success, if the message is not an inline message, the edited [objects.Message] is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message with live location to stop
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// A JSON-serialized object for a new inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e StopMessageLiveLocation) Validate() error {
	var err gotely.ErrFailedValidation
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Use this method to edit only the reply markup of messages.
// On success, if the edited message is not an inline message, the edited [objects.Message] is returned,
// otherwise True is returned. Note that business messages that were not sent by the bot and
// do not contain an inline keyboard can only be edited within 48 hours from the time they were sent.
type EditMessageReplyMarkup struct {
	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Required if inline_message_id is not specified.
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId *string `json:"chat_id,omitempty"`
	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageId *int `json:"message_id,omitempty"`
	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageId *string `json:"inline_message_id,omitempty"`
	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (e EditMessageReplyMarkup) Validate() error {
	var err gotely.ErrFailedValidation
	if e.ChatId == nil && e.MessageId == nil {
		if e.InlineMessageId == nil {
			err = append(err, fmt.Errorf("inline_message_id parameter can'be empty if chat_id and message_id are not specified"))
		}
	}
	if e.InlineMessageId == nil {
		if e.ChatId == nil || len(*e.ChatId) == 0 {
			err = append(err, fmt.Errorf("chat_id parameter can't be empty if inline_message_id is not specified"))
		}
		if e.MessageId == nil || *e.MessageId < 1 {
			err = append(err, fmt.Errorf("message_id parameter can't be empty if inline_message_id is not specified"))
		}
	}
	if e.ReplyMarkup != nil {
		if er := e.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

// Use this method to stop a poll which was sent by the bot.
// On success, the stopped [objects.Poll] is returned.
type StopPoll struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Identifier of the original message with the poll
	MessageId int `json:"message_id"`

	// Unique identifier of the business connection on behalf of which the message to be edited was sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup *objects.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (s StopPoll) Validate() error {
	var err gotely.ErrFailedValidation
	if s.ChatId == "" {
		err = append(err, fmt.Errorf("chat_id parameter can't be empty"))
	}
	if s.MessageId < 1 {
		err = append(err, fmt.Errorf("message_id parameter can't be empty"))
	}
	if s.ReplyMarkup != nil {
		if er := s.ReplyMarkup.Validate(); er != nil {
			err = append(err, er)
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s StopPoll) Endpoint() string {
	return "stopPoll"
}

func (s StopPoll) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s StopPoll) ContentType() string {
	return "application/json"
}

// Use this method to delete a message, including service messages, with the following limitations:
//
//   - A message can only be deleted if it was sent less than 48 hours ago.
//
//   - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
//
//   - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//
//   - Bots can delete outgoing messages in private chats, groups, and supergroups.
//
//   - Bots can delete incoming messages in private chats.
//
//   - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//
//   - If the bot is an administrator of a group, it can delete any message there.
//
//   - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
//
// Returns True on success.
type DeleteMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Identifier of the message to delete
	MessageId int `json:"message_id"`
}

func (d DeleteMessage) Validate() error {
	var err gotely.ErrFailedValidation
	if d.ChatId == "" {
		err = append(err, fmt.Errorf("chat_id parameter can't be empty"))
	}
	if d.MessageId < 1 {
		err = append(err, fmt.Errorf("message_id parameter can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s DeleteMessage) Endpoint() string {
	return "deleteMessage"
}

func (s DeleteMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteMessage) ContentType() string {
	return "application/json"
}

// Use this method to delete multiple messages simultaneously.
// If some of the specified messages can't be found, they are skipped.
// Returns True on success.
type DeleteMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// A JSON-serialized list of 1-100 identifiers of messages to delete.
	// See deleteMessage for limitations on which messages can be deleted
	MessageIds []int `json:"message_ids"`
}

func (d DeleteMessages) Validate() error {
	var err gotely.ErrFailedValidation
	if d.ChatId == "" {
		err = append(err, fmt.Errorf("chat_id parameter can't be empty"))
	}
	if len(d.MessageIds) < 1 || len(d.MessageIds) > 100 {
		err = append(err, fmt.Errorf("message_ids parameter must be between 1 and 100"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s DeleteMessages) Endpoint() string {
	return "deleteMessages"
}

func (s DeleteMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteMessages) ContentType() string {
	return "application/json"
}
