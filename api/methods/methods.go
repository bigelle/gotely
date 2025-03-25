package methods

import (
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/bigelle/gotely"
	"github.com/bigelle/gotely/api/objects"
)

// A simple method for testing your bot's authentication token.
// Requires no parameters.
// Returns basic information about the bot in form of a [objects.User] object.
type GetMe struct{}

func (g GetMe) Validate() error {
	return nil
}

func (g GetMe) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g GetMe) Endpoint() string {
	return "getMe"
}

func (g GetMe) ContentType() string {
	return "application/json"
}

// Use this method to log out from the cloud Bot API server before launching the bot locally.
// You must log out the bot before running it locally, otherwise there is no guarantee that the bot will receive updates.
// After a successful call, you can immediately log in on a local server,
// but will not be able to log in back to the cloud Bot API server for 10 minutes.
// Returns True on success. Requires no parameters.
type LogOut struct{}

func (g LogOut) Validate() error {
	return nil
}

func (g LogOut) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g LogOut) Endpoint() string {
	return "logOut"
}

func (g LogOut) ContentType() string {
	return "application/json"
}

// Use this method to close the bot instance before moving it from one local server to another.
// You need to delete the webhook before calling this method to ensure that the bot isn't launched again after server restart.
// The method will return error 429 in the first 10 minutes after the bot is launched.
// Returns True on success. Requires no parameters.
type Close struct{}

func (g Close) Validate() error {
	return nil
}

func (g Close) Reader() io.Reader {
	return gotely.EncodeJSON(g)
}

func (g Close) Endpoint() string {
	return "close"
}

func (g Close) ContentType() string {
	return "application/json"
}

// Use this method to send text messages.
// On success, the sent [objects.Message] is returned.
type SendMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty,"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty,"`
	// Mode for parsing entities in the message text.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty,"`
	// A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities *[]objects.MessageEntity `json:"entities,omitempty,"`
	// Link preview generation options for the message
	LinkPreviewOptions *objects.LinkPreviewOptions `json:"link_preview_options,omitempty,"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty,"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty,"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty,"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty,"`
	// Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty,"`
}

func (s SendMessage) Validate() error {
	l := len(s.Text)
	if l < 1 || l > 4096 {
		return fmt.Errorf("text parameter must be between 1 and 4096 characters")
	}
	if s.ChatId == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.LinkPreviewOptions != nil {
		if err := s.LinkPreviewOptions.Validate(); err != nil {
			return err
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMessage) Endpoint() string {
	return "sendMessage"
}

func (s SendMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendMessage) ContentType() string {
	return "application/json"
}

// Use this method to forward messages of any kind.
// Service messages and messages with protected content can't be forwarded.
// On success, the sent [objects.Message] is returned.
type ForwardMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername
	FromChatId string `json:"from_chat_id"`

	// Message identifier in the chat specified in from_chat_id
	MessageId int `json:"message_id"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
}

func (f ForwardMessage) Validate() error {
	if strings.TrimSpace(f.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(f.FromChatId) == "" {
		return fmt.Errorf("from_chat_id parameter can't be empty")
	}
	if f.MessageId < 1 {
		return fmt.Errorf("message_id parameter can't be empty")
	}
	return nil
}

func (s ForwardMessage) Endpoint() string {
	return "forwardMessage"
}

func (s ForwardMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ForwardMessage) ContentType() string {
	return "application/json"
}

// Use this method to forward multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
// On success, an array of [objects.MessageId] of the sent messages is returned.
type ForwardMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`
	// REQUIRED:
	// A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward.
	// The identifiers must be specified in a strictly increasing order.
	MessageIds []int `json:"message_ids"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
}

func (f ForwardMessages) Validate() error {
	if strings.TrimSpace(f.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(f.FromChatId) == "" {
		return fmt.Errorf("from_chat_id parameter can't be empty")
	}
	if len(f.MessageIds) < 1 {
		return fmt.Errorf("message_ids parameter can't be empty")
	}
	return nil
}

func (s ForwardMessages) Endpoint() string {
	return "forwardMessages"
}

func (s ForwardMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ForwardMessages) ContentType() string {
	return "application/json"
}

// Use this method to copy messages of any kind.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message.
// Returns the [objects.MessageId] of the sent message on success.
type CopyMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`
	// REQUIRED:
	// Message identifier in the chat specified in from_chat_id
	MessageId int `json:"message_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the new caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options.
	// A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (c CopyMessage) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.FromChatId) == "" {
		return fmt.Errorf("from_chat_id parameter can't be empty")
	}
	if c.MessageId < 1 {
		return fmt.Errorf("message_ids parameter can't be empty")
	}
	if c.ReplyParameters != nil {
		if err := c.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s CopyMessage) Endpoint() string {
	return "copyMessage"
}

func (s CopyMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CopyMessage) ContentType() string {
	return "application/json"
}

// Use this method to copy messages of any kind.
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages.
// On success, an array of [objects.MessageId] of the sent messages is returned.
type CopyMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy.
	// The identifiers must be specified in a strictly increasing order.
	MessageIds []int `json:"message_ids"`
	// REQUIRED:
	// Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`

	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to copy the messages without their captions
	RemoveCaption *bool `json:"remove_caption,omitempty"`
}

func (c CopyMessages) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.FromChatId) == "" {
		return fmt.Errorf("from_chat_id parameter can't be empty")
	}
	if len(c.MessageIds) < 1 {
		return fmt.Errorf("message_ids parameter can't be empty")
	}
	return nil
}

func (s CopyMessages) Endpoint() string {
	return "copyMessages"
}

func (s CopyMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CopyMessages) ContentType() string {
	return "application/json"
}

// Use this method to send photos.
// On success, the sent [objects.Message] is returned.
type SendPhoto struct {
	//REQUIRED:
	//Unique identifier for the target chat or username of the target channel
	//(in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a photo from the Internet,
	// or upload a new photo using multipart/form-data.
	// The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total.
	// Width and height ratio must be at most 20.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Photo objects.InputFile `json:"photo"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the photo caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard,
	// custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendPhoto) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendPhoto) Endpoint() string {
	return "sendPhoto"
}

