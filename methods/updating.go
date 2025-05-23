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

		if err := gotely.WriteJSONToForm(mw, "media", s.Media); err != nil {
			pw.CloseWithError(err)
			return
		}
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

// Sends a gift to the given user or channel chat.
// The gift can't be converted to Telegram Stars by the receiver.
// Returns True on success.
type SendGift struct {
	// REQUIRED:
	// Identifier of the gift
	GiftId string `json:"gift_id"`

	// Text that will be shown along with the gift; 0-128 characters
	Text *string `json:"text"`
	// Required if chat_id is not specified. Unique identifier of the target user who will receive the gift.
	UserId *int `json:"user_id"`
	// Required if user_id is not specified.
	// Unique identifier for the chat or username of the channel (in the format @channelusername) that will receive the gift.
	ChatId *string `json:"chat_id,omitempty,"`
	// Pass True to pay for the gift upgrade from the bot's balance,
	// thereby making the upgrade free for the receiver
	PayForUpgrade *bool `json:"pay_for_upgrade,omitempty"`
	// Mode for parsing entities in the text. See formatting options for more details.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty,"`
	// A JSON-serialized list of special entities that appear in the gift text.
	// It can be specified instead of text_parse_mode.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextEntities *[]objects.MessageEntity `json:"text_entities,omitempty,"`
}

