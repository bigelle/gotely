package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"slices"
	"strings"

	"github.com/bigelle/tele.go/objects"
	iso6391 "github.com/emvi/iso-639-1"
)

// Use this method to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	//Required
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Required
	//Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`
	//Optional.
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty,"`
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty,"`
	//Optional.
	//Mode for parsing entities in the message text.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty,"`
	//Optional.
	//A JSON-serialized list of special entities that appear in message text, which can be specified instead of parse_mode
	Entities *[]objects.MessageEntity `json:"entities,omitempty,"`
	//Optional.
	//Link preview generation options for the message
	LinkPreviewOptions *objects.LinkPreviewOptions `json:"link_preview_options,omitempty,"`
	//Optional.
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty,"`
	//Optional.
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty,"`
	//Optional.
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Optional.
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty,"`
	//Optional.
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty,"`
	//Optional.
	//Additional interface options. A JSON-serialized object for an inline keyboard,
	//custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty,"`
}

func (s SendMessage) Validate() error {
	if strings.TrimSpace(s.Text) == "" {
		return objects.ErrInvalidParam("text parameter can't be empty")
	}

	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (s SendMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendMessage) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendMessage", s)
}

// Use this method to forward messages of any kind.
// Service messages and messages with protected content can't be forwarded.
// On success, the sent Message is returned.
type ForwardMessage struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Required.
	//Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername
	FromChatId string `json:"from_chat_id"`
	//Required.
	//Message identifier in the chat specified in from_chat_id
	MessageId int `json:"message_id"`
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Optional.
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Optional.
	//Protects the contents of the forwarded message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
}

func (f ForwardMessage) Validate() error {
	if strings.TrimSpace(f.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(f.FromChatId) == "" {
		return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
	}
	if f.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (f ForwardMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(f)
}

func (f ForwardMessage) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("forwardMessage", f)
}

// Use this method to forward multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
// On success, an array of MessageId of the sent messages is returned.
type ForwardMessages struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Required.
	//Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`
	//Required.
	//A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward.
	//The identifiers must be specified in a strictly increasing order.
	MessageIds []int `json:"message_ids"`
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Optional.
	//Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Optional.
	//Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
}

func (f ForwardMessages) Validate() error {
	if strings.TrimSpace(f.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(f.FromChatId) == "" {
		return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
	}
	if len(f.MessageIds) < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (f ForwardMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(f)
}

func (f ForwardMessages) Execute() (*[]objects.MessageId, error) {
	return MakePostRequest[[]objects.MessageId]("forwardMessages", f)
}

// Use this method to copy messages of any kind.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessageut the copied message doesn't have a link to the original message.
// Returns the MessageId of the sent message on success.
type CopyMessage struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Required.
	//Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`
	//Required.
	//Message identifier in the chat specified in from_chat_id
	MessageId int `json:"message_id"`
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Optional.
	//New caption for media, 0-1024 characters after entities parsing. If not specified, the original caption is kept
	Caption *string `json:"caption,omitempty"`
	//Optional.
	//Mode for parsing entities in the new caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//Optional.
	//A JSON-serialized list of special entities that appear in the new caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Optional.
	//Pass True, if the caption must be shown above the message media. Ignored if a new caption isn't specified.
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Optional.
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Optional.
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Optional.
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Optional.
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Optional.
	//Additional interface options.
	//A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (c CopyMessage) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.FromChatId) == "" {
		return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
	}
	if c.MessageId < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CopyMessage) Execute() (*objects.MessageId, error) {
	return MakePostRequest[objects.MessageId]("copyMessage", c)
}

// Use this method to copy messages of any kind.
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessagesut the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages. On success, an array of MessageId of the sent messages is returned.
type CopyMessages struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatId string `json:"from_chat_id"`
	//A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to copy.
	//The identifiers must be specified in a strictly increasing order.
	MessageIds []int `json:"message_ids"`
	//Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to copy the messages without their captions
	RemoveCaption *bool `json:"remove_caption,omitempty"`
}

func (c CopyMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CopyMessages) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.FromChatId) == "" {
		return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
	}
	if len(c.MessageIds) < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessages) Execute() (*[]objects.MessageId, error) {
	return MakePostRequest[[]objects.MessageId]("copyMessages", c)
}