func (s *SendPhoto) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if err := mw.WriteField("chat_id", s.ChatId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if err := s.Photo.WriteTo(mw, "photo"); err != nil {
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
		if s.ShowCaptionAboveMedia != nil {
			if err := mw.WriteField("show_caption_above_media", fmt.Sprint(s.ShowCaptionAboveMedia)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.HasSpoiler != nil {
			if err := mw.WriteField("has_spoiler", fmt.Sprint(s.HasSpoiler)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", *s.MessageEffectId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

func (s SendPhoto) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data" // may not work
	}
	return s.contentType
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
// On success, the sent [objects.Message] is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the [SendVoice] method instead.
type SendAudio struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an audio file from the Internet,
	// or upload a new one using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Audio objects.InputFile `json:"audio"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Audio caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the audio caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Duration of the audio in seconds
	Duration *int `json:"duration,omitempty"`
	// Performer
	Performer *string `json:"performer,omitempty"`
	// Track name
	Title *string `json:"title,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendAudio) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Audio.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendAudio) Endpoint() string {
	return "sendAudio"
}

func (s *SendAudio) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
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
		if s.Duration != nil {
			if err := mw.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Title != nil {
			if err := mw.WriteField("title", fmt.Sprint(*s.Title)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Performer != nil {
			if err := mw.WriteField("performer", fmt.Sprint(*s.Performer)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.Audio.WriteTo(mw, "audio"); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Thumbnail != nil {
			if err := s.Thumbnail.WriteTo(mw, "thumbnail"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendAudio) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data" // may not work
	}
	return s.contentType
}

// Use this method to send general files.
// On success, the sent [objects.Message] is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Document objects.InputFile `json:"document"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	// Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the document caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendDocument) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}

	if err := s.Document.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendDocument) Endpoint() string {
	return "sendDocument"
}

func (s *SendDocument) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
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
		if s.DisableContentTypeDetection != nil {
			if err := mw.WriteField("disable_content_type_detection", fmt.Sprint(*s.DisableContentTypeDetection)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.Document.WriteTo(mw, "document"); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Thumbnail != nil {
			if err := s.Thumbnail.WriteTo(mw, "thumbnail"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendDocument) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document).
// On success, the sent [objects.Message] is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Video objects.InputFile `json:"video"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Duration of sent video in seconds
	Duration *int `json:"duration,omitempty"`
	// Video width
	Width *int `json:"width,omitempty"`
	// Video height
	Height *int `json:"height,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	// Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the video caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	// Pass True if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendVideo) Endpoint() string {
	return "sendVideo"
}

func (s *SendVideo) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Duration != nil {
			if err := mw.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Height != nil {
			if err := mw.WriteField("height", fmt.Sprint(*s.Height)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Width != nil {
			if err := mw.WriteField("width", fmt.Sprint(*s.Width)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
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
		if s.HasSpoiler != nil {
			if err := mw.WriteField("has_spoiler", fmt.Sprint(*s.HasSpoiler)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.SupportsStreaming != nil {
			if err := mw.WriteField("supports_streaming", fmt.Sprint(*s.SupportsStreaming)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.Video.WriteTo(mw, "video"); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Thumbnail != nil {
			if err := s.Thumbnail.WriteTo(mw, "thumbnail"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendVideo) Validate() error {
	if s.ChatId == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Video.Validate(); err != nil {
		return err
	}
	if s.Thumbnail != nil {
		if err := s.Thumbnail.Validate(); err != nil {
			return err
		}
	}
	if s.Caption != nil {
		if len(*s.Caption) > 1024 {
			return fmt.Errorf("caption must not be longer than 1024 characters if specified")
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendVideo) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
// On success, the sent [objects.Message] is returned.
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Animation objects.InputFile `json:"animation"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Duration of sent animation in seconds
	Duration *int `json:"duration,omitempty"`
	// Animation width
	Width *int `json:"width,omitempty"`
	// Animation height
	Height *int `json:"height,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	// so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	// Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the animation caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendAnimation) Endpoint() string {
	return "sendAnimation"
}

func (s SendAnimation) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Animation.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SendAnimation) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Duration != nil {
			if err := mw.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Height != nil {
			if err := mw.WriteField("height", fmt.Sprint(*s.Height)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Width != nil {
			if err := mw.WriteField("width", fmt.Sprint(*s.Width)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
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
		if s.HasSpoiler != nil {
			if err := mw.WriteField("has_spoiler", fmt.Sprint(*s.HasSpoiler)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.Animation.WriteTo(mw, "animation"); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Thumbnail != nil {
			if err := s.Thumbnail.WriteTo(mw, "thumbnail"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendAnimation) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format
// (other formats may be sent as Audio or Document).
// On success, the sent [objects.Message] is returned. Bots can currently send voice messages of up to 50 MB in size,
// this limit may be changed in the future.
type SendVoice struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Voice objects.InputFile `json:"voice"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Voice message caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the voice message caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities"`
	// Duration of the voice message in seconds
	Duration *int `json:"duration,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendVoice) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Voice.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendVoice) Endpoint() string {
	return "sendVoice"
}

func (s *SendVoice) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Duration != nil {
			if err := mw.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Caption != nil {
			if err := mw.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ParseMode != nil {
			if err := mw.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
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
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.Voice.WriteTo(mw, "voice"); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()
	return pr
}

func (s SendVoice) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
// Use this method to send video messages.
// On success, the sent [objects.Message] is returned.
type SendVideoNote struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or
	// upload a new video using multipart/form-data.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files.
	// Sending video notes by a URL is currently unsupported
	VideoNote objects.InputFile `json:"video_note"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	// Duration of sent video in seconds
	Duration *int `json:"duration,omitempty"`
	// Video width and height, i.e. diameter of the video message
	Length *int `json:"length,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	// the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
	contentType string
}

func (s SendVideoNote) Endpoint() string {
	return "sendVideoNote"
}

func (s SendVideoNote) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.VideoNote.Validate(); err != nil {
		return err
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SendVideoNote) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Duration != nil {
			if err := mw.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Length != nil {
			if err := mw.WriteField("length", fmt.Sprint(*s.Duration)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

		if err := s.VideoNote.WriteTo(mw, "video_note"); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.Thumbnail != nil {
			if err := s.Thumbnail.WriteTo(mw, "thumbnail"); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendVideoNote) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send paid media.
// On success, the sent Message is returned.
type SendPaidMedia struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	// If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance.
	// Otherwise, they will be credited to the bot's balance.
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	StarCount int `json:"star_count"`
	// REQUIRED:
	// An array describing the media to be sent; up to 10 items
	Media []objects.InputPaidMedia `json:"media"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload *string `json:"payload,omitempty"`
	// Media caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	// Mode for parsing entities in the media caption.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	// Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`

	contentType string
}

func (s SendPaidMedia) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.StarCount < 1 || s.StarCount > 2500 {
		return fmt.Errorf("star_count parameter must be between 1 and 2500")
	}
	if len(s.Media) < 1 {
		return fmt.Errorf("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return fmt.Errorf("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendPaidMedia) Endpoint() string {
	return "sendPaidMedia"
}

func (s SendPaidMedia) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if err := mw.WriteField("chat_id", fmt.Sprint(s.ChatId)); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := mw.WriteField("star_count", fmt.Sprint(s.StarCount)); err != nil {
			pw.CloseWithError(err)
			return
		}
		for _, media := range s.Media {
			if err := media.WriteTo(mw); err != nil {
				pw.CloseWithError(err)
				return
			}
		}

		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.Payload != nil {
			if err := mw.WriteField("payload", *s.Payload); err != nil {
				pw.CloseWithError(err)
				return
			}
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
		if s.ShowCaptionAboveMedia != nil {
			if err := mw.WriteField("show_caption_above_media", fmt.Sprint(*s.ShowCaptionAboveMedia)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
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

func (s SendPaidMedia) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of [objects.Messages] that were sent is returned.
type SendMediaGroup struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// An array describing messages to be sent, must include 2-10 items
	Media []objects.InputMedia `json:"media"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// Sends messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	contentType     string
}

func (s SendMediaGroup) Endpoint() string {
	return "sendMediaGroup"
}

func (s SendMediaGroup) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if len(s.Media) < 1 {
		return fmt.Errorf("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return fmt.Errorf("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SendMediaGroup) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if s.BusinessConnectionId != nil {
			if err := mw.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if err := mw.WriteField("chat_id", s.ChatId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.MessageThreadId != nil {
			if err := mw.WriteField("message_thread_id", *s.MessageThreadId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if err := gotely.WriteJSONToForm(mw, "media", s.Media); err != nil {
			pw.CloseWithError(err)
			return
		}
		if s.DisableNotification != nil {
			if err := mw.WriteField("disable_notification", fmt.Sprint(s.DisableNotification)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ProtectContent != nil {
			if err := mw.WriteField("protect_content", fmt.Sprint(s.ProtectContent)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.AllowPaidBroadcast != nil {
			if err := mw.WriteField("allow_paid_broadcast", fmt.Sprint(s.AllowPaidBroadcast)); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.MessageEffectId != nil {
			if err := mw.WriteField("message_effect_id", *s.MessageEffectId); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
		if s.ReplyParameters != nil {
			if err := gotely.WriteJSONToForm(mw, "reply_parameters", *s.ReplyParameters); err != nil {
				pw.CloseWithError(err)
				return
			}
		}

		for _, media := range s.Media {
			if err := media.WriteTo(mw); err != nil {
				pw.CloseWithError(err)
				return
			}
		}
	}()
	return pr
}

func (s SendMediaGroup) ContentType() string {
	if s.contentType == "" {
		return "multipart/form-data"
	}
	return s.contentType
}

// Use this method to send point on the map.
// On success, the sent [objects.Message] is returned.
type SendLocation struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Latitude of the location
	Latitude *float64 `json:"latitude"`
	// REQUIRED:
	// Longitude of the location
	Longitude *float64 `json:"longitude"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
	//Period in seconds during which the location will be updated
	//(see https://telegram.org/blog/live-locations), should be between 60 and 86400,
	//or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int `json:"live_period,omitempty"`
	// For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int `json:"heading,omitempty"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int `json:"proximity_alert_radius,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendLocation) Endpoint() string {
	return "sendLocation"
}

func (s SendLocation) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.Latitude == nil {
		return fmt.Errorf("latitude parameter can't be empty")
	}
	if s.Longitude == nil {
		return fmt.Errorf("longitude parameter can't be empty")
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendLocation) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendLocation) ContentType() string {
	return "application/json"
}

// Use this method to send information about a venue.
// On success, the sent [objects.Message] is returned.
type SendVenue struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Latitude of the venue
	Latitude *float64 `json:"latitude"`
	// REQUIRED:
	// Longitude of the venue
	Longitude *float64 `json:"longitude"`
	// REQUIRED:
	// Name of the venue
	Title string `json:"title"`
	// REQUIRED:
	// Address of the venue
	Address string `json:"address"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// Foursquare identifier of the venue
	FoursquareId *string `json:"foursquare_id,omitempty"`
	// Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	// Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	// Google Places type of the venue. (See https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendVenue) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.Latitude == nil {
		return fmt.Errorf("latitude parameter can't be empty")
	}
	if s.Longitude == nil {
		return fmt.Errorf("longitude parameter can't be empty")
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendVenue) Endpoint() string {
	return "sendVenue"
}

func (s SendVenue) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendVenue) ContentType() string {
	return "application/json"
}

// Use this method to send phone contacts.
// On success, the sent [objects.Message] is returned.
type SendContact struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// REQUIRED:
	// Contact's first name
	FirstName string `json:"first_name"`

	// Contact's last name
	LastName *string `json:"last_name,omitempty"`
	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard *string `json:"vcard,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendContact) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.PhoneNumber) == "" {
		return fmt.Errorf("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return fmt.Errorf("first_name parameter can't be empty")
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendContact) Endpoint() string {
	return "sendContact"
}

func (s SendContact) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendContact) ContentType() string {
	return "application/json"
}

// Use this method to send a native poll.
// On success, the sent [objects.Message] is returned.
type SendPoll struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// REQUIRED:
	// A JSON-serialized list of 2-10 answer options
	Options []objects.InputPollOption `json:"options"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// Mode for parsing entities in the question.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	// Currently, only custom emoji entities are allowed
	QuestionParseMode *string `json:"question_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the poll question.
	// It can be specified instead of question_parse_mode
	QuestionEntities *[]objects.MessageEntity `json:"question_entities,omitempty"`
	// True, if the poll needs to be anonymous, defaults to True
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type *string `json:"type,omitempty"`
	// True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	AllowMultipleAnswers *bool `json:"allow_multiple_answers,omitempty"`
	// 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionId *int `json:"correct_option_id,omitempty"`
	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll,
	// 0-200 characters with at most 2 line feeds after entities parsing
	Explanation *string `json:"explanation,omitempty"`
	// Mode for parsing entities in the explanation.
	// See https://core.telegram.org/bots/api#formatting-options for more details.
	ExplanationParseMode *string `json:"explanation_parse_mode,omitempty"`
	// A JSON-serialized list of special entities that appear in the poll explanation.
	// It can be specified instead of explanation_parse_mode
	ExplanationEntities *[]objects.MessageEntity `json:"explanation_entities,omitempty"`
	// Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	OpenPeriod *int `json:"open_period,omitempty"`
	// Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future.
	// Can't be used together with open_period.
	CloseDate *int `json:"close_date,omitempty"`
	// Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed *bool `json:"is_closed,omitempty"`
	// /Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendPoll) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Question) == "" {
		return fmt.Errorf("question parameter can't be empty")
	}
	for _, opt := range s.Options {
		if err := opt.Validate(); err != nil {
			return err
		}
	}
	if len(s.Options) < 2 || len(s.Options) > 10 {
		return fmt.Errorf("options parameter must be between 2 and 10")
	}
	if s.Type != nil {
		if *s.Type != "quiz" && *s.Type != "regular" {
			return fmt.Errorf("type parameter must be 'regular' or 'quiz' if specified")
		}
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendPoll) Endpoint() string {
	return "sendPoll"
}

func (s SendPoll) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendPoll) ContentType() string {
	return "application/json"
}

// Use this method to send an animated emoji that will display a random value.
// On success, the sent [objects.Message] is returned.
type SendDice struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.
	// Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”.
	// Defaults to “🎲”
	Emoji string `json:"emoji"`

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	// Protects the contents of the sent message from forwarding
	ProtectContent *bool `json:"protect_content,omitempty"`
	// Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	// The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	// Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	// Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendDice) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Emoji) == "" {
		return fmt.Errorf("emoji parameter can't be empty")
	}
	if s.ReplyParameters != nil {
		if err := s.ReplyParameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendDice) Endpoint() string {
	return "sendDice"
}

func (s SendDice) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendDice) ContentType() string {
	return "application/json"
}

// Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status).
// Returns True on success.
//
// Example: The ImageBot needs some time to process a request and upload the image.
// Instead of sending a text message along the lines of “Retrieving image, please wait…”,
// the bot may use sendChatAction with action = upload_photo. The user will see a “sending photo” status for the bot.
//
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
type SendChatAction struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread; for supergroups only
	Action string `json:"action"`

	// Unique identifier of the business connection on behalf of which the action will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Type of action to broadcast. Choose one, depending on what the user is about to receive: 'typing' for text messages,
	//'upload_photo' for photos, 'record_video' or 'upload_video' for videos, 'record_voice' or 'upload_voice' for voice notes,
	//'upload_document' for general files, 'choose_sticker' for stickers, 'find_location' for location data,
	//'record_video_note' or 'upload_video_note' for video notes.
	MessageThreadId *string `json:"message_thread_id,omitempty"`
}

func (s SendChatAction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Action) == "" {
		return fmt.Errorf("action parameter can't be empty")
	}
	allowed := map[string]struct{}{
		"typing":            {},
		"upload_photo":      {},
		"record_video":      {},
		"upload_video":      {},
		"record_voice":      {},
		"upload_voice":      {},
		"upload_document":   {},
		"choose_sticker":    {},
		"find_location":     {},
		"record_video_note": {},
		"upload_video_note": {},
	}
	if _, ok := allowed[s.Action]; !ok {
		return fmt.Errorf("invalid action parameter: see https://core.telegram.org/bots/api#sendchataction for a list of available actions")
	}
	return nil
}