func (s SendGift) Validate() error {
	var err gotely.ErrFailedValidation
	if s.UserId != nil {
		if *s.UserId < 1 {
			err = append(err, fmt.Errorf("user_id parameter can't be empty"))
		}
	}
	if s.ChatId != nil {
		if *s.ChatId == "" {
			err = append(err, fmt.Errorf("user_id parameter can't be empty"))
		}
	}
	if s.GiftId == "" {
		err = append(err, fmt.Errorf("gift_id parameter can't be empty"))
	}
	if s.Text != nil {
		if len(*s.Text) > 255 {
			err = append(err, fmt.Errorf("text parameter must not be longer than 255 characters"))
		}
	}
	if s.TextParseMode != nil && s.TextEntities != nil {
		err = append(err, fmt.Errorf("parse_mode can't be used if entities are provided"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SendGift) Endpoint() string {
	return "sendGift"
}

func (s SendGift) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendGift) ContentType() string {
	return "application/json"
}

// Gifts a Telegram Premium subscription to the given user.
// Returns True on success.
type GiftPremiumSubscription struct {
	// Unique identifier of the target user who will receive a Telegram Premium subscription
	UserId int `json:"user_id"`
	// Number of months the Telegram Premium subscription will be active for the user; must be one of 3, 6, or 12
	MonthCount int `json:"month_count"`
	// Number of Telegram Stars to pay for the Telegram Premium subscription;
	// must be 1000 for 3 months, 1500 for 6 months, and 2500 for 12 months
	StarCount int `json:"star_count"`
	// Text that will be shown along with the service message about the subscription; 0-128 characters
	Text *string `json:"text,omitempty"`
	// Mode for parsing entities in the text. See formatting options for more details.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextParseMode *string `json:"text_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the gift text.
	// It can be specified instead of text_parse_mode.
	// Entities other than “bold”, “italic”, “underline”, “strikethrough”, “spoiler”, and “custom_emoji” are ignored.
	TextEntities *[]objects.MessageEntity `json:"text_entities,omitempty"`
}

func (g GiftPremiumSubscription) Validate() error {
	var err gotely.ErrFailedValidation
	if g.UserId == 0 {
		err = append(err, fmt.Errorf("user_id can't be empty"))
	}
	if g.MonthCount != 3 && g.MonthCount != 6 && g.MonthCount != 12 {
		err = append(err, fmt.Errorf("month_count must be one of 3, 6 or 12"))
	}
	if (g.MonthCount == 3 && g.StarCount != 1000) || (g.MonthCount == 6 && g.StarCount != 1500) || (g.MonthCount == 12 && g.StarCount != 2500) {
		err = append(err, fmt.Errorf("star_count must be 1000 for 3 months, 1500 for 6 months, and 2500 for 12 months"))
	}
	if g.Text != nil {
		if len(*g.Text) > 128 {
			err = append(err, fmt.Errorf("text must be between 0 and 128 characters if specified"))
		}
	}
	if g.TextEntities != nil {
		for _, ent := range *g.TextEntities {
			if e := ent.Validate(); e != nil {
				err = append(err, e)
				break
			}
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (g GiftPremiumSubscription) Endpoint() string {
	return "giftPremiumSubscription"
}

func (g GiftPremiumSubscription) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g GiftPremiumSubscription) ContentType() string {
	return "application/json"
}

// Verifies a user on behalf of the organization which is represented by the bot.
// Returns True on success.
type VerifyUser struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// Custom description for the verification; 0-70 characters.
	// Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty"`
}

func (v VerifyUser) Validate() error {
	var err gotely.ErrFailedValidation
	if v.UserId <= 0 {
		err = append(err, fmt.Errorf("user_id can't be empty or negative"))
	}
	if v.CustomDescription != nil {
		if len(*v.CustomDescription) > 70 {
			err = append(err, fmt.Errorf("custom_description must be between 0 and 70 characters"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s VerifyUser) Endpoint() string {
	return "verifyUser"
}

func (s VerifyUser) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s VerifyUser) ContentType() string {
	return "application/json"
}

// Verifies a chat on behalf of the organization which is represented by the bot.
// Returns True on success.
type VerifyChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`

	// Custom description for the verification; 0-70 characters.
	// Must be empty if the organization isn't allowed to provide a custom verification description.
	CustomDescription *string `json:"custom_description,omitempty"`
}

func (v VerifyChat) Validate() error {
	var err gotely.ErrFailedValidation
	if v.ChatId == "" {
		err = append(err, fmt.Errorf("user_id can't be empty or negative"))
	}
	if v.CustomDescription != nil {
		if len(*v.CustomDescription) > 70 {
			err = append(err, fmt.Errorf("custom_description must be between 0 and 70 characters"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s VerifyChat) Endpoint() string {
	return "verifyChat"
}

func (s VerifyChat) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s VerifyChat) ContentType() string {
	return "application/json"
}

// Removes verification from a user who is currently verified on behalf of the organization represented by the bot.
// Returns True on success.
type RemoveUserVerification struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (r RemoveUserVerification) Validate() error {
	var err gotely.ErrFailedValidation
	if r.UserId <= 0 {
		err = append(err, fmt.Errorf("user_id can't be empty or negative"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s RemoveUserVerification) Endpoint() string {
	return "removeUserVerification"
}

func (s RemoveUserVerification) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s RemoveUserVerification) ContentType() string {
	return "application/json"
}

// Removes verification from a chat that is currently verified on behalf of the organization represented by the bot.
// Returns True on success.
type RemoveChatVerification struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (r RemoveChatVerification) Validate() error {
	var err gotely.ErrFailedValidation
	if r.ChatId == "" {
		err = append(err, fmt.Errorf("chat_id can't be empty or negative"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s RemoveChatVerification) Endpoint() string {
	return "removeChatVerification"
}

func (s RemoveChatVerification) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s RemoveChatVerification) ContentType() string {
	return "application/json"
}

// Marks incoming message as read on behalf of a business account.
// Requires the can_read_messages business bot right.
// Returns True on success.
type ReadBusinessMessage struct {
	// REQUIRED:
	// Unique identifier of the business connection on behalf of which to read the message
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Unique identifier of the chat in which the message was received.
	// The chat must have been active in the last 24 hours.
	ChatId int `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the message to mark as read
	MessageId int `json:"message_id"`
}

func (r ReadBusinessMessage) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.ChatId == 0 {
		err = append(err, fmt.Errorf("chat_id can't be empty"))
	}
	if r.MessageId == 0 {
		err = append(err, fmt.Errorf("message_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s ReadBusinessMessage) Endpoint() string {
	return "readBusinessMessage"
}

func (s ReadBusinessMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ReadBusinessMessage) ContentType() string {
	return "application/json"
}

// Delete messages on behalf of a business account.
// Requires the can_delete_outgoing_messages business bot right to delete messages sent by the bot itself,
// or the can_delete_all_messages business bot right to delete any message.
// Returns True on success.
type DeleteBusinessMessage struct {
	// REQUIRED:
	// Unique identifier of the business connection on behalf of which to delete the messages
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// A JSON-serialized list of 1-100 identifiers of messages to delete.
	// All messages must be from the same chat.
	// See deleteMessage for limitations on which messages can be deleted
	MessageIds []int `json:"message_ids"`
}

func (r DeleteBusinessMessage) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if len(r.MessageIds) < 1 || len(r.MessageIds) > 100 {
		err = append(err, fmt.Errorf("message_ids accepts only 1-100 IDs"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s DeleteBusinessMessage) Endpoint() string {
	return "deleteBusinessMessage"
}

func (s DeleteBusinessMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteBusinessMessage) ContentType() string {
	return "application/json"
}

// Changes the first and last name of a managed business account.
// Requires the can_change_name business bot right.
// Returns True on success.
type SetBusinessAccountName struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// The new value of the first name for the business account; 1-64 characters
	FirstName string `json:"first_name"`

	// The new value of the last name for the business account; 0-64 characters
	LastName *string `json:"last_name,omitempty"`
}

func (r SetBusinessAccountName) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if len(r.FirstName) < 1 || len(r.FirstName) > 64 {
		err = append(err, fmt.Errorf("first_name accepts only 1-64 characters"))
	}
	if r.LastName != nil {
		if len(*r.LastName) < 0 || len(*r.LastName) > 64 {
			err = append(err, fmt.Errorf("last_name accepts only 0-64 characters if specified"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetBusinessAccountName) Endpoint() string {
	return "setBusinessAccountName"
}

func (s SetBusinessAccountName) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetBusinessAccountName) ContentType() string {
	return "application/json"
}

// Changes the username of a managed business account.
// Requires the can_change_username business bot right.
// Returns True on success.
type SetBusinessAccountUsername struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`

	// The new value of the username for the business account; 0-32 characters
	Username *string `json:"username,omitempty"`
}

func (r SetBusinessAccountUsername) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.Username != nil {
		if len(*r.Username) < 0 || len(*r.Username) > 32 {
			err = append(err, fmt.Errorf("username accepts only 0-32 characters if specified"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetBusinessAccountUsername) Endpoint() string {
	return "setBusinessAccountUsername"
}

func (s SetBusinessAccountUsername) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetBusinessAccountUsername) ContentType() string {
	return "application/json"
}

// Changes the bio of a managed business account.
// Requires the can_change_bio business bot right.
// Returns True on success.
type SetBusinessAccountBio struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`

	// The new value of the bio for the business account; 0-140 characters
	Bio *string `json:"bio,omitempty"`
}

func (r SetBusinessAccountBio) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.Bio != nil {
		if len(*r.Bio) < 0 || len(*r.Bio) > 140 {
			err = append(err, fmt.Errorf("bio accepts only 0-140 characters if specified"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetBusinessAccountBio) Endpoint() string {
	return "setBusinessAccountBi"
}

func (s SetBusinessAccountBio) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetBusinessAccountBio) ContentType() string {
	return "application/json"
}

// Changes the profile photo of a managed business account.
// Requires the can_edit_profile_photo business bot right.
// Returns True on success.
type SetBusinessAccountProfilePhoto struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// The new profile photo to set
	Photo objects.InputProfilePhoto `json:"photo"`

	// Pass True to set the public photo, which will be visible even if the main photo is hidden by the business account's privacy settings.
	// An account can have only one public photo.
	IsPublic *bool `json:"is_public,omitempty"`

	contentType string
}

func (r SetBusinessAccountProfilePhoto) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if e := r.Photo.Validate(); e != nil {
		err = append(err, e)
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetBusinessAccountProfilePhoto) Endpoint() string {
	return "setBusinessAccountProfilePhoto"
}

func (s *SetBusinessAccountProfilePhoto) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("business_connection_id", s.BusinessConnectionId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.IsPublic != nil {
			if err := mw.WriteField("is_public", fmt.Sprint(*s.IsPublic)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if err := gotely.WriteJSONToForm(mw, "photo", s.Photo); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := s.Photo.WriteTo(mw); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()

	return pr
}

func (s SetBusinessAccountProfilePhoto) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Removes the current profile photo of a managed business account.
// Requires the can_edit_profile_photo business bot right.
// Returns True on success.
type RemoveBusinessAccountProfilePhoto struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`

	// Pass True to remove the public photo, which is visible even if the main photo is hidden by the business account's privacy settings.
	// After the main photo is removed, the previous profile photo (if present) becomes the main photo.
	IsPublic *bool `json:"is_public,omitempty"`
}

func (r RemoveBusinessAccountProfilePhoto) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s RemoveBusinessAccountProfilePhoto) Endpoint() string {
	return "removeBusinessAccountProfilePhoto"
}

func (s RemoveBusinessAccountProfilePhoto) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s RemoveBusinessAccountProfilePhoto) ContentType() string {
	return "application/json"
}

// Changes the privacy settings pertaining to incoming gifts in a managed business account.
// Requires the can_change_gift_settings business bot right.
// Returns True on success.
type SetBusinessAccountGiftSettings struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Pass True, if a button for sending a gift to the user or by the business account must always be shown in the input field
	ShowGiftButton bool `json:"show_gift_button"`
	// REQUIRED:
	// Types of gifts accepted by the business account
	AcceptedGiftTypes objects.AcceptedGiftTypes `json:"accepted_gift_types"`
}

func (r SetBusinessAccountGiftSettings) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s SetBusinessAccountGiftSettings) Endpoint() string {
	return "setBusinessAccountGiftSettings"
}

func (s SetBusinessAccountGiftSettings) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetBusinessAccountGiftSettings) ContentType() string {
	return "application/json"
}

// Returns the amount of Telegram Stars owned by a managed business account.
// Requires the can_view_gifts_and_stars business bot right.
// Returns [objects.StarAmount] on success.
type GetBusinessAccountStarBalance struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
}

func (r GetBusinessAccountStarBalance) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s GetBusinessAccountStarBalance) Endpoint() string {
	return "getBusinessAccountStarBalance"
}

func (s GetBusinessAccountStarBalance) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetBusinessAccountStarBalance) ContentType() string {
	return "application/json"
}

// Transfers Telegram Stars from the business account balance to the bot's balance.
// Requires the can_transfer_stars business bot right.
// Returns True on success.
type TransferBusinessAccountStars struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Number of Telegram Stars to transfer; 1-10000
	StarCount int `json:"star_count"`
}

func (r TransferBusinessAccountStars) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.StarCount < 1 || r.StarCount > 10_000 {
		err = append(err, fmt.Errorf("star_count must be between 1 and 10 000"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s TransferBusinessAccountStars) Endpoint() string {
	return "transferBusinessAccountStar"
}

func (s TransferBusinessAccountStars) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s TransferBusinessAccountStars) ContentType() string {
	return "application/json"
}

// Returns the gifts received and owned by a managed business account.
// Requires the can_view_gifts_and_stars business bot right.
// Returns [objects.OwnedGifts] on success.
type GetBusinessAccountGifts struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`

	// Pass True to exclude gifts that aren't saved to the account's profile page
	ExcludeUnsaved *bool `json:"exclude_unsaved,omitempty"`
	// Pass True to exclude gifts that are saved to the account's profile page
	ExcludeSaved *bool `json:"exclude_saved,omitempty"`
	// Pass True to exclude gifts that can be purchased an unlimited number of times
	ExcludeUnlimited *bool `json:"exclude_unlimited,omitempty"`
	// Pass True to exclude gifts that can be purchased a limited number of times
	ExcludeLimited *bool `json:"exclude_limited,omitempty"`
	// Pass True to exclude unique gifts
	ExcludeUnique *bool `json:"exclude_unique,omitempty"`
	// Pass True to sort results by gift price instead of send date.
	// Sorting is applied before pagination.
	SortByPrice *bool `json:"sort_by_price,omitempty"`
	// Offset of the first entry to return as received from the previous request;
	// use empty string to get the first chunk of results
	Offset *string `json:"offset,omitempty"`
	// The maximum number of gifts to be returned; 1-100. Defaults to 100
	Limit *int `json:"limit,omitempty"`
}

func (r GetBusinessAccountGifts) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.Limit != nil {
		if *r.Limit < 1 || *r.Limit > 100 {
			err = append(err, fmt.Errorf("limit must be between 1 and 100"))
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s GetBusinessAccountGifts) Endpoint() string {
	return "getBusinessAccountGifts"
}

func (s GetBusinessAccountGifts) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetBusinessAccountGifts) ContentType() string {
	return "application/json"
}

// Converts a given regular gift to Telegram Stars.
// Requires the can_convert_gifts_to_stars business bot right.
// Returns True on success.
type ConvertGiftToStarts struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Unique identifier of the regular gift that should be converted to Telegram Stars
	OwnedGiftId string `json:"owned_gift_id"`
}

func (r ConvertGiftToStarts) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.OwnedGiftId == "" {
		err = append(err, fmt.Errorf("owned_gift_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s ConvertGiftToStarts) Endpoint() string {
	return "convertGiftToStarts"
}

func (s ConvertGiftToStarts) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ConvertGiftToStarts) ContentType() string {
	return "application/json"
}

// Upgrades a given regular gift to a unique gift.
// Requires the can_transfer_and_upgrade_gifts business bot right.
// Additionally requires the can_transfer_stars business bot right if the upgrade is paid.
// Returns True on success.
type UpgradeGift struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Unique identifier of the regular gift that should be upgraded to a unique one
	OwnedGiftId string `json:"owned_gift_id"`

	// Pass True to keep the original gift text, sender and receiver in the upgraded gift
	KeepOriginalDetails *bool `json:"keep_original_details,omitempty"`
	// The amount of Telegram Stars that will be paid for the upgrade from the business account balance.
	// If gift.prepaid_upgrade_star_count > 0, then pass 0, otherwise,
	// the can_transfer_stars business bot right is required and gift.upgrade_star_count must be passed.
	StartCount *int `json:"star_count,omitempty"`
}

func (r UpgradeGift) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.OwnedGiftId == "" {
		err = append(err, fmt.Errorf("owned_gift_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s UpgradeGift) Endpoint() string {
	return "upgradeGift"
}

func (s UpgradeGift) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UpgradeGift) ContentType() string {
	return "application/json"
}

// Transfers an owned unique gift to another user.
// Requires the can_transfer_and_upgrade_gifts business bot right.
// Requires can_transfer_stars business bot right if the transfer is paid.
// Returns True on success.
type TransferGift struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Unique identifier of the regular gift that should be transferred
	OwnedGiftId string `json:"owned_gift_id"`
	// REQUIRED:
	// Unique identifier of the chat which will own the gift.
	// The chat must be active in the last 24 hours.
	NewOwnerChatId string `json:"new_owner_chat_id"`

	// The amount of Telegram Stars that will be paid for the transfer from the business account balance.
	// If positive, then the can_transfer_stars business bot right is required.
	StarCount *int `json:"star_count,omitempty"`
}

func (r TransferGift) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if r.OwnedGiftId == "" {
		err = append(err, fmt.Errorf("owned_gift_id can't be empty"))
	}
	if r.NewOwnerChatId == "" {
		err = append(err, fmt.Errorf("new_owner_chat_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s TransferGift) Endpoint() string {
	return "transferGift"
}

func (s TransferGift) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s TransferGift) ContentType() string {
	return "application/json"
}

// Posts a story on behalf of a managed business account.
// Requires the can_manage_stories business bot right.
// Returns [objects.Story] on success.
type PostStory struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Content of the story
	Content objects.InputStoryContent `json:"content"`
	// REQUIRED:
	// Period after which the story is moved to the archive, in seconds;
	// must be one of 6 * 3600, 12 * 3600, 86400, or 2 * 86400
	ActivePeriod int `json:"active_period"`

	// Caption of the story, 0-2048 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the story caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// A JSON-serialized list of clickable areas to be shown on the story
	Areas *[]objects.StoryArea `json:"areas,omitempty"`
	// Pass True to keep the story accessible after it expires
	PostToChatPage *bool `json:"post_to_chat_page,omitempty"`
	// Pass True if the content of the story must be protected from forwarding and screenshotting
	ProtectContent *bool `json:"protect_content,omitempty"`

	contentType string
}

func (r PostStory) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if e := r.Content.Validate(); e != nil {
		err = append(err, e)
	}
	if r.ActivePeriod != 6*3600 && r.ActivePeriod != 12*3600 && r.ActivePeriod != 86400 && r.ActivePeriod != 2*86400 {
		err = append(err, fmt.Errorf("active_period must be one of 6 *3600, 12 * 3600, 86400, or 2 * 86400"))
	}
	if r.CaptionEntities != nil {
		for _, ent := range *r.CaptionEntities {
			if e := ent.Validate(); e != nil {
				err = append(err, e)
				break
			}
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s PostStory) Endpoint() string {
	return "postStory"
}

func (s *PostStory) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("business_connection_id", s.BusinessConnectionId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := gotely.WriteJSONToForm(mw, "content", s.Content); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := s.Content.WriteTo(mw); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := mw.WriteField("active_period", fmt.Sprint(s.ActivePeriod)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", *s.Caption); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", *s.ParseMode); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.CaptionEntities != nil {
			if err := gotely.WriteJSONToForm(mw, "caption_entities", *s.CaptionEntities); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Areas != nil {
			if err := gotely.WriteJSONToForm(mw, "areas", *s.Areas); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.PostToChatPage != nil {
			if err := mw.WriteField("post_to_charge", fmt.Sprint(&s.PostToChatPage)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(&s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s PostStory) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Edits a story previously posted by the bot on behalf of a managed business account.
// Requires the can_manage_stories business bot right.
// Returns [objects.Story] on success.
type EditStory struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// REQUIRED:
	// Content of the story
	Content objects.InputStoryContent `json:"content"`
	// Unique identifier of the story to edit
	StoryId int `json:"story_id"`

	// Caption of the story, 0-2048 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the story caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// A JSON-serialized list of clickable areas to be shown on the story
	Areas *[]objects.StoryArea `json:"areas,omitempty"`

	contentType string
}

func (r EditStory) Validate() error {
	var err gotely.ErrFailedValidation
	if r.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if e := r.Content.Validate(); e != nil {
		err = append(err, e)
	}
	if r.CaptionEntities != nil {
		for _, ent := range *r.CaptionEntities {
			if e := ent.Validate(); e != nil {
				err = append(err, e)
				break
			}
		}
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (s EditStory) Endpoint() string {
	return "editStory"
}

func (s *EditStory) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("business_connection_id", s.BusinessConnectionId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := gotely.WriteJSONToForm(mw, "content", s.Content); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := s.Content.WriteTo(mw); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := mw.WriteField("story_id", fmt.Sprint(s.StoryId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", *s.Caption); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", *s.ParseMode); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.CaptionEntities != nil {
			if err := gotely.WriteJSONToForm(mw, "caption_entities", *s.CaptionEntities); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Areas != nil {
			if err := gotely.WriteJSONToForm(mw, "areas", *s.Areas); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s EditStory) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Deletes a story previously posted by the bot on behalf of a managed business account.
// Requires the can_manage_stories business bot right.
// Returns True on success.
type DeleteStory struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Unique identifier of the story to delete
	StoryId int `json:"story_id"`
}

func (d DeleteStory) Validate() error {
	var err gotely.ErrFailedValidation
	if d.BusinessConnectionId == "" {
		err = append(err, fmt.Errorf("business_connection_id can't be empty"))
	}
	if d.StoryId == 0 {
		err = append(err, fmt.Errorf("story_id can't be empty"))
	}
	if len(err) > 0 {
		return err
	}
	return nil
}

func (d DeleteStory) Endpoint() string {
	return "deleteStory"
}

func (d DeleteStory) Reader() io.Reader {
	return gotely.EncodeJSON(d)
}

func (d DeleteStory) ContentType() string {
	return "application/json"
}