// Use this method to send photos. On success, the sent Message is returned.
type SendPhoto struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel
	//(in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get a photo from the Internet,
	//or upload a new photo using multipart/form-data.
	//The photo must be at most 10 MB in size. The photo's width and height must not exceed 10000 in total.
	//Width and height ratio must be at most 20.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Photo objects.InputFile `json:"photo"`
	//Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the photo caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Pass True if the photo needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard,
	//custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendPhoto) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendPhoto) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	//writing text fields
	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ShowCaptionAboveMedia != nil {
		if err := w.WriteField("show_caption_above_media", fmt.Sprint(*s.ShowCaptionAboveMedia)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	//writing file
	part, err := w.CreateFormFile("photo", s.Photo.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Photo.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	w.Close()
	return buf, w, nil
}

func (s SendPhoto) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPhoto) Execute() (*objects.Message, error) {
	if s.Photo.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendPhoto", s)
	}
	return MakePostRequest[objects.Message]("sendPhoto", s)
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned.
// Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
type SendAudio struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get an audio file from the Internet,
	//or upload a new one using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Audio objects.InputFile `json:"audio"`
	//Audio caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the audio caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Duration of the audio in seconds
	Duration *int `json:"duration,omitempty"`
	//Performer
	Performer *string `json:"performer,omitempty"`
	//Track name
	Title *string `json:"title,omitempty"`
	//Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	//Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file,
	//so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendAudio) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Audio.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendAudio) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.Duration != nil {
		if err := w.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.Title != nil {
		if err := w.WriteField("title", fmt.Sprint(*s.Title)); err != nil {
			return nil, nil, err
		}
	}
	if s.Performer != nil {
		if err := w.WriteField("performer", fmt.Sprint(*s.Performer)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("audio", s.Audio.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Audio.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	if s.Thumbnail != nil {
		switch (s.Thumbnail).(type) {
		case objects.InputFileFromRemote:
			if err := w.WriteField("thumbnail", fmt.Sprint(s.Thumbnail)); err != nil {
				return nil, nil, err
			}
		default:
			part, err := w.CreateFormFile("thumbnail", s.Thumbnail.Name())
			if err != nil {
				return nil, nil, err
			}
			reader, err := s.Thumbnail.Reader()
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(part, reader)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	w.Close()
	return buf, w, nil
}

func (s SendAudio) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAudio) Execute() (*objects.Message, error) {
	if s.Audio.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendAudio", s)
	}
	return MakePostRequest[objects.Message]("sendAudio", s)
}

// Use this method to send general files. On success, the sent Message is returned.
// Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Document objects.InputFile `json:"document"`
	//Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	//Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	//so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	//Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the document caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Disables automatic server-side content type detection for files uploaded using multipart/form-data
	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendDocument) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}

	if err := s.Document.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendDocument) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableContentTypeDetection != nil {
		if err := w.WriteField("disable_content_type_detection", fmt.Sprint(*s.DisableContentTypeDetection)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("document", s.Document.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Document.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	if s.Thumbnail != nil {
		switch (s.Thumbnail).(type) {
		case objects.InputFileFromRemote:
			if err := w.WriteField("thumbnail", fmt.Sprint(s.Thumbnail)); err != nil {
				return nil, nil, err
			}
		default:
			part, err := w.CreateFormFile("thumbnail", s.Thumbnail.Name())
			if err != nil {
				return nil, nil, err
			}
			reader, err := s.Thumbnail.Reader()
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(part, reader)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	w.Close()
	return buf, w, nil
}

func (s SendDocument) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDocument) Execute() (*objects.Message, error) {
	if s.Document.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendDocument", s)
	}
	return MakePostRequest[objects.Message]("sendDocument", s)
}

// Use this method to send video files, Telegram clients support MPEG4 videos (other formats may be sent as Document).
// On success, the sent Message is returned.
// Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Video objects.InputFile `json:"video"`
	//Duration of sent video in seconds
	Duration *int `json:"duration,omitempty"`
	//Video width
	Width *int `json:"width,omitempty"`
	//Video height
	Height *int `json:"height,omitempty"`
	//Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	//Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	//so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	//Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the video caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Pass True if the video needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	//Pass True if the uploaded video is suitable for streaming
	SupportsStreaming *bool `json:"supports_streaming,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendVideo) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Video.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendVideo) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Duration != nil {
		if err := w.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.Height != nil {
		if err := w.WriteField("height", fmt.Sprint(*s.Height)); err != nil {
			return nil, nil, err
		}
	}
	if s.Width != nil {
		if err := w.WriteField("width", fmt.Sprint(*s.Width)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.HasSpoiler != nil {
		if err := w.WriteField("has_spoiler", fmt.Sprint(*s.HasSpoiler)); err != nil {
			return nil, nil, err
		}
	}
	if s.SupportsStreaming != nil {
		if err := w.WriteField("supports_streaming", fmt.Sprint(*s.SupportsStreaming)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("video", s.Video.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Video.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	if s.Thumbnail != nil {
		switch (s.Thumbnail).(type) {
		case objects.InputFileFromRemote:
			if err := w.WriteField("thumbnail", fmt.Sprint(s.Thumbnail)); err != nil {
				return nil, nil, err
			}
		default:
			part, err := w.CreateFormFile("thumbnail", s.Thumbnail.Name())
			if err != nil {
				return nil, nil, err
			}
			reader, err := s.Thumbnail.Reader()
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(part, reader)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	w.Close()
	return buf, w, nil
}

func (s SendVideo) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideo) Execute() (*objects.Message, error) {
	if s.Video.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendVideo", s)
	}
	return MakePostRequest[objects.Message]("sendVideo", s)
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned.
// Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Animation objects.InputFile `json:"animation"`
	//Duration of sent animation in seconds
	Duration *int `json:"duration,omitempty"`
	//Animation width
	Width *int `json:"width,omitempty"`
	//Animation height
	Height *int `json:"height,omitempty"`
	//Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	//Ignored if the file is not uploaded using multipart/form-data. Thumbnails can't be reused and can be only uploaded as a new file,
	//so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	//Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the animation caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Pass True if the animation needs to be covered with a spoiler animation
	HasSpoiler *bool `json:"has_spoiler,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendAnimation) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Animation.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendAnimation) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Duration != nil {
		if err := w.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.Height != nil {
		if err := w.WriteField("height", fmt.Sprint(*s.Height)); err != nil {
			return nil, nil, err
		}
	}
	if s.Width != nil {
		if err := w.WriteField("width", fmt.Sprint(*s.Width)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.HasSpoiler != nil {
		if err := w.WriteField("has_spoiler", fmt.Sprint(*s.HasSpoiler)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("animation", s.Animation.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Animation.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	if s.Thumbnail != nil {
		switch (s.Thumbnail).(type) {
		case objects.InputFileFromRemote:
			if err := w.WriteField("thumbnail", fmt.Sprint(s.Thumbnail)); err != nil {
				return nil, nil, err
			}
		default:
			part, err := w.CreateFormFile("thumbnail", s.Thumbnail.Name())
			if err != nil {
				return nil, nil, err
			}
			reader, err := s.Thumbnail.Reader()
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(part, reader)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	w.Close()
	return buf, w, nil
}

func (s SendAnimation) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAnimation) Execute() (*objects.Message, error) {
	if s.Animation.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendAnimation", s)
	}
	return MakePostRequest[objects.Message]("sendAnimation", s)
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS, or in .MP3 format, or in .M4A format
// (other formats may be sent as Audio or Document).
// On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size,
// this limit may be changed in the future.
type SendVoice struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	//pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Voice objects.InputFile `json:"voice"`
	//Voice message caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the voice message caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities"`
	//Duration of the voice message in seconds
	Duration *int `json:"duration,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendVoice) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Voice.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendVoice) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVoice) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Duration != nil {
		if err := w.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.Caption != nil {
		if err := w.WriteField("caption", fmt.Sprint(*s.Caption)); err != nil {
			return nil, nil, err
		}
	}
	if s.ParseMode != nil {
		if err := w.WriteField("parse_mode", fmt.Sprint(*s.ParseMode)); err != nil {
			return nil, nil, err
		}
	}
	if s.CaptionEntities != nil {
		b, err := json.Marshal(s.CaptionEntities)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("caption_entities", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("voice", s.Voice.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Voice.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	w.Close()
	return buf, w, nil
}

func (s SendVoice) Execute() (*objects.Message, error) {
	if s.Voice.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendVoice", s)
	}
	return MakePostRequest[objects.Message]("sendVoice", s)
}

// As of v.4.0, Telegram clients support rounded square MPEG4 videos of up to 1 minute long.
// Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int `json:"message_thread_id,omitempty"`
	//Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or
	//upload a new video using multipart/form-data.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files.
	//Sending video notes by a URL is currently unsupported
	VideoNote objects.InputFile `json:"video_note"`
	//Duration of sent video in seconds
	Duration *int `json:"duration,omitempty"`
	//Video width and height, i.e. diameter of the video message
	Length *int `json:"length,omitempty"`
	//Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	//The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail's width and height should not exceed 320.
	//Ignored if the file is not uploaded using multipart/form-data.
	//Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if
	//the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	//More information on Sending Files: https://core.telegram.org/bots/api#sending-files
	Thumbnail objects.InputFile `json:"thumbnail,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendVideoNote) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.VideoNote.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SendVideoNote) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", fmt.Sprintf("%v", s.ChatId)); err != nil {
		return nil, nil, err
	}
	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", fmt.Sprint(*s.MessageThreadId)); err != nil {
			return nil, nil, err
		}
	}
	if s.Duration != nil {
		if err := w.WriteField("duration", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.Length != nil {
		if err := w.WriteField("length", fmt.Sprint(*s.Duration)); err != nil {
			return nil, nil, err
		}
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(*s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(*s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(*s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", fmt.Sprint(*s.MessageEffectId)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		b, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(b)); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyMarkup != nil {
		b, err := json.Marshal(s.ReplyMarkup)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_markup", string(b)); err != nil {
			return nil, nil, err
		}
	}

	part, err := w.CreateFormFile("video_note", s.VideoNote.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.VideoNote.Reader()
	if err != nil {
		return nil, nil, err
	}
	_, err = io.Copy(part, reader)
	if err != nil {
		return nil, nil, err
	}

	if s.Thumbnail != nil {
		switch (s.Thumbnail).(type) {
		case objects.InputFileFromRemote:
			if err := w.WriteField("thumbnail", fmt.Sprint(s.Thumbnail)); err != nil {
				return nil, nil, err
			}
		default:
			part, err := w.CreateFormFile("thumbnail", s.Thumbnail.Name())
			if err != nil {
				return nil, nil, err
			}
			reader, err := s.Thumbnail.Reader()
			if err != nil {
				return nil, nil, err
			}
			_, err = io.Copy(part, reader)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	w.Close()
	return buf, w, nil
}

func (s SendVideoNote) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideoNote) Execute() (*objects.Message, error) {
	if s.VideoNote.IsLocal() {
		return MakeMultipartRequest[objects.Message]("sendVideoNote", s)
	}
	return MakePostRequest[objects.Message]("sendVideoNote", s)
}

// Use this method to send paid media. On success, the sent Message is returned.
type SendPaidMedia struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	//If the chat is a channel, all Telegram Star proceeds from this media will be credited to the chat's balance.
	//Otherwise, they will be credited to the bot's balance.
	ChatId string `json:"chat_id"`
	//The number of Telegram Stars that must be paid to buy access to the media; 1-2500
	StarCount int `json:"star_count"`
	//An array describing the media to be sent; up to 10 items
	Media []objects.InputPaidMedia `json:"media"`
	//Bot-defined paid media payload, 0-128 bytes. This will not be displayed to the user, use it for your internal processes.
	Payload *string `json:"payload,omitempty"`
	//Media caption, 0-1024 characters after entities parsing
	Caption *string `json:"caption,omitempty"`
	//Mode for parsing entities in the media caption.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ParseMode *string `json:"parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the caption, which can be specified instead of parse_mode
	CaptionEntities *[]objects.MessageEntity `json:"caption_entities,omitempty"`
	//Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia *bool `json:"show_caption_above_media,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendPaidMedia) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.StarCount < 1 || s.StarCount > 2500 {
		return objects.ErrInvalidParam("star_count parameter must be between 1 and 2500")
	}
	if len(s.Media) < 1 {
		return objects.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return objects.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendPaidMedia) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPaidMedia) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPaidMedia", s)
}

// Use this method to send a group of photos, videos, documents or audios as an album.
// Documents and audio files can be only grouped in an album with messages of the same type.
// On success, an array of Messages that were sent is returned.
type SendMediaGroup struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//An array describing messages to be sent, must include 2-10 items
	Media []objects.InputMedia `json:"media"`
	//Sends messages silently. Users will receive a notif,omitemptyication with no sound.
	DisableNotification *bool `json:"disable_notification"`
	//Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// /Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
}

func (s SendMediaGroup) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(s.Media) < 1 {
		return objects.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return objects.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMediaGroup) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendMediaGroup) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendMediaGroup", s)
}

// Use this method to send point on the map. On success, the sent Message is returned.
type SendLocation struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//Latitude of the location
	Latitude *float64 `json:"latitude"`
	//Longitude of the location
	Longtitude *float64 `json:"longtitude"`
	//The radius of uncertainty for the location, measured in meters; 0-1500
	HorizontalAccuracy *float64 `json:"horizontal_accuracy,omitempty"`
	//Period in seconds during which the location will be updated
	//(see https://telegram.org/blog/live-locations), should be between 60 and 86400,
	//or 0x7FFFFFFF for live locations that can be edited indefinitely.
	LivePeriod *int `json:"live_period,omitempty"`
	//For live locations, a direction in which the user is moving, in degrees. Must be between 1 and 360 if specified.
	Heading *int `json:"heading,omitempty"`
	//For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	//Must be between 1 and 100000 if specified.
	ProximityAlertRadius *int `json:"proximity_alert_radius,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendLocation) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return objects.ErrInvalidParam("longtitude parameter can't be empty")
	}
	// TODO: validate replyparameters everywhere
	return nil
}

func (s SendLocation) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendLocation) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendLocation", s)
}

// Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//Latitude of the venue
	Latitude *float64 `json:"latitude"`
	//Longitude of the venue
	Longtitude *float64 `json:"longtitude"`
	//Name of the venue
	Title string `json:"title"`
	//Address of the venue
	Address string `json:"address"`
	//Foursquare identifier of the venue
	FoursquareId *string `json:"foursquare_id,omitempty"`
	//Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType *string `json:"foursquare_type,omitempty"`
	//Google Places identifier of the venue
	GooglePlaceId *string `json:"google_place_id,omitempty"`
	//Google Places type of the venue. (See https://developers.google.com/places/web-service/supported_types)
	GooglePlaceType *string `json:"google_place_type,omitempty"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	//Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	//Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	//instructions to remove a reply keyboard or to force a reply from the user
	ReplyMarkup *objects.ReplyMarkup `json:"reply_markup,omitempty"`
}

func (s SendVenue) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return objects.ErrInvalidParam("longtitude parameter can't be empty")
	}
	return nil
}