func (s SendChatAction) Endpoint() string {
	return "sendChatAction"
}

func (s SendChatAction) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SendChatAction) ContentType() string {
	return "application/json"
}

// Use this method to change the chosen reactions on a message.
// Service messages can't be reacted to.
// Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel.
// Bots can't use paid reactions.
// Returns True on success.
type SetMessageReaction struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	MessageId int `json:"message_id"`

	// A JSON-serialized list of reaction types to set on the message.
	// Currently, as non-premium users, bots can set up to one reaction per message.
	// A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators.
	// Paid reactions can't be used by bots.
	Reaction *[]objects.ReactionType `json:"reaction,omitempty"`
	// Pass True to set the reaction with a big animation
	IsBig *bool `json:"is_big,omitempty"`
}

func (s SetMessageReaction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.MessageId < 1 {
		return fmt.Errorf("message_id parameter can't be empty")
	}
	for _, r := range *s.Reaction {
		if err := r.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetMessageReaction) Endpoint() string {
	return "setMessageReaction"
}

func (s SetMessageReaction) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMessageReaction) ContentType() string {
	return "application/json"
}

// Use this method to get a list of profile pictures for a user.
// Returns a [objects.UserProfilePhotos] object.
type GetUserProfilePhotos struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset *int `json:"offset,omitempty"`
	// Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit *int `json:"limit,omitempty"`
}