func (s SendVenue) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVenue) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendVenue", s)
}

// Use this method to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId              string                   `json:"chat_id"`
	PhoneNumber         string                   `json:"phone_number"`
	FirstName           string                   `json:"first_name"`
	LastName            *string                  `json:"last_name,omitempty"`
	Vcard               *string                  `json:"vcard,omitempty"`
	MessageThreadId     *string                  `json:"message_thread_id,omitempty"`
	DisableNotification *bool                    `json:"disable_notification,omitempty"`
	ProtectContent      *bool                    `json:"protect_content,omitempty"`
	AllowPaidBroadcast  *bool                    `json:"allow_paid_broadcast,omitempty"`
	MessageEffectId     *string                  `json:"message_effect_id,omitempty"`
	ReplyParameters     *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         *objects.ReplyMarkup     `json:"reply_markup,omitempty"`
}

func (s SendContact) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.PhoneNumber) == "" {
		return objects.ErrInvalidParam("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return objects.ErrInvalidParam("first_name parameter can't be empty")
	}
	return nil
}

func (s SendContact) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendContact) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendContact", s)
}

type SendPoll struct {
	ChatId               string
	Question             string
	Options              []objects.InputPollOption
	QuestionParseMode    *string
	QuestionEntities     *[]objects.MessageEntity
	IsAnonymous          *bool
	Type                 *string
	AllowMultipleAnswers *bool
	CorrectOptionId      *int
	Explanation          *string
	ExplanationParseMode *string
	ExplanationEntities  *[]objects.MessageEntity
	OpenPeriod           *int
	CloseDate            *int
	IsClosed             *bool
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendPoll) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Question) == "" {
		return objects.ErrInvalidParam("question parameter can't be empty")
	}
	if len(s.Options) < 2 || len(s.Options) > 10 {
		return objects.ErrInvalidParam("options parameter must be between 2 and 10")
	}
	return nil
}

func (s SendPoll) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPoll) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPoll", s)
}

type SendDice struct {
	ChatId               string
	Emoji                string
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendDice) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Emoji) == "" {
		return objects.ErrInvalidParam("emoji parameter can't be empty")
	}
	return nil
}

func (s SendDice) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDice) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPoll", s)
}

type SendChatAction struct {
	ChatId               string
	Action               string
	BusinessConnectionId *string
	MessageThreadId      *string
}

func (s SendChatAction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Action) == "" {
		return objects.ErrInvalidParam("action parameter can't be empty")
	}
	allowed := []string{
		"typing",
		"upload_photo",
		"record_video",
		"record_voice",
		"upload_voice",
		"upload_document",
		"choose_sticker",
		"find_location",
		"record_video_note",
		"upload_video_note",
	}
	// NOTE: maybe there's a better way
	if !slices.Contains(allowed, s.Action) {
		return objects.ErrInvalidParam(fmt.Sprintf("action must be %s or upload_video_note", strings.Join(allowed[:len(allowed)-1], ", ")))
	}
	return nil
}

type SetMessageReaction struct {
	ChatId    string
	MessageId int
	Reaction  *[]objects.ReactionType
	IsBig     *bool
}

func (s SetMessageReaction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (s SetMessageReaction) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMessageReaction) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMessageReaction", s)
}