func (g GetUserProfilePhotos) Validate() error {
	if g.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return fmt.Errorf("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (s GetUserProfilePhotos) Endpoint() string {
	return "getUserProfilePhotos"
}

func (s GetUserProfilePhotos) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetUserProfilePhotos) ContentType() string {
	return "application/json"
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the
// Mini App method requestEmojiStatusAccess.
// Returns True on success.
type SetUserEmojiStatus struct {
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`

	// Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusCustomEmojiId *string `json:"emoji_status_custom_emoji_id,omitempty"`
	// Expiration date of the emoji status, if any
	EmojiStatusExpirationDate *int `json:"emoji_status_expiration_date,omitempty"`
}

func (s SetUserEmojiStatus) Validate() error {
	if s.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s SetUserEmojiStatus) Endpoint() string {
	return "setUserEmojiStatus"
}

func (s SetUserEmojiStatus) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetUserEmojiStatus) ContentType() string {
	return "application/json"
}

// Use this method to get basic information about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size.
// On success, a [objects.File] object is returned.
// The file can then be downloaded via the link https://gotely.telegram.org/file/bot<token>/<file_path>,
// where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type.
// You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	// REQUIRED:
	// File identifier to get information about
	FileId string `json:"file_id"`
}

func (g GetFile) Validate() error {
	if strings.TrimSpace(g.FileId) == "" {
		return fmt.Errorf("file_id parameter can't be empty")
	}
	return nil
}

func (s GetFile) Endpoint() string {
	return "getFile"
}

func (s GetFile) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetFile) ContentType() string {
	return "application/json"
}

// Use this method to ban a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links,
// etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type BanChatMember struct {
	// REQUIRED:
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`

	// Date when the user will be unbanned; Unix time.
	// If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever.
	// Applied for supergroups and channels only.
	UntilDate *int `json:"until_date,omitempty"`
	// Pass True to delete all messages from the chat for the user that is being removed.
	// If False, the user will be able to see messages in the group that were sent before the user was removed.
	// Always True for supergroups and channels.
	RevokeMessages *bool `json:"revoke_messages,omitempty"`
}

func (b BanChatMember) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if b.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s BanChatMember) Endpoint() string {
	return "banChatMember"
}

func (s BanChatMember) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s BanChatMember) ContentType() string {
	return "application/json"
}

// Use this method to unban a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work. By default,
// this method guarantees that after the call the user is not a member of the chat, but will be able to join it.
// So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned.
// Returns True on success.
type UnbanChatMember struct {
	// REQUIRED:
	// Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`

	// Do nothing if the user is not banned
	OnlyIfBanned *bool `json:"only_if_banned,omitempty"`
}

func (b UnbanChatMember) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if b.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s UnbanChatMember) Endpoint() string {
	return "unbanChatMember"
}

func (s UnbanChatMember) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnbanChatMember) ContentType() string {
	return "application/json"
}

// Use this method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights.
// Pass True for all permissions to lift restrictions from a user.
// Returns True on success.
type RestrictChatMember struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// REQUIRED:
	// A JSON-serialized object for new user permissions
	Permissions objects.ChatPermissions `json:"permissions"`

	// Pass True if chat permissions are set independently.
	// Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	// can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	// the can_send_polls permission will imply the can_send_messages permission.
	UserIndependentChatPermissions *bool `json:"user_independent_chat_permissions,omitempty"`
	// Date when restrictions will be lifted for the user; Unix time.
	// If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilDate *int `json:"until_date,omitempty"`
}

func (r RestrictChatMember) Validate() error {
	if strings.TrimSpace(r.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if r.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s RestrictChatMember) Endpoint() string {
	return "restrictChatMember"
}

func (s RestrictChatMember) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s RestrictChatMember) ContentType() string {
	return "application/json"
}

// Use this method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Pass False for all boolean parameters to demote a user.
// Returns True on success
type PromoteChatMember struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`

	// Pass True if the administrator's presence in the chat is hidden
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	// Pass True if the administrator can access the chat event log, get boost list,
	// see hidden supergroup and channel members, report spam messages and ignore slow mode.
	// Implied by any other administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty"`
	// Pass True if the administrator can delete messages of other users
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`
	// Pass True if the administrator can manage video chats
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`
	// Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`
	// Pass True if the administrator can add new administrators with a subset of their own privileges or
	// demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`
	// Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty"`
	// Pass True if the administrator can invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`
	// Pass True if the administrator can post stories to the chat
	CanPostStories *bool `json:"can_post_stories,omitempty"`
	// Pass True if the administrator can edit stories posted by other users,
	// post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories *bool `json:"can_edit_stories,omitempty"`
	// Pass True if the administrator can delete stories posted by other users
	CanDeleteStories *bool `json:"can_delete_stories,omitempty"`
	// Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	// Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	// Pass True if the administrator can pin messages; for supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	// Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