type GetUserProfilePhotos struct {
	UserId int
	Offset *int
	Limit  *int
}

func (g GetUserProfilePhotos) Validate() error {
	if g.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if g.Limit != nil {
		if *g.Limit < 1 || *g.Limit > 100 {
			return objects.ErrInvalidParam("limit parameter must be between 1 and 100")
		}
	}
	return nil
}

func (g GetUserProfilePhotos) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetUserProfilePhotos) Execute() (*objects.UserProfilePhotos, error) {
	return MakeGetRequest[objects.UserProfilePhotos]("getUserProfilePhotos", g)
}

type SetUserEmojiStatus struct {
	UserId                    int
	EmojiStatusCustomEmojiId  *string
	EmojiStatusExpirationDate *int
}

func (s SetUserEmojiStatus) Validate() error {
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s SetUserEmojiStatus) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetUserEmojiStatus) Execute() (*bool, error) {
	return MakePostRequest[bool]("setUserEmojiStatus", s)
}

type GetFile struct {
	FileId string
}

func (g GetFile) Validate() error {
	if strings.TrimSpace(g.FileId) == "" {
		return objects.ErrInvalidParam("file_id parameter can't be empty")
	}
	return nil
}

func (g GetFile) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetFile) Execute() (*objects.File, error) {
	return MakeGetRequest[objects.File]("getFile", g)
}

type BanChatMember struct {
	ChatId         string
	UserId         int
	UntilDate      *int
	RevokeMessages *bool
}

func (b BanChatMember) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if b.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b BanChatMember) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatMember) Execute() (*bool, error) {
	return MakeGetRequest[bool]("banChatMember", b)
}

type UnbanChatMember struct {
	ChatId       string
	UserId       int
	OnlyIfBanned *bool
}

func (b UnbanChatMember) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if b.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatMember) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatMember) Execute() (*bool, error) {
	return MakeGetRequest[bool]("unbanChatMember", b)
}

type RestrictChatMember struct {
	ChatId                         string
	UserId                         int
	Permissions                    objects.ChatPermissions
	UserIndependentChatPermissions *bool
	UntilDate                      *int
}

func (r RestrictChatMember) Validate() error {
	if strings.TrimSpace(r.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if r.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (r RestrictChatMember) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r RestrictChatMember) Execute() (*bool, error) {
	return MakePostRequest[bool]("restrictChatMember", r)
}

type PromoteChatMember struct {
	ChatId              string
	UserId              int
	IsAnonymous         *bool
	CanManageChat       *bool
	CanDeleteMessages   *bool
	CanManageVideoChats *bool
	CanRestrictMembers  *bool
	CanPromoteMembers   *bool
	CanChangeInfo       *bool
	CanInviteUsers      *bool
	CanPostStories      *bool
	CanEditStories      *bool
	CanDeleteStories    *bool
	CanPostMessages     *bool
	CanEditMessages     *bool
	CanPinMessages      *bool
	CanManageTopics     *bool
}

func (p PromoteChatMember) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p PromoteChatMember) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PromoteChatMember) Execute() (*bool, error) {
	return MakePostRequest[bool]("promoteChatMember", p)
}

type SetChatAdministratorCustomTitle struct {
	ChatId      string
	UserId      int
	CustomTitle string
}

func (s SetChatAdministratorCustomTitle) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.CustomTitle) > 16 {
		return objects.ErrInvalidParam("custom_title parameter must be not longer than 16 characters")
	}
	for _, r := range s.CustomTitle {
		if (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
			(r >= 0x1F300 && r <= 0x1F5FF) || // Miscellaneous Symbols and Pictographs
			(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
			(r >= 0x1F700 && r <= 0x1F77F) { // Alchemical Symbols
			return objects.ErrInvalidParam("invalid custom_title parameter: emojis are not allowed")
		}
	}
	return nil
}

func (s SetChatAdministratorCustomTitle) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatAdministratorCustomTitle) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatAdministratorCustomTitle", s)
}

type BanChatSenderChat struct {
	ChatId       string
	SenderChatId int
}

func (b BanChatSenderChat) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if b.SenderChatId < 1 {
		return objects.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b BanChatSenderChat) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatSenderChat) Execute() (*bool, error) {
	return MakePostRequest[bool]("banChatSenderChat", b)
}

type UnbanChatSenderChat struct {
	ChatId       string
	SenderChatId int
}

func (b UnbanChatSenderChat) Validate() error {
	if strings.TrimSpace(b.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if b.SenderChatId < 1 {
		return objects.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatSenderChat) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatSenderChat) Execute() (*bool, error) {
	return MakePostRequest[bool]("unbanChatSenderChat", b)
}

type SetChatPermissions struct {
	ChatId                         string
	Permissions                    objects.ChatPermissions
	UserIndependentChatPermissions *bool
}

func (s SetChatPermissions) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (s SetChatPermissions) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPermissions) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatPermissions", s)
}

type ExportChatInviteLink struct {
	ChatId string
}

func (e ExportChatInviteLink) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e ExportChatInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ExportChatInviteLink) Execute() (*string, error) {
	return MakePostRequest[string]("exportChatInviteLink", e)
}

type CreateInviteLink struct {
	ChatId             string
	Name               *string
	ExpireDate         *int
	MemberLimit        *int
	CreatesJoinRequest *bool
}

func (c CreateInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return objects.ErrInvalidParam("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (c CreateInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateInviteLink) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("createInviteLink", c)
}

type EditInviteLink struct {
	ChatId             string
	InviteLink         string
	Name               *string
	ExpireDate         *int
	MemberLimit        *int
	CreatesJoinRequest *bool
}

func (c EditInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return objects.ErrInvalidParam("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (c EditInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditInviteLink) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("editInviteLink", c)
}

type CreateChatSubscriptionInviteLink struct {
	ChatId             string
	SubscriptionPeriod int
	SubscriptionPrice  int
	Name               *string
}

func (c CreateChatSubscriptionInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if c.SubscriptionPeriod != 2592000 {
		return objects.ErrInvalidParam("subscription_period currently must always be 2592000 seconds (30 days)")
	}
	if c.SubscriptionPrice < 1 || c.SubscriptionPrice > 2500 {
		return objects.ErrInvalidParam("subscription_price must be between 1 and 2500")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c CreateChatSubscriptionInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateChatSubscriptionInviteLink) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("createChatSubscriptionInviteLink", c)
}

type EditChatSubscriptionInviteLink struct {
	ChatId     string
	InviteLink string
	Name       *string
}

func (c EditChatSubscriptionInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.InviteLink) == "" {
		return objects.ErrInvalidParam("invite_link parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c EditChatSubscriptionInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditChatSubscriptionInviteLink) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("editChatSubscriptionInviteLink", c)
}

type RevokeInviteLink struct {
	ChatId     string
	InviteLink string
	Name       *string
}

func (c RevokeInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c RevokeInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c RevokeInviteLink) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("revokeInviteLink", c)
}

type ApproveChatJoinRequest struct {
	ChatId string
	UserId int
}

func (s ApproveChatJoinRequest) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s ApproveChatJoinRequest) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s ApproveChatJoinRequest) Execute() (*bool, error) {
	return MakePostRequest[bool]("approveChatJoinRequest", s)
}

type DeclineChatJoinRequest struct {
	ChatId string
	UserId int
}

func (s DeclineChatJoinRequest) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s DeclineChatJoinRequest) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeclineChatJoinRequest) Execute() (*bool, error) {
	return MakePostRequest[bool]("declineChatJoinRequest", s)
}

type SetChatPhoto struct {
	ChatId string
	Photo  objects.InputFile
}

func (s SetChatPhoto) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SetChatPhoto) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPhoto) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatPhoto", s)
}

type DeleteChatPhoto struct {
	ChatId string
}

func (d DeleteChatPhoto) Validate() error {
	if strings.TrimSpace(d.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (d DeleteChatPhoto) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteChatPhoto) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteChatPhoto", d)
}

type SetChatTitle struct {
	ChatId string
	Title  string
}

func (s SetChatTitle) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(s.Title) < 1 || len(s.Title) > 128 {
		return objects.ErrInvalidParam("title parameter must be between 1 and 128 characters long")
	}
	return nil
}

func (s SetChatTitle) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatTitle) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatTitle", s)
}

type SetChatDescription struct {
	ChatId      string
	Description string
}

func (s SetChatDescription) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(s.Description) > 255 {
		return objects.ErrInvalidParam("description parameter must not be longer than 255 characters")
	}
	return nil
}

func (s SetChatDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatDescription) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatTitle", s)
}

type PinChatMessage struct {
	ChatId               string
	MessageId            int
	BusinessConnectionId *string
	DisableNotification  *bool
}

func (p PinChatMessage) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p PinChatMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PinChatMessage) Execute() (*bool, error) {
	return MakePostRequest[bool]("pinChatMessage", p)
}

type UnpinChatMessage struct {
	ChatId               string
	MessageId            int
	BusinessConnectionId *string
}

func (p UnpinChatMessage) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p UnpinChatMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinChatMessage) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinChatMessage", p)
}

type UnpinAllChatMessages struct {
	ChatId string
}

func (p UnpinAllChatMessages) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p UnpinAllChatMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinAllChatMessages) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinAllChatMessages", p)
}

type LeaveChat struct {
	ChatId string
}

func (p LeaveChat) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p LeaveChat) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p LeaveChat) Execute() (*bool, error) {
	return MakePostRequest[bool]("leaveChat", p)
}