func (p PromoteChatMember) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if p.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s PromoteChatMember) Endpoint() string {
	return "promoteChatMember"
}

func (s PromoteChatMember) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s PromoteChatMember) ContentType() string {
	return "application/json"
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot.
// Returns True on success.
type SetChatAdministratorCustomTitle struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
	// REQUIRED:
	// New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title"`
}

func (s SetChatAdministratorCustomTitle) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	if len(s.CustomTitle) > 16 {
		return fmt.Errorf("custom_title parameter must be not longer than 16 characters")
	}
	for _, r := range s.CustomTitle {
		if (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
			(r >= 0x1F300 && r <= 0x1F5FF) || // Miscellaneous Symbols and Pictographs
			(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
			(r >= 0x1F700 && r <= 0x1F77F) { // Alchemical Symbols
			return fmt.Errorf("invalid custom_title parameter: emojis are not allowed")
		}
	}
	return nil
}

func (s SetChatAdministratorCustomTitle) Endpoint() string {
	return "setChatAdministratorCustomTitle"
}

func (s SetChatAdministratorCustomTitle) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatAdministratorCustomTitle) ContentType() string {
	return "application/json"
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned,
// the owner of the banned chat won't be able to send messages on behalf of any of their channels.
// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights.
// Returns True on success.
type BanChatSenderChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target sender chat
	SenderChatId int `json:"sender_chat_id"`
}

func (b BanChatSenderChat) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if b.SenderChatId < 1 {
		return fmt.Errorf("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (s BanChatSenderChat) Endpoint() string {
	return "banChatSenderChat"
}

func (s BanChatSenderChat) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s BanChatSenderChat) ContentType() string {
	return "application/json"
}

// Use this method to unban a previously banned channel chat in a supergroup or channel.
// The bot must be an administrator for this to work and must have the appropriate administrator rights.
// Returns True on success.
type UnbanChatSenderChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target sender chat
	SenderChatId int `json:"sender_chat_id"`
}

func (b UnbanChatSenderChat) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if b.SenderChatId < 1 {
		return fmt.Errorf("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (s UnbanChatSenderChat) Endpoint() string {
	return "unbanChatSenderChat"
}

func (s UnbanChatSenderChat) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnbanChatSenderChat) ContentType() string {
	return "application/json"
}

// Use this method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights.
// Returns True on success.
type SetChatPermissions struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// A JSON-serialized object for new default chat permissions
	Permissions objects.ChatPermissions `json:"permissions"`

	// Pass True if chat permissions are set independently.
	// Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the
	// can_send_messages, can_send_audios, can_send_documents, can_send_photos,
	// can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	// the can_send_polls permission will imply the can_send_messages permission.
	UserIndependentChatPermissions *bool `json:"user_independent_chat_permissions"`
}

func (s SetChatPermissions) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s SetChatPermissions) Endpoint() string {
	return "setChatPermissions"
}

func (s SetChatPermissions) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatPermissions) ContentType() string {
	return "application/json"
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators.
// If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method.
// If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLink struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (e ExportChatInviteLink) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s ExportChatInviteLink) Endpoint() string {
	return "exportChatInviteLink"
}

func (s ExportChatInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ExportChatInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to create an additional invite link for a chat.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method [RevokeChatInviteLink].
// Returns the new invite link as [objects.ChatInviteLink] object.
type CreateInviteLink struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`

	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	// Point in time (Unix timestamp) when the link will expire
	ExpireDate *int `json:"expire_date,omitempty"`
	// The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int `json:"member_limit,omitempty"`
	// True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty"`
}

func (c CreateInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return fmt.Errorf("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return fmt.Errorf("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (s CreateInviteLink) Endpoint() string {
	return "createInviteLink"
}

func (s CreateInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CreateInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to edit a non-primary invite link created by the bot.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the edited invite link as a [objects.ChatInviteLink] object.
type EditChatInviteLink struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// The invite link to edit
	InviteLink string `json:"invite_link"`

	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	// Point in time (Unix timestamp) when the link will expire
	ExpireDate *int `json:"expire_date,omitempty"`
	// The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int `json:"member_limit,omitempty"`
	// True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty"`
}

func (c EditChatInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return fmt.Errorf("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return fmt.Errorf("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (s EditChatInviteLink) Endpoint() string {
	return "editChatInviteLink"
}

func (s EditChatInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditChatInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to create a subscription invite link for a channel chat.
// The bot must have the can_invite_users administrator rights.
// The link can be edited using the method [EditChatSubscriptionInviteLink] or revoked using the method [RevokeChatInviteLink].
// Returns the new invite link as a [objects.ChatInviteLink] object.
type CreateChatSubscriptionInviteLink struct {
	// REQUIRED:
	// Unique identifier for the target channel chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPeriod int `json:"subscription_period"`
	// REQUIRED:
	// The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
	SubscriptionPrice int `json:"subscription_price"`

	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
}

func (c CreateChatSubscriptionInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if c.SubscriptionPeriod != 2592000 {
		return fmt.Errorf("subscription_period currently must always be 2592000 seconds (30 days)")
	}
	if c.SubscriptionPrice < 1 || c.SubscriptionPrice > 2500 {
		return fmt.Errorf("subscription_price must be between 1 and 2500")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return fmt.Errorf("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (s CreateChatSubscriptionInviteLink) Endpoint() string {
	return "createChatSubscriptionInviteLink"
}

func (s CreateChatSubscriptionInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CreateChatSubscriptionInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to edit a subscription invite link created by the bot.
// The bot must have the can_invite_users administrator rights.
// Returns the edited invite link as a [objects.ChatInviteLink] object.
type EditChatSubscriptionInviteLink struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// The invite link to edit
	InviteLink string `json:"invite_link"`

	// Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
}

func (c EditChatSubscriptionInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.InviteLink) == "" {
		return fmt.Errorf("invite_link parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return fmt.Errorf("name parameter must not be longer than 32 characters if specified")
		}
	}
	return nil
}

func (s EditChatSubscriptionInviteLink) Endpoint() string {
	return "editChatSubscriptionInviteLink"
}

func (s EditChatSubscriptionInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditChatSubscriptionInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to revoke an invite link created by the bot.
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the revoked invite link as [objects.ChatInviteLink] object.
type RevokeInviteLink struct {
	// Unique identifier of the target chat or username of the target channel (in the format @channelusername)
	// REQUIRED:
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// The invite link to revoke
	InviteLink string `json:"invite_link"`
}

func (c RevokeInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.InviteLink) == "" {
		return fmt.Errorf("invite_link parameter can't be empty")
	}
	return nil
}

func (s RevokeInviteLink) Endpoint() string {
	return "revokeInviteLink"
}

func (s RevokeInviteLink) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s RevokeInviteLink) ContentType() string {
	return "application/json"
}

// Use this method to approve a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
type ApproveChatJoinRequest struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (s ApproveChatJoinRequest) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s ApproveChatJoinRequest) Endpoint() string {
	return "approveChatJoinRequest"
}

func (s ApproveChatJoinRequest) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ApproveChatJoinRequest) ContentType() string {
	return "application/json"
}

// Use this method to decline a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
type DeclineChatJoinRequest struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (s DeclineChatJoinRequest) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s DeclineChatJoinRequest) Endpoint() string {
	return "declineChatJoinRequest"
}

func (s DeclineChatJoinRequest) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeclineChatJoinRequest) ContentType() string {
	return "application/json"
}

// Use this method to set a new profile photo for the chat.
// Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatPhoto struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// New chat photo, uploaded using multipart/form-data
	Photo       objects.InputFile `json:"photo"`
	contentType string
}