type GetChat struct {
	ChatId string
}

func (p GetChat) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p GetChat) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChat) Execute() (*objects.ChatFullInfo, error) {
	return MakeGetRequest[objects.ChatFullInfo]("getChat", p)
}

type GetChatAdministrators struct {
	ChatId string
}

func (p GetChatAdministrators) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p GetChatAdministrators) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatAdministrators) Execute() (*[]objects.ChatMember, error) {
	return MakeGetRequest[[]objects.ChatMember]("getChatAdministrators", p)
}

type GetChatMemberCount struct {
	ChatId string
}

func (p GetChatMemberCount) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p GetChatMemberCount) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatMemberCount) Execute() (*int, error) {
	return MakeGetRequest[int]("getChatMemberCount", p)
}

type GetChatMember struct {
	ChatId string
	UserId int
}

func (p GetChatMember) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p GetChatMember) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatMember) Execute() (*objects.ChatMember, error) {
	return MakeGetRequest[objects.ChatMember]("getChatMember", p)
}

type SetChatStickerSet struct {
	ChatId         string
	StickerSetName string
}

func (p SetChatStickerSet) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(p.StickerSetName) == "" {
		return objects.ErrInvalidParam("sticker_set_name parameter can't be empty")
	}
	return nil
}

func (p SetChatStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p SetChatStickerSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatStickerSet", p)
}

type DeleteChatStickerSet struct {
	ChatId string
}

func (p DeleteChatStickerSet) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (p DeleteChatStickerSet) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p DeleteChatStickerSet) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteChatStickerSet", p)
}

type GetForumTopicIconStickers struct {
}

// always nil
func (g GetForumTopicIconStickers) Validate() error {
	return nil
}

// alwways empty json
func (g GetForumTopicIconStickers) ToRequestBody() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetForumTopicIconStickers) Execute() (*[]objects.Sticker, error) {
	return MakeGetRequest[[]objects.Sticker]("getForumTopicStickers", g)
}

var validIconColors = map[int]struct{}{
	7322096:  {},
	16766590: {},
	13338331: {},
	9367192:  {},
	16749490: {},
	16478047: {},
}

type CreateForumTopic struct {
	ChatId            string
	Name              string
	IconColor         *int
	IconCustomEmojiId *string
}

func (c CreateForumTopic) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 128 {
		return objects.ErrInvalidParam("name parameter must be between 1 and 128 characters long")
	}
	if c.IconColor != nil {
		if _, ok := validIconColors[*c.IconColor]; !ok {
			return objects.ErrInvalidParam("icon_color must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)")
		}
	}
	return nil
}

func (c CreateForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateForumTopic) Execute() (*objects.ForumTopic, error) {
	return MakePostRequest[objects.ForumTopic]("createForumTopic", c)
}

type EditForumTopic struct {
	ChatId            string
	MessageThreadId   string
	Name              *string
	IconCustomEmojiId *string
}

func (e EditForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	if e.Name != nil {
		if len(*e.Name) > 128 {
			return objects.ErrInvalidParam("name parameter must not be longer than 128 characters")
		}
	}
	return nil
}

func (e EditForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("editForumTopic", e)
}

type CloseForumTopic struct {
	ChatId          string
	MessageThreadId string
}

func (e CloseForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e CloseForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("closeForumTopic", e)
}

type ReopenForumTopic struct {
	ChatId          string
	MessageThreadId string
}

func (e ReopenForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e ReopenForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("reopenForumTopic", e)
}

type DeleteForumTopic struct {
	ChatId          string
	MessageThreadId string
}

func (e DeleteForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e DeleteForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e DeleteForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteForumTopic", e)
}

type UnpinAllForumTopicMessages struct {
	ChatId          string
	MessageThreadId string
}

func (e UnpinAllForumTopicMessages) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e UnpinAllForumTopicMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllForumTopicMessages) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinAllForumTopicMessages", e)
}

type EditGeneralForumTopic struct {
	ChatId string
	Name   string
}

func (e EditGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(e.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (e EditGeneralForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditGeneralForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("editGeneralForumTopic", e)
}

type CloseGeneralForumTopic struct {
	ChatId string
}

func (e CloseGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e CloseGeneralForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseGeneralForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("closeGeneralForumTopic", e)
}

type ReopenGeneralForumTopic struct {
	ChatId string
}

func (e ReopenGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e ReopenGeneralForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenGeneralForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("reopenGeneralForumTopic", e)
}

type HideGeneralForumTopic struct {
	ChatId string
}

func (e HideGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e HideGeneralForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e HideGeneralForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("hideGeneralForumTopic", e)
}

type UnhideGeneralForumTopic struct {
	ChatId string
}

func (e UnhideGeneralForumTopic) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e UnhideGeneralForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnhideGeneralForumTopic) Execute() (*bool, error) {
	return MakePostRequest[bool]("unhideGeneralForumTopic", e)
}

type UnpinAllGeneralForumTopicMessages struct {
	ChatId string
}

func (e UnpinAllGeneralForumTopicMessages) Validate() error {
	if strings.TrimSpace(e.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	return nil
}

func (e UnpinAllGeneralForumTopicMessages) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllGeneralForumTopicMessages) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinAllGeneralForumTopicMessages", e)
}