func (s SetChatPhoto) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SetChatPhoto) Reader() io.Reader {
	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)
	s.contentType = mw.FormDataContentType()

	go func() {
		defer pw.Close()
		defer mw.Close()

		if _, ok := s.Photo.(objects.InputFileFromRemote); ok {
			pw.CloseWithError(fmt.Errorf("can't use remote file in setChatPhoto method"))
			return
		}

		if err := mw.WriteField("chat_id", s.ChatId); err != nil {
			pw.CloseWithError(err)
			return
		}
		if err := s.Photo.WriteTo(mw, "photo"); err != nil {
			pw.CloseWithError(err)
			return
		}
	}()
	return pr
}

func (s SetChatPhoto) Endpoint() string {
	return "setChatPhoto"
}

func (s SetChatPhoto) ContentType() string {
	return "application/json"
}

// Use this method to delete a chat photo. Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type DeleteChatPhoto struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (d DeleteChatPhoto) Validate() error {
	if strings.TrimSpace(d.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s DeleteChatPhoto) Endpoint() string {
	return "deleteChatPhoto"
}

func (s DeleteChatPhoto) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteChatPhoto) ContentType() string {
	return "application/json"
}

// Use this method to change the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatTitle struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// New chat title, 1-128 characters
	Title string `json:"title"`
}

func (s SetChatTitle) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if len(s.Title) < 1 || len(s.Title) > 128 {
		return fmt.Errorf("title parameter must be between 1 and 128 characters long")
	}
	return nil
}

func (s SetChatTitle) Endpoint() string {
	return "setChatTitle"
}

func (s SetChatTitle) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatTitle) ContentType() string {
	return "application/json"
}

// Use this method to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatDescription struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// New chat description, 0-255 characters
	Description string `json:"description"`
}

func (s SetChatDescription) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if len(s.Description) > 255 {
		return fmt.Errorf("description parameter must not be longer than 255 characters")
	}
	return nil
}

func (s SetChatDescription) Endpoint() string {
	return "setChatDescription"
}

func (s SetChatDescription) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatDescription) ContentType() string {
	return "application/json"
}

// Use this method to add a message to the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type PinChatMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Identifier of a message to pin
	MessageId int `json:"message_id"`

	// Unique identifier of the business connection on behalf of which the message will be pinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Pass True if it is not necessary to send a notification to all chat members about the new pinned message.
	// Notifications are always disabled in channels and private chats.
	DisableNotification *bool `json:"disable_notification,omitempty"`
}

func (p PinChatMessage) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if p.MessageId < 1 {
		return fmt.Errorf("message_id parameter can't be empty")
	}
	return nil
}

func (s PinChatMessage) Endpoint() string {
	return "pinChatMessage"
}

func (s PinChatMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s PinChatMessage) ContentType() string {
	return "application/json"
}

// Use this method to remove a message from the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type UnpinChatMessage struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`

	// Unique identifier of the business connection on behalf of which the message will be unpinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	// Identifier of the message to unpin. Required if business_connection_id is specified.
	// If not specified, the most recent pinned message (by sending date) will be unpinned.
	MessageId *int `json:"message_id,omitempty"`
}

func (p UnpinChatMessage) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if p.MessageId != nil {
		if *p.MessageId < 1 {
			return fmt.Errorf("message_id parameter can't be empty")
		}
	}
	return nil
}

func (s UnpinChatMessage) Endpoint() string {
	return "unpinChatMessage"
}

func (s UnpinChatMessage) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnpinChatMessage) ContentType() string {
	return "application/json"
}

// Use this method to clear the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type UnpinAllChatMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (s UnpinAllChatMessages) Endpoint() string {
	return "unpinAllChatMessages"
}

func (s UnpinAllChatMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnpinAllChatMessages) ContentType() string {
	return "application/json"
}

func (p UnpinAllChatMessages) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

// Use this method for your bot to leave a group, supergroup or channel.
// Returns True on success.
type LeaveChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (p LeaveChat) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s LeaveChat) Endpoint() string {
	return "leaveChat"
}

func (s LeaveChat) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s LeaveChat) ContentType() string {
	return "application/json"
}

// Use this method to get up-to-date information about the chat.
// Returns a [objects.ChatFullInfo] object on success.
type GetChat struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (p GetChat) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s GetChat) Endpoint() string {
	return "getChat"
}

func (s GetChat) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetChat) ContentType() string {
	return "application/json"
}

// Use this method to get a list of administrators in a chat, which aren't bots.
// Returns an Array of [objects.ChatMember] objects.
type GetChatAdministrators struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (p GetChatAdministrators) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s GetChatAdministrators) Endpoint() string {
	return "getChatAdministrators"
}

func (s GetChatAdministrators) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetChatAdministrators) ContentType() string {
	return "application/json"
}

// Use this method to get the number of members in a chat.
// Returns Int on success.
type GetChatMemberCount struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
}

func (p GetChatMemberCount) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s GetChatMemberCount) Endpoint() string {
	return "getChatMemberCount"
}

func (s GetChatMemberCount) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetChatMemberCount) ContentType() string {
	return "application/json"
}

// Use this method to get information about a member of a chat.
// The method is only guaranteed to work for other users if the bot is an administrator in the chat.
// Returns a [objects.ChatMember] object on success.
type GetChatMember struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (p GetChatMember) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if p.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s GetChatMember) Endpoint() string {
	return "getChatMember"
}

func (s GetChatMember) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetChatMember) ContentType() string {
	return "application/json"
}

// Use this method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True on success.
type SetChatStickerSet struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
}

func (p SetChatStickerSet) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(p.StickerSetName) == "" {
		return fmt.Errorf("sticker_set_name parameter can't be empty")
	}
	return nil
}

func (s SetChatStickerSet) Endpoint() string {
	return "setChatStickerSet"
}

func (s SetChatStickerSet) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatStickerSet) ContentType() string {
	return "application/json"
}

// Use this method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True on success.
type DeleteChatStickerSet struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (p DeleteChatStickerSet) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s DeleteChatStickerSet) Endpoint() string {
	return "deleteChatStickerSet"
}