type AnswerCallbackQuery struct {
	CallbackQueryId string
	Text            *string
	ShowAlert       *bool
	Url             *string
	CacheTime       *int
}

func (a AnswerCallbackQuery) Validate() error {
	if strings.TrimSpace(a.CallbackQueryId) == "" {
		return objects.ErrInvalidParam("callback_query_id parameter can't be empty")
	}
	if a.Text != nil {
		if len(*a.Text) > 200 {
			return objects.ErrInvalidParam("text parameter must not be longer than 200 characters ")
		}
	}
	return nil
}

func (a AnswerCallbackQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerCallbackQuery) Execute() (*bool, error) {
	return MakePostRequest[bool]("answerCallbackQuery", a)
}

type GetUserChatBoosts struct {
	ChatId string
	UserId int
}

func (g GetUserChatBoosts) Validate() error {
	if strings.TrimSpace(g.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if g.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (g GetUserChatBoosts) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetUserChatBoosts) Execute() (*objects.UserChatBoosts, error) {
	return MakeGetRequest[objects.UserChatBoosts]("getUserChatBoosts", g)
}

type GetBusinessConnection struct {
	BusinessConnectionId string
}

func (g GetBusinessConnection) Validate() error {
	if strings.TrimSpace(g.BusinessConnectionId) == "" {
		return objects.ErrInvalidParam("business_connection_id parameter can't be empty")
	}
	return nil
}

func (g GetBusinessConnection) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetBusinessConnection) Execute() (*objects.BusinessConnection, error) {
	return MakeGetRequest[objects.BusinessConnection]("getBusinessConnection", g)
}

type SetMyCommands struct {
	Commands     []objects.BotCommand
	Scope        *objects.BotCommandScope
	LanguageCode *string
}

func (s SetMyCommands) Validate() error {
	for _, command := range s.Commands {
		if err := command.Validate(); err != nil {
			return err
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		// FIXME: should validate it using no dependencies
		//https://ru.wikipedia.org/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%BA%D0%BE%D0%B4%D0%BE%D0%B2_ISO_639-1
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyCommands) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMyCommands", s)
}

type DeleteMyCommands struct {
	Scope        *objects.BotCommandScope
	LanguageCode *string
}

func (s DeleteMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s DeleteMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeleteMyCommands) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteMyCommands", s)
}

type GetMyCommands struct {
	Scope        *objects.BotCommandScope
	LanguageCode *string
}

func (s GetMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyCommands) Execute() (*[]objects.BotCommand, error) {
	return MakeGetRequest[[]objects.BotCommand]("getMyCommands", s)
}

type SetMyName struct {
	Name         *string
	LanguageCode *string
}

func (s SetMyName) Validate() error {
	if s.Name != nil {
		if len(*s.Name) > 64 {
			return objects.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyName) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyName) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMyName", s)
}

type GetMyName struct {
	LanguageCode *string
}

func (s GetMyName) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyName) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyName) Execute() (*objects.BotName, error) {
	return MakeGetRequest[objects.BotName]("getMyName", s)
}

type SetMyDescription struct {
	Description  *string
	LanguageCode *string
}

func (s SetMyDescription) Validate() error {
	if s.Description != nil {
		if len(*s.Description) > 64 {
			return objects.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyDescription) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMyDescription", s)
}

type GetMyDescription struct {
	LanguageCode *string
}

func (s GetMyDescription) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyDescription) Execute() (*objects.BotDescription, error) {
	return MakeGetRequest[objects.BotDescription]("getMyDescription", s)
}

type SetMyShortDescription struct {
	ShortDescription *string
	LanguageCode     *string
}

func (s SetMyShortDescription) Validate() error {
	if s.ShortDescription != nil {
		if len(*s.ShortDescription) > 64 {
			return objects.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyShortDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyShortDescription) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMyShortDescription", s)
}

type GetMyShortDescription struct {
	LanguageCode *string
}

func (s GetMyShortDescription) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s GetMyShortDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyShortDescription) Execute() (*bool, error) {
	return MakeGetRequest[bool]("getMyShortDescription", s)
}

type SetChatMenuButton struct {
	ChatId     *string
	MenuButton objects.MenuButton
}

func (s SetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if strings.TrimSpace(*s.ChatId) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if specified")
		}
	}
	if s.MenuButton != nil {
		if err := s.MenuButton.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetChatMenuButton) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatMenuButton) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatMenuButton", s)
}

type GetChatMenuButton struct {
	ChatId *int
}

func (s GetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if *s.ChatId < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if specified")
		}
	}
	return nil
}

func (s GetChatMenuButton) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetChatMenuButton) Execute() (*objects.MenuButtonResponse, error) {
	return MakeGetRequest[objects.MenuButtonResponse]("setChatMenuButton", s)
}

type SetMyDefaultAdministratorRights struct {
	Rights      *objects.ChatAdministratorRights
	ForChannels *bool
}

// always nil
func (s SetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s SetMyDefaultAdministratorRights) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyDefaultAdministratorRights) Execute() (*bool, error) {
	return MakePostRequest[bool]("setMyDefaultAdministratorRights", s)
}

type GetMyDefaultAdministratorRights struct {
	ForChannels *bool
}

// always nil
func (s GetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s GetMyDefaultAdministratorRights) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyDefaultAdministratorRights) Execute() (*objects.ChatAdministratorRights, error) {
	return MakePostRequest[objects.ChatAdministratorRights]("getMyDefaultAdministratorRights", s)
}