func (s DeleteChatStickerSet) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteChatStickerSet) ContentType() string {
	return "application/json"
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user.
// Requires no parameters. Returns an Array of Sticker objects.
type GetForumTopicIconStickers struct{}

func (g GetForumTopicIconStickers) Validate() error {
	return nil
}

func (s GetForumTopicIconStickers) Endpoint() string {
	return "getForumTopicIconStickers"
}

func (s GetForumTopicIconStickers) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetForumTopicIconStickers) ContentType() string {
	return "application/json"
}

// Use this method to create a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns information about the created topic as a [ForumTopic] object.
type CreateForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Topic name, 1-128 characters
	Name string `json:"name"`

	// Color of the topic icon in RGB format.
	// Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconColor *int `json:"icon_color,omitempty"`
	// Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

func (c CreateForumTopic) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 128 {
		return fmt.Errorf("name parameter must be between 1 and 128 characters long")
	}
	if c.IconColor != nil {
		validIconColors := map[int]struct{}{
			7322096:  {},
			16766590: {},
			13338331: {},
			9367192:  {},
			16749490: {},
			16478047: {},
		}
		if _, ok := validIconColors[*c.IconColor]; !ok {
			return fmt.Errorf("icon_color must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)")
		}
	}
	return nil
}

func (s CreateForumTopic) Endpoint() string {
	return "createForumTopic"
}

func (s CreateForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CreateForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to edit name and icon of a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type EditForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`

	// New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	Name *string `json:"name,omitempty"`
	// New unique identifier of the custom emoji shown as the topic icon.
	// Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	// Pass an empty string to remove the icon. If not specified, the current icon will be kept
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
}

func (e EditForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return fmt.Errorf("message_thread_id parameter can't be empty")
	}
	if e.Name != nil {
		if len(*e.Name) > 128 {
			return fmt.Errorf("name parameter must not be longer than 128 characters")
		}
	}
	return nil
}

func (s EditForumTopic) Endpoint() string {
	return "editForumTopic"
}

func (s EditForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to close an open topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type CloseForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
}

func (e CloseForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return fmt.Errorf("message_thread_id parameter can't be empty")
	}
	return nil
}

func (s CloseForumTopic) Endpoint() string {
	return "closeForumTopic"
}

func (s CloseForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CloseForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to reopen a closed topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type ReopenForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
}

func (e ReopenForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return fmt.Errorf("message_thread_id parameter can't be empty")
	}
	return nil
}

func (s ReopenForumTopic) Endpoint() string {
	return "reopenForumTopic"
}

func (s ReopenForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ReopenForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights.
// Returns True on success.
type DeleteForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
}

func (e DeleteForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return fmt.Errorf("message_thread_id parameter can't be empty")
	}
	return nil
}

func (s DeleteForumTopic) Endpoint() string {
	return "deleteForumTopic"
}

func (s DeleteForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to clear the list of pinned messages in a forum topic.
// The bot must be an administrator in the chat for this to work and
// must have the can_pin_messages administrator right in the supergroup.
// Returns True on success.
type UnpinAllForumTopicMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
}

func (e UnpinAllForumTopicMessages) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return fmt.Errorf("message_thread_id parameter can't be empty")
	}
	return nil
}

func (s UnpinAllForumTopicMessages) Endpoint() string {
	return "unpinAllForumTopicMessages"
}

func (s UnpinAllForumTopicMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnpinAllForumTopicMessages) ContentType() string {
	return "application/json"
}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns True on success.
type EditGeneralForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// New topic name, 1-128 characters
	Name string `json:"name"`
}

func (e EditGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.Name) == "" {
		return fmt.Errorf("name parameter can't be empty")
	}
	return nil
}

func (s EditGeneralForumTopic) Endpoint() string {
	return "editGeneralForumTopic"
}

func (s EditGeneralForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s EditGeneralForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to close an open 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights.
// Returns True on success.
type CloseGeneralForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (e CloseGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s CloseGeneralForumTopic) Endpoint() string {
	return "closeGeneralForumTopic"
}

func (s CloseGeneralForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s CloseGeneralForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// The topic will be automatically unhidden if it was hidden.
// Returns True on success.
type ReopenGeneralForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (e ReopenGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s ReopenGeneralForumTopic) Endpoint() string {
	return "reopenGeneralForumTopic"
}

func (s ReopenGeneralForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s ReopenGeneralForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to hide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// The topic will be automatically closed if it was open.
// Returns True on success.
type HideGeneralForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (e HideGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s HideGeneralForumTopic) Endpoint() string {
	return "hideGeneralForumTopic"
}

func (s HideGeneralForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s HideGeneralForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to unhide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns True on success.
type UnhideGeneralForumTopic struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (e UnhideGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s UnhideGeneralForumTopic) Endpoint() string {
	return "unhideGeneralForumTopic"
}

func (s UnhideGeneralForumTopic) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnhideGeneralForumTopic) ContentType() string {
	return "application/json"
}

// Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and
// must have the can_pin_messages administrator right in the supergroup.
// Returns True on success.
type UnpinAllGeneralForumTopicMessages struct {
	// REQUIRED:
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
}

func (e UnpinAllGeneralForumTopicMessages) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	return nil
}

func (s UnpinAllGeneralForumTopicMessages) Endpoint() string {
	return "unpinAllGeneralForumTopicMessages"
}

func (s UnpinAllGeneralForumTopicMessages) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s UnpinAllGeneralForumTopicMessages) ContentType() string {
	return "application/json"
}

// Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL.
// For this option to work, you must first create a game for your bot via @BotFather and accept the terms.
// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQuery struct {
	// REQUIRED:
	// Unique identifier for the query to be answered
	CallbackQueryId string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text *string `json:"text,omitempty"`
	// If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert *bool `json:"show_alert,omitempty"`
	//URL that will be opened by the user's client.
	//If you have created a Game and accepted the conditions via @BotFather,
	//specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.
	//
	//Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	Url *string `json:"url,omitempty"`
	// The maximum amount of time in seconds that the result of the callback query may be cached client-side.
	// Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime *int `json:"cache_time,omitempty"`
}

func (a AnswerCallbackQuery) Validate() error {
	if strings.TrimSpace(a.CallbackQueryId) == "" {
		return fmt.Errorf("callback_query_id parameter can't be empty")
	}
	if a.Text != nil {
		if len(*a.Text) > 200 {
			return fmt.Errorf("text parameter must not be longer than 200 characters if specified")
		}
	}
	return nil
}

func (s AnswerCallbackQuery) Endpoint() string {
	return "answerCallbackQuery"
}

func (s AnswerCallbackQuery) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s AnswerCallbackQuery) ContentType() string {
	return "application/json"
}

// Use this method to get the list of boosts added to a chat by a user.
// Requires administrator rights in the chat.
// Returns a [objects.UserChatBoosts] object.
type GetUserChatBoosts struct {
	// REQUIRED:
	// Unique identifier for the chat or username of the channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	// REQUIRED:
	// Unique identifier of the target user
	UserId int `json:"user_id"`
}

func (g GetUserChatBoosts) Validate() error {
	if strings.TrimSpace(g.ChatId) == "" {
		return fmt.Errorf("chat_id parameter can't be empty")
	}
	if g.UserId < 1 {
		return fmt.Errorf("user_id parameter can't be empty")
	}
	return nil
}

func (s GetUserChatBoosts) Endpoint() string {
	return "getUserChatBoosts"
}

func (s GetUserChatBoosts) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetUserChatBoosts) ContentType() string {
	return "application/json"
}

// Use this method to get information about the connection of the bot with a business account.
// Returns a [objects.BusinessConnection] object on success.
type GetBusinessConnection struct {
	// REQUIRED:
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
}

func (g GetBusinessConnection) Validate() error {
	if strings.TrimSpace(g.BusinessConnectionId) == "" {
		return fmt.Errorf("business_connection_id parameter can't be empty")
	}
	return nil
}

func (s GetBusinessConnection) Endpoint() string {
	return "getBusinessConnection"
}

func (s GetBusinessConnection) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetBusinessConnection) ContentType() string {
	return "application/json"
}

// Use this method to change the list of the bot's commands.
// See this manual for more details about bot commands.
// Returns True on success.
type SetMyCommands struct {
	// REQUIRED:
	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []objects.BotCommand `json:"commands"`

	// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s SetMyCommands) Validate() error {
	for _, command := range s.Commands {
		if err := command.Validate(); err != nil {
			return err
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetMyCommands) Endpoint() string {
	return "setMyCommands"
}

func (s SetMyCommands) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMyCommands) ContentType() string {
	return "application/json"
}

// Use this method to delete the list of the bot's commands for the given scope and user language.
// After deletion, higher level commands will be shown to affected users.
// Returns True on success.
type DeleteMyCommands struct {
	// A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s DeleteMyCommands) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s DeleteMyCommands) Endpoint() string {
	return "deleteMyCommands"
}

func (s DeleteMyCommands) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s DeleteMyCommands) ContentType() string {
	return "application/json"
}

// Use this method to get the current list of the bot's commands for the given scope and user language.
// Returns an Array of [objects.BotCommand] objects. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	// A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s GetMyCommands) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s GetMyCommands) Endpoint() string {
	return "getMyCommands"
}

func (s GetMyCommands) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetMyCommands) ContentType() string {
	return "application/json"
}

// Use this method to change the bot's name.
// Returns True on success.
type SetMyName struct {
	// New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	Name *string `json:"name,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s SetMyName) Validate() error {
	if s.Name != nil {
		if len(*s.Name) > 64 {
			return fmt.Errorf("name parameter must not be longer than 64 characters")
		}
	}
	return nil
}

func (s SetMyName) Endpoint() string {
	return "setMyName"
}

func (s SetMyName) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMyName) ContentType() string {
	return "application/json"
}

// Use this method to get the current bot name for the given user language.
// Returns BotName on success.
type GetMyName struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s GetMyName) Validate() error {
	return nil
}

func (s GetMyName) Endpoint() string {
	return "getMyName"
}

func (s GetMyName) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetMyName) ContentType() string {
	return "application/json"
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty.
// Returns True on success.
type SetMyDescription struct {
	// New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description *string `json:"description,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s SetMyDescription) Validate() error {
	if s.Description != nil {
		if len(*s.Description) > 64 {
			return fmt.Errorf("name parameter must not be longer than 64 characters")
		}
	}
	return nil
}

func (s SetMyDescription) Endpoint() string {
	return "setMyDescription"
}

func (s SetMyDescription) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMyDescription) ContentType() string {
	return "application/json"
}

// Use this method to get the current bot description for the given user language.
// Returns [objects.BotDescription] on success.
type GetMyDescription struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s GetMyDescription) Validate() error {
	return nil
}

func (s GetMyDescription) Endpoint() string {
	return "getMyDescription"
}

func (s GetMyDescription) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetMyDescription) ContentType() string {
	return "application/json"
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot.
// Returns True on success.
type SetMyShortDescription struct {
	// New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription *string `json:"short_description,omitempty"`
	// A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s SetMyShortDescription) Validate() error {
	if s.ShortDescription != nil {
		if len(*s.ShortDescription) > 64 {
			return fmt.Errorf("name parameter must not be longer than 64 characters")
		}
	}
	return nil
}

func (s SetMyShortDescription) Endpoint() string {
	return "setMyShortDescription"
}

func (s SetMyShortDescription) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMyShortDescription) ContentType() string {
	return "application/json"
}

// Use this method to get the current bot short description for the given user language.
// Returns [objects.BotShortDescription] on success.
type GetMyShortDescription struct {
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
}

func (s GetMyShortDescription) Validate() error {
	return nil
}

func (s GetMyShortDescription) Endpoint() string {
	return "getMyShortDescription"
}

func (s GetMyShortDescription) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetMyShortDescription) ContentType() string {
	return "application/json"
}

// Use this method to change the bot's menu button in a private chat, or the default menu button.
// Returns True on success.
type SetChatMenuButton struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	ChatId *string `json:"chat_id,omitempty"`
	// A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
	MenuButton objects.MenuButton `json:"menu_button,omitempty"`
}

func (s SetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if strings.TrimSpace(*s.ChatId) == "" {
			return fmt.Errorf("chat_id parameter can't be empty if specified")
		}
	}
	if s.MenuButton != nil {
		if err := s.MenuButton.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetChatMenuButton) Endpoint() string {
	return "setChatMenuButton"
}

func (s SetChatMenuButton) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetChatMenuButton) ContentType() string {
	return "application/json"
}

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button.
// Returns [objects.MenuButtonResponse] on success.
type GetChatMenuButton struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
	ChatId *int `json:"chat_id,omitempty"`
}

func (s GetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if *s.ChatId < 1 {
			return fmt.Errorf("chat_id parameter can't be empty if specified")
		}
	}
	return nil
}

func (s GetChatMenuButton) Endpoint() string {
	return "getChatMenuButton"
}

func (s GetChatMenuButton) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetChatMenuButton) ContentType() string {
	return "application/json"
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels.
// These rights will be suggested to users, but they are free to modify the list before adding the bot.
// Returns True on success.
type SetMyDefaultAdministratorRights struct {
	// A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	Rights *objects.ChatAdministratorRights `json:"rights,omitempty"`
	// Pass True to change the default administrator rights of the bot in channels.
	// Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels *bool `json:"for_channels,omitempty"`
}

func (s SetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s SetMyDefaultAdministratorRights) Endpoint() string {
	return "setMyDefaultAdministratorRights"
}

func (s SetMyDefaultAdministratorRights) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s SetMyDefaultAdministratorRights) ContentType() string {
	return "application/json"
}

// Use this method to get the current default administrator rights of the bot.
// Returns [objects.ChatAdministratorRights] on success.
type GetMyDefaultAdministratorRights struct {
	// Pass True to get default administrator rights of the bot in channels.
	// Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels *bool `json:"for_channels,omitempty"`
}

func (s GetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s GetMyDefaultAdministratorRights) Endpoint() string {
	return "getMyDefaultAdministratorRights"
}

func (s GetMyDefaultAdministratorRights) Reader() io.Reader {
	return gotely.EncodeJSON(s)
}

func (s GetMyDefaultAdministratorRights) ContentType() string {
	return "application/json"
}
