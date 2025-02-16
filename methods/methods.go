// TODO: make optional and required fields more obvious
// TODO: add documentation for struct methods in Go style
// TODO: replace iso6391 dependency with self-made function OR package (do i need it?)
package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/bigelle/gotely/objects"
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
	client      *http.Client
	baseUrl     string
}

func (s *SendMessage) WithClient(c *http.Client) *SendMessage {
	s.client = c
	return s
}

func (s SendMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendMessage) WithApiBaseUrl(u string) *SendMessage {
	s.baseUrl = u
	return s
}

func (s SendMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendMessage) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendMessage", s)
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
	client         *http.Client
	baseUrl        string
}

func (s *ForwardMessage) WithClient(c *http.Client) *ForwardMessage {
	s.client = c
	return s
}

func (s ForwardMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ForwardMessage) WithApiBaseUrl(u string) *ForwardMessage {
	s.baseUrl = u
	return s
}

func (s ForwardMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (f ForwardMessage) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "forwardMessage", f)
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
	client         *http.Client
	baseUrl        string
}

func (s *ForwardMessages) WithClient(c *http.Client) *ForwardMessages {
	s.client = c
	return s
}

func (s ForwardMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ForwardMessages) WithApiBaseUrl(u string) *ForwardMessages {
	s.baseUrl = u
	return s
}

func (s ForwardMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (f ForwardMessages) Execute(token string) (*[]objects.MessageId, error) {
	return SendTelegramPostRequest[[]objects.MessageId](token, "forwardMessages", f)
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
	client      *http.Client
	baseUrl     string
}

func (s *CopyMessage) WithClient(c *http.Client) *CopyMessage {
	s.client = c
	return s
}

func (s CopyMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CopyMessage) WithApiBaseUrl(u string) *CopyMessage {
	s.baseUrl = u
	return s
}

func (s CopyMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CopyMessage) Execute(token string) (*objects.MessageId, error) {
	return SendTelegramPostRequest[objects.MessageId](token, "copyMessage", c)
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
	client        *http.Client
	baseUrl       string
}

func (s *CopyMessages) WithClient(c *http.Client) *CopyMessages {
	s.client = c
	return s
}

func (s CopyMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CopyMessages) WithApiBaseUrl(u string) *CopyMessages {
	s.baseUrl = u
	return s
}

func (s CopyMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CopyMessages) Execute(token string) (*[]objects.MessageId, error) {
	return SendTelegramPostRequest[[]objects.MessageId](token, "copyMessages", c)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendPhoto) WithClient(c *http.Client) *SendPhoto {
	s.client = c
	return s
}

func (s SendPhoto) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendPhoto) WithApiBaseUrl(u string) *SendPhoto {
	s.baseUrl = u
	return s
}

func (s SendPhoto) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendPhoto) Execute(token string) (*objects.Message, error) {
	if s.Photo.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendPhoto", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendPhoto", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendAudio) WithClient(c *http.Client) *SendAudio {
	s.client = c
	return s
}

func (s SendAudio) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendAudio) WithApiBaseUrl(u string) *SendAudio {
	s.baseUrl = u
	return s
}

func (s SendAudio) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendAudio) Execute(token string) (*objects.Message, error) {
	if s.Audio.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendAudio", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendAudio", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendDocument) WithClient(c *http.Client) *SendDocument {
	s.client = c
	return s
}

func (s SendDocument) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendDocument) WithApiBaseUrl(u string) *SendDocument {
	s.baseUrl = u
	return s
}

func (s SendDocument) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendDocument) Execute(token string) (*objects.Message, error) {
	if s.Document.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendDocument", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendDocument", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendVideo) WithClient(c *http.Client) *SendVideo {
	s.client = c
	return s
}

func (s SendVideo) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendVideo) WithApiBaseUrl(u string) *SendVideo {
	s.baseUrl = u
	return s
}

func (s SendVideo) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendVideo) Execute(token string) (*objects.Message, error) {
	if s.Video.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendVideo", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendVideo", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendAnimation) WithClient(c *http.Client) *SendAnimation {
	s.client = c
	return s
}

func (s SendAnimation) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendAnimation) WithApiBaseUrl(u string) *SendAnimation {
	s.baseUrl = u
	return s
}

func (s SendAnimation) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendAnimation) Execute(token string) (*objects.Message, error) {
	if s.Animation.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendAnimation", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendAnimation", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendVoice) WithClient(c *http.Client) *SendVoice {
	s.client = c
	return s
}

func (s SendVoice) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendVoice) WithApiBaseUrl(u string) *SendVoice {
	s.baseUrl = u
	return s
}

func (s SendVoice) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendVoice) Execute(token string) (*objects.Message, error) {
	if s.Voice.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendVoice", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendVoice", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendVideoNote) WithClient(c *http.Client) *SendVideoNote {
	s.client = c
	return s
}

func (s SendVideoNote) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendVideoNote) WithApiBaseUrl(u string) *SendVideoNote {
	s.baseUrl = u
	return s
}

func (s SendVideoNote) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendVideoNote) Execute(token string) (*objects.Message, error) {
	if s.VideoNote.IsLocal() {
		return SendTelegramMultipartRequest[objects.Message](token, "sendVideoNote", s)
	}
	return SendTelegramPostRequest[objects.Message](token, "sendVideoNote", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendPaidMedia) WithClient(c *http.Client) *SendPaidMedia {
	s.client = c
	return s
}

func (s SendPaidMedia) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendPaidMedia) WithApiBaseUrl(u string) *SendPaidMedia {
	s.baseUrl = u
	return s
}

func (s SendPaidMedia) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendPaidMedia) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendPaidMedia", s)
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
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent messages from forwarding and saving
	ProtectContent *bool `json:"protect_content,omitempty"`
	// /Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message.
	//The relevant Stars will be withdrawn from the bot's balance
	AllowPaidBroadcast *bool `json:"allow_paid_broadcast,omitempty"`
	//Unique identifier of the message effect to be added to the message; for private chats only
	MessageEffectId *string `json:"message_effect_id,omitempty"`
	//Description of the message to reply to
	ReplyParameters *objects.ReplyParameters `json:"reply_parameters,omitempty"`
	client          *http.Client
	baseUrl         string
}

func (s *SendMediaGroup) WithClient(c *http.Client) *SendMediaGroup {
	s.client = c
	return s
}

func (s SendMediaGroup) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendMediaGroup) WithApiBaseUrl(u string) *SendMediaGroup {
	s.baseUrl = u
	return s
}

func (s SendMediaGroup) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendMediaGroup) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)

	if s.BusinessConnectionId != nil {
		if err := w.WriteField("business_connection_id", *s.BusinessConnectionId); err != nil {
			return nil, nil, err
		}
	}
	if err := w.WriteField("chat_id", s.ChatId); err != nil {
		return nil, nil, err
	}
	if s.MessageThreadId != nil {
		if err := w.WriteField("message_thread_id", *s.MessageThreadId); err != nil {
			return nil, nil, err
		}
	}
	by, err := json.Marshal(s.Media)
	if err != nil {
		return nil, nil, err
	}
	if err := w.WriteField("media", string(by)); err != nil {
		return nil, nil, err
	}
	if s.DisableNotification != nil {
		if err := w.WriteField("disable_notification", fmt.Sprint(s.DisableNotification)); err != nil {
			return nil, nil, err
		}
	}
	if s.ProtectContent != nil {
		if err := w.WriteField("protect_content", fmt.Sprint(s.ProtectContent)); err != nil {
			return nil, nil, err
		}
	}
	if s.AllowPaidBroadcast != nil {
		if err := w.WriteField("allow_paid_broadcast", fmt.Sprint(s.AllowPaidBroadcast)); err != nil {
			return nil, nil, err
		}
	}
	if s.MessageEffectId != nil {
		if err := w.WriteField("message_effect_id", *s.MessageEffectId); err != nil {
			return nil, nil, err
		}
	}
	if s.ReplyParameters != nil {
		by, err := json.Marshal(s.ReplyParameters)
		if err != nil {
			return nil, nil, err
		}
		if err := w.WriteField("reply_parameters", string(by)); err != nil {
			return nil, nil, err
		}
	}

	for _, media := range s.Media {
		if media.IsLocalFile() {
			part, err := w.CreateFormFile(media.Deattach(), media.Deattach())
			if err != nil {
				return nil, nil, err
			}
			if _, err := io.Copy(part, media.GetReader()); err != nil {
				return nil, nil, err
			}
		}
	}
	w.Close()
	return b, w, nil
}

func (s SendMediaGroup) Execute(token string) (*[]objects.Message, error) {
	return SendTelegramMultipartRequest[[]objects.Message](token, "sendMediaGroup", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendLocation) WithClient(c *http.Client) *SendLocation {
	s.client = c
	return s
}

func (s SendLocation) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendLocation) WithApiBaseUrl(u string) *SendLocation {
	s.baseUrl = u
	return s
}

func (s SendLocation) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendLocation) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendLocation", s)
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
	client      *http.Client
	baseUrl     string
}

func (s *SendVenue) WithClient(c *http.Client) *SendVenue {
	s.client = c
	return s
}

func (s SendVenue) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendVenue) WithApiBaseUrl(u string) *SendVenue {
	s.baseUrl = u
	return s
}

func (s SendVenue) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendVenue) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendVenue", s)
}

// Use this method to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//Contact's phone number
	PhoneNumber string `json:"phone_number"`
	//Contact's first name
	FirstName string `json:"first_name"`
	//Contact's last name
	LastName *string `json:"last_name,omitempty"`
	//Additional data about the contact in the form of a vCard, 0-2048 bytes
	Vcard *string `json:"vcard,omitempty"`
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
	client      *http.Client
	baseUrl     string
}

func (s *SendContact) WithClient(c *http.Client) *SendContact {
	s.client = c
	return s
}

func (s SendContact) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendContact) WithApiBaseUrl(u string) *SendContact {
	s.baseUrl = u
	return s
}

func (s SendContact) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendContact) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendContact", s)
}

// Use this method to send a native poll. On success, the sent Message is returned.
type SendPoll struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//Poll question, 1-300 characters
	Question string `json:"question"`
	//Mode for parsing entities in the question.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	//Currently, only custom emoji entities are allowed
	QuestionParseMode *string `json:"question_parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the poll question.
	//It can be specified instead of question_parse_mode
	QuestionEntities *[]objects.MessageEntity `json:"question_entities,omitempty"`
	//A JSON-serialized list of 2-10 answer options
	Options []objects.InputPollOption `json:"options"`
	//True, if the poll needs to be anonymous, defaults to True
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	//Poll type, “quiz” or “regular”, defaults to “regular”
	Type *string `json:"type,omitempty"`
	//True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	AllowMultipleAnswers *bool `json:"allow_multiple_answers,omitempty"`
	//0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionId *int `json:"correct_option_id,omitempty"`
	//Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll,
	//0-200 characters with at most 2 line feeds after entities parsing
	Explanation *string `json:"explanation,omitempty"`
	//Mode for parsing entities in the explanation.
	//See https://core.telegram.org/bots/api#formatting-options for more details.
	ExplanationParseMode *string `json:"explanation_parse_mode,omitempty"`
	//A JSON-serialized list of special entities that appear in the poll explanation.
	//It can be specified instead of explanation_parse_mode
	ExplanationEntities *[]objects.MessageEntity `json:"explanation_entities,omitempty"`
	//Amount of time in seconds the poll will be active after creation, 5-600. Can't be used together with close_date.
	OpenPeriod *int `json:"open_period,omitempty"`
	//Point in time (Unix timestamp) when the poll will be automatically closed. Must be at least 5 and no more than 600 seconds in the future.
	//Can't be used together with open_period.
	CloseDate *int `json:"close_date,omitempty"`
	//Pass True if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed *bool `json:"is_closed,omitempty"`
	// /Sends the message silently. Users will receive a notification with no sound.
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
	client      *http.Client
	baseUrl     string
}

func (s *SendPoll) WithClient(c *http.Client) *SendPoll {
	s.client = c
	return s
}

func (s SendPoll) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendPoll) WithApiBaseUrl(u string) *SendPoll {
	s.baseUrl = u
	return s
}

func (s SendPoll) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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
	if s.Type != nil {
		if *s.Type != "quiz" && *s.Type != "regular" {
			return objects.ErrInvalidParam("type parameter must be 'regular' or 'quiz' if specified")
		}
	}
	return nil
}

func (s SendPoll) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPoll) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendPoll", s)
}

// Use this method to send an animated emoji that will display a random value. On success, the sent Message is returned.
type SendDice struct {
	//Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	//Emoji on which the dice throw animation is based. Currently, must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”.
	//Dice can have values 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64 for “🎰”.
	//Defaults to “🎲”
	Emoji string `json:"emoji"`
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	//Protects the contents of the sent message from forwarding
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
	client      *http.Client
	baseUrl     string
}

func (s *SendDice) WithClient(c *http.Client) *SendDice {
	s.client = c
	return s
}

func (s SendDice) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendDice) WithApiBaseUrl(u string) *SendDice {
	s.baseUrl = u
	return s
}

func (s SendDice) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SendDice) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendPoll", s)
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
	//Unique identifier of the business connection on behalf of which the action will be sent
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread; for supergroups only
	Action string `json:"action"`
	//Type of action to broadcast. Choose one, depending on what the user is about to receive: 'typing' for text messages,
	//'upload_photo' for photos, 'record_video' or 'upload_video' for videos, 'record_voice' or 'upload_voice' for voice notes,
	//'upload_document' for general files, 'choose_sticker' for stickers, 'find_location' for location data,
	//'record_video_note' or 'upload_video_note' for video notes.
	MessageThreadId *string `json:"message_thread_id,omitempty"`
	client          *http.Client
	baseUrl         string
}

func (s *SendChatAction) WithClient(c *http.Client) *SendChatAction {
	s.client = c
	return s
}

func (s SendChatAction) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendChatAction) WithApiBaseUrl(u string) *SendChatAction {
	s.baseUrl = u
	return s
}

func (s SendChatAction) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s SendChatAction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.Action) == "" {
		return objects.ErrInvalidParam("action parameter can't be empty")
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
		return objects.ErrInvalidParam("invalid action parameter: see https://core.telegram.org/bots/api#sendchataction for a list of available actions")
	}
	return nil
}

func (s SendChatAction) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendChatAction) Execute(token string) (*bool, error) {
	return SendTelegramGetRequest[bool](token, "sendChatAction", s)
	//TODO: make a function that will be a part of sendable interface and will return endpoint of this method as a string
	//name should be Endpoint()
	//then there would be no need to pass endpoint to any of those functions
	//function signature would be smaller and easier
}

// Use this method to change the chosen reactions on a message.
// Service messages can't be reacted to.
// Automatically forwarded messages from a channel to its discussion group have the same available reactions as messages in the channel.
// Bots can't use paid reactions. Returns True on success.
type SetMessageReaction struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Identifier of the target message. If the message belongs to a media group, the reaction is set to the first non-deleted message in the group instead.
	MessageId int `json:"message_id"`
	//A JSON-serialized list of reaction types to set on the message.
	//Currently, as non-premium users, bots can set up to one reaction per message.
	//A custom emoji reaction can be used if it is either already present on the message or explicitly allowed by chat administrators.
	//Paid reactions can't be used by bots.
	Reaction *[]objects.ReactionType `json:"reaction,omitempty"`
	//Pass True to set the reaction with a big animation
	IsBig   *bool `json:"is_big,omitempty"`
	client  *http.Client
	baseUrl string
}

func (s *SetMessageReaction) WithClient(c *http.Client) *SetMessageReaction {
	s.client = c
	return s
}

func (s SetMessageReaction) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMessageReaction) WithApiBaseUrl(u string) *SetMessageReaction {
	s.baseUrl = u
	return s
}

func (s SetMessageReaction) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s SetMessageReaction) Validate() error {
	if strings.TrimSpace(s.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if s.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	for _, r := range *s.Reaction {
		if err := r.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetMessageReaction) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMessageReaction) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMessageReaction", s)
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset *int `json:"offset,omitempty"`
	//Limits the number of photos to be retrieved. Values between 1-100 are accepted. Defaults to 100.
	Limit   *int `json:"limit,omitempty"`
	client  *http.Client
	baseUrl string
}

func (s *GetUserProfilePhotos) WithClient(c *http.Client) *GetUserProfilePhotos {
	s.client = c
	return s
}

func (s GetUserProfilePhotos) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetUserProfilePhotos) WithApiBaseUrl(u string) *GetUserProfilePhotos {
	s.baseUrl = u
	return s
}

func (s GetUserProfilePhotos) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetUserProfilePhotos) Execute(token string) (*objects.UserProfilePhotos, error) {
	return SendTelegramGetRequest[objects.UserProfilePhotos](token, "getUserProfilePhotos", g)
}

// Changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the
// Mini App method requestEmojiStatusAccess. Returns True on success.
type SetUserEmojiStatus struct {
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//Custom emoji identifier of the emoji status to set. Pass an empty string to remove the status.
	EmojiStatusCustomEmojiId *string `json:"emoji_status_custom_emoji_id,omitempty"`
	//Expiration date of the emoji status, if any
	EmojiStatusExpirationDate *int `json:"emoji_status_expiration_date,omitempty"`
	client                    *http.Client
	baseUrl                   string
}

func (s *SetUserEmojiStatus) WithClient(c *http.Client) *SetUserEmojiStatus {
	s.client = c
	return s
}

func (s SetUserEmojiStatus) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetUserEmojiStatus) WithApiBaseUrl(u string) *SetUserEmojiStatus {
	s.baseUrl = u
	return s
}

func (s SetUserEmojiStatus) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetUserEmojiStatus) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setUserEmojiStatus", s)
}

// Use this method to get basic information about a file and prepare it for downloading.
// For the moment, bots can download files of up to 20MB in size. On success, a File object is returned.
// The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>,
// where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile again.
//
// Note: This function may not preserve the original file name and MIME type.
// You should save the file's MIME type and name (if available) when the File object is received.
type GetFile struct {
	//File identifier to get information about
	FileId  string `json:"file_id"`
	client  *http.Client
	baseUrl string
} //TODO: probably a good idea to add a method that will download the file and will return it as reader

func (s *GetFile) WithClient(c *http.Client) *GetFile {
	s.client = c
	return s
}

func (s GetFile) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetFile) WithApiBaseUrl(u string) *GetFile {
	s.baseUrl = u
	return s
}

func (s GetFile) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetFile) Execute(token string) (*objects.File, error) {
	return SendTelegramGetRequest[objects.File](token, "getFile", g)
}

// Use this method to ban a user in a group, a supergroup or a channel.
// In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links,
// etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type BanChatMember struct {
	//Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//Date when the user will be unbanned; Unix time.
	//If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever.
	//Applied for supergroups and channels only.
	UntilDate *int `json:"until_date,omitempty"`
	//Pass True to delete all messages from the chat for the user that is being removed.
	//If False, the user will be able to see messages in the group that were sent before the user was removed.
	//Always True for supergroups and channels.
	RevokeMessages *bool `json:"revoke_messages,omitempty"`
	client         *http.Client
	baseUrl        string
}

func (s *BanChatMember) WithClient(c *http.Client) *BanChatMember {
	s.client = c
	return s
}

func (s BanChatMember) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *BanChatMember) WithApiBaseUrl(u string) *BanChatMember {
	s.baseUrl = u
	return s
}

func (s BanChatMember) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (b BanChatMember) Execute(token string) (*bool, error) {
	return SendTelegramGetRequest[bool](token, "banChatMember", b)
}

// Use this method to unban a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be able to join via link, etc.
// The bot must be an administrator for this to work. By default,
// this method guarantees that after the call the user is not a member of the chat, but will be able to join it.
// So if the user is a member of the chat they will also be removed from the chat. If you don't want this, use the parameter only_if_banned.
// Returns True on success.
type UnbanChatMember struct {
	//Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//Do nothing if the user is not banned
	OnlyIfBanned *bool `json:"only_if_banned,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *UnbanChatMember) WithClient(c *http.Client) *UnbanChatMember {
	s.client = c
	return s
}

func (s UnbanChatMember) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnbanChatMember) WithApiBaseUrl(u string) *UnbanChatMember {
	s.baseUrl = u
	return s
}

func (s UnbanChatMember) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (b UnbanChatMember) Execute(token string) (*bool, error) {
	return SendTelegramGetRequest[bool](token, "unbanChatMember", b)
}

// Use this method to restrict a user in a supergroup.
// The bot must be an administrator in the supergroup for this to work and must have the appropriate administrator rights.
// Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//A JSON-serialized object for new user permissions
	Permissions objects.ChatPermissions `json:"permissions"`
	//Pass True if chat permissions are set independently.
	//Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	//can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	//the can_send_polls permission will imply the can_send_messages permission.
	UserIndependentChatPermissions *bool `json:"user_independent_chat_permissions,omitempty"`
	//Date when restrictions will be lifted for the user; Unix time.
	//If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilDate *int `json:"until_date,omitempty"`
	client    *http.Client
	baseUrl   string
}

func (s *RestrictChatMember) WithClient(c *http.Client) *RestrictChatMember {
	s.client = c
	return s
}

func (s RestrictChatMember) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *RestrictChatMember) WithApiBaseUrl(u string) *RestrictChatMember {
	s.baseUrl = u
	return s
}

func (s RestrictChatMember) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (r RestrictChatMember) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "restrictChatMember", r)
}

// Use this method to promote or demote a user in a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Pass False for all boolean parameters to demote a user. Returns True on success
type PromoteChatMember struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//Pass True if the administrator's presence in the chat is hidden
	IsAnonymous *bool `json:"is_anonymous,omitempty"`
	//Pass True if the administrator can access the chat event log, get boost list,
	//see hidden supergroup and channel members, report spam messages and ignore slow mode.
	//Implied by any other administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty"`
	//Pass True if the administrator can delete messages of other users
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`
	//Pass True if the administrator can manage video chats
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`
	//Pass True if the administrator can restrict, ban or unban chat members, or access supergroup statistics
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`
	//Pass True if the administrator can add new administrators with a subset of their own privileges or
	//demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him)
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`
	//Pass True if the administrator can change chat title, photo and other settings
	CanChangeInfo *bool `json:"can_change_info,omitempty"`
	//Pass True if the administrator can invite new users to the chat
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`
	//Pass True if the administrator can post stories to the chat
	CanPostStories *bool `json:"can_post_stories,omitempty"`
	//Pass True if the administrator can edit stories posted by other users,
	//post stories to the chat page, pin chat stories, and access the chat's story archive
	CanEditStories *bool `json:"can_edit_stories,omitempty"`
	//Pass True if the administrator can delete stories posted by other users
	CanDeleteStories *bool `json:"can_delete_stories,omitempty"`
	//Pass True if the administrator can post messages in the channel, or access channel statistics; for channels only
	CanPostMessages *bool `json:"can_post_messages,omitempty"`
	//Pass True if the administrator can edit messages of other users and can pin messages; for channels only
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`
	//Pass True if the administrator can pin messages; for supergroups only
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`
	//Pass True if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
	client          *http.Client
	baseUrl         string
}

func (s *PromoteChatMember) WithClient(c *http.Client) *PromoteChatMember {
	s.client = c
	return s
}

func (s PromoteChatMember) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *PromoteChatMember) WithApiBaseUrl(u string) *PromoteChatMember {
	s.baseUrl = u
	return s
}

func (s PromoteChatMember) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p PromoteChatMember) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "promoteChatMember", p)
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitle struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId int `json:"user_id"`
	//New custom title for the administrator; 0-16 characters, emoji are not allowed
	CustomTitle string `json:"custom_title"`
	client      *http.Client
	baseUrl     string
}

func (s *SetChatAdministratorCustomTitle) WithClient(c *http.Client) *SetChatAdministratorCustomTitle {
	s.client = c
	return s
}

func (s SetChatAdministratorCustomTitle) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatAdministratorCustomTitle) WithApiBaseUrl(u string) *SetChatAdministratorCustomTitle {
	s.baseUrl = u
	return s
}

func (s SetChatAdministratorCustomTitle) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatAdministratorCustomTitle) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatAdministratorCustomTitle", s)
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the chat is unbanned,
// the owner of the banned chat won't be able to send messages on behalf of any of their channels.
// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights.
// Returns True on success.
type BanChatSenderChat struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target sender chat
	SenderChatId int `json:"sender_chat_id"`
	client       *http.Client
	baseUrl      string
}

func (s *BanChatSenderChat) WithClient(c *http.Client) *BanChatSenderChat {
	s.client = c
	return s
}

func (s BanChatSenderChat) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *BanChatSenderChat) WithApiBaseUrl(u string) *BanChatSenderChat {
	s.baseUrl = u
	return s
}

func (s BanChatSenderChat) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (b BanChatSenderChat) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "banChatSenderChat", b)
}

// Use this method to unban a previously banned channel chat in a supergroup or channel.
// The bot must be an administrator for this to work and must have the appropriate administrator rights.
// Returns True on success.
type UnbanChatSenderChat struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target sender chat
	SenderChatId int `json:"sender_chat_id"`
	client       *http.Client
	baseUrl      string
}

func (s *UnbanChatSenderChat) WithClient(c *http.Client) *UnbanChatSenderChat {
	s.client = c
	return s
}

func (s UnbanChatSenderChat) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnbanChatSenderChat) WithApiBaseUrl(u string) *UnbanChatSenderChat {
	s.baseUrl = u
	return s
}

func (s UnbanChatSenderChat) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (b UnbanChatSenderChat) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unbanChatSenderChat", b)
}

// Use this method to set default chat permissions for all members.
// The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members administrator rights.
// Returns True on success.
type SetChatPermissions struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//A JSON-serialized object for new default chat permissions
	Permissions objects.ChatPermissions `json:"permissions"`
	//Pass True if chat permissions are set independently.
	//Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the
	//can_send_messages, can_send_audios, can_send_documents, can_send_photos,
	// can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	//the can_send_polls permission will imply the can_send_messages permission.
	UserIndependentChatPermissions *bool `json:"user_independent_chat_permissions"`
	client                         *http.Client
	baseUrl                        string
}

func (s *SetChatPermissions) WithClient(c *http.Client) *SetChatPermissions {
	s.client = c
	return s
}

func (s SetChatPermissions) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatPermissions) WithApiBaseUrl(u string) *SetChatPermissions {
	s.baseUrl = u
	return s
}

func (s SetChatPermissions) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatPermissions) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatPermissions", s)
}

// Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated by other administrators.
// If you want your bot to work with invite links, it will need to generate its own link using exportChatInviteLink or by calling the getChat method.
// If your bot needs to generate a new primary invite link replacing its previous one, use exportChatInviteLink again.
type ExportChatInviteLink struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *ExportChatInviteLink) WithClient(c *http.Client) *ExportChatInviteLink {
	s.client = c
	return s
}

func (s ExportChatInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ExportChatInviteLink) WithApiBaseUrl(u string) *ExportChatInviteLink {
	s.baseUrl = u
	return s
}

func (s ExportChatInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e ExportChatInviteLink) Execute(token string) (*string, error) {
	return SendTelegramPostRequest[string](token, "exportChatInviteLink", e)
}

// Use this method to create an additional invite link for a chat.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// The link can be revoked using the method revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
type CreateInviteLink struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	//Point in time (Unix timestamp) when the link will expire
	ExpireDate *int `json:"expire_date,omitempty"`
	//The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int `json:"member_limit,omitempty"`
	//True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty"`
	client             *http.Client
	baseUrl            string
}

func (s *CreateInviteLink) WithClient(c *http.Client) *CreateInviteLink {
	s.client = c
	return s
}

func (s CreateInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CreateInviteLink) WithApiBaseUrl(u string) *CreateInviteLink {
	s.baseUrl = u
	return s
}

func (s CreateInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CreateInviteLink) Execute(token string) (*objects.ChatInviteLink, error) {
	return SendTelegramPostRequest[objects.ChatInviteLink](token, "createInviteLink", c)
}

// Use this method to edit a non-primary invite link created by the bot.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the edited invite link as a ChatInviteLink object.
type EditChatInviteLink struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//The invite link to edit
	InviteLink string `json:"invite_link"`
	//Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	//Point in time (Unix timestamp) when the link will expire
	ExpireDate *int `json:"expire_date,omitempty"`
	//The maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit *int `json:"member_limit,omitempty"`
	//True, if users joining the chat via the link need to be approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest *bool `json:"creates_join_request,omitempty"`
	client             *http.Client
	baseUrl            string
}

func (s *EditChatInviteLink) WithClient(c *http.Client) *EditChatInviteLink {
	s.client = c
	return s
}

func (s EditChatInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditChatInviteLink) WithApiBaseUrl(u string) *EditChatInviteLink {
	s.baseUrl = u
	return s
}

func (s EditChatInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (c EditChatInviteLink) Validate() error {
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

func (c EditChatInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditChatInviteLink) Execute(token string) (*objects.ChatInviteLink, error) {
	return SendTelegramPostRequest[objects.ChatInviteLink](token, "editInviteLink", c)
}

// Use this method to create a subscription invite link for a channel chat.
// The bot must have the can_invite_users administrator rights.
// The link can be edited using the method editChatSubscriptionInviteLink or revoked using the method revokeChatInviteLink.
// Returns the new invite link as a ChatInviteLink object.
type CreateChatSubscriptionInviteLink struct {
	//Unique identifier for the target channel chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Invite link name; 0-32 characters
	Name *string `json:"name,omitempty"`
	//The number of seconds the subscription will be active for before the next payment. Currently, it must always be 2592000 (30 days).
	SubscriptionPeriod int `json:"subscription_period"`
	//The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat; 1-2500
	SubscriptionPrice int `json:"subscription_price"`
	client            *http.Client
	baseUrl           string
}

func (s *CreateChatSubscriptionInviteLink) WithClient(c *http.Client) *CreateChatSubscriptionInviteLink {
	s.client = c
	return s
}

func (s CreateChatSubscriptionInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CreateChatSubscriptionInviteLink) WithApiBaseUrl(u string) *CreateChatSubscriptionInviteLink {
	s.baseUrl = u
	return s
}

func (s CreateChatSubscriptionInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (c CreateChatSubscriptionInviteLink) Execute(token string) (*objects.ChatInviteLink, error) {
	return SendTelegramPostRequest[objects.ChatInviteLink](token, "createChatSubscriptionInviteLink", c)
}

// Use this method to edit a subscription invite link created by the bot.
// The bot must have the can_invite_users administrator rights.
// Returns the edited invite link as a ChatInviteLink object.
type EditChatSubscriptionInviteLink struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//The invite link to edit
	InviteLink string `json:"invite_link"`
	//Invite link name; 0-32 characters
	Name    *string `json:"name,omitempty"`
	client  *http.Client
	baseUrl string
}

func (s *EditChatSubscriptionInviteLink) WithClient(c *http.Client) *EditChatSubscriptionInviteLink {
	s.client = c
	return s
}

func (s EditChatSubscriptionInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditChatSubscriptionInviteLink) WithApiBaseUrl(u string) *EditChatSubscriptionInviteLink {
	s.baseUrl = u
	return s
}

func (s EditChatSubscriptionInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters if specified")
		}
	}
	return nil
}

func (c EditChatSubscriptionInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditChatSubscriptionInviteLink) Execute(token string) (*objects.ChatInviteLink, error) {
	return SendTelegramPostRequest[objects.ChatInviteLink](token, "editChatSubscriptionInviteLink", c)
}

// Use this method to revoke an invite link created by the bot.
// If the primary link is revoked, a new link is automatically generated.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns the revoked invite link as ChatInviteLink object.
type RevokeInviteLink struct {
	//nique identifier of the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//The invite link to revoke
	InviteLink string `json:"invite_link"`
	client     *http.Client
	baseUrl    string
}

func (s *RevokeInviteLink) WithClient(c *http.Client) *RevokeInviteLink {
	s.client = c
	return s
}

func (s RevokeInviteLink) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *RevokeInviteLink) WithApiBaseUrl(u string) *RevokeInviteLink {
	s.baseUrl = u
	return s
}

func (s RevokeInviteLink) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (c RevokeInviteLink) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(c.InviteLink) == "" {
		return objects.ErrInvalidParam("invite_link parameter can't be empty")
	}
	return nil
}

func (c RevokeInviteLink) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c RevokeInviteLink) Execute(token string) (*objects.ChatInviteLink, error) {
	return SendTelegramPostRequest[objects.ChatInviteLink](token, "revokeInviteLink", c)
}

// Use this method to approve a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
type ApproveChatJoinRequest struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId  int `json:"user_id"`
	client  *http.Client
	baseUrl string
}

func (s *ApproveChatJoinRequest) WithClient(c *http.Client) *ApproveChatJoinRequest {
	s.client = c
	return s
}

func (s ApproveChatJoinRequest) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ApproveChatJoinRequest) WithApiBaseUrl(u string) *ApproveChatJoinRequest {
	s.baseUrl = u
	return s
}

func (s ApproveChatJoinRequest) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s ApproveChatJoinRequest) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "approveChatJoinRequest", s)
}

// Use this method to decline a chat join request.
// The bot must be an administrator in the chat for this to work and must have the can_invite_users administrator right.
// Returns True on success.
type DeclineChatJoinRequest struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId  int `json:"user_id"`
	client  *http.Client
	baseUrl string
}

func (s *DeclineChatJoinRequest) WithClient(c *http.Client) *DeclineChatJoinRequest {
	s.client = c
	return s
}

func (s DeclineChatJoinRequest) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeclineChatJoinRequest) WithApiBaseUrl(u string) *DeclineChatJoinRequest {
	s.baseUrl = u
	return s
}

func (s DeclineChatJoinRequest) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s DeclineChatJoinRequest) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "declineChatJoinRequest", s)
}

type SetChatPhoto struct {
	ChatId  string            `json:"chat_id"`
	Photo   objects.InputFile `json:"photo"`
	client  *http.Client
	baseUrl string
}

func (s *SetChatPhoto) WithClient(c *http.Client) *SetChatPhoto {
	s.client = c
	return s
}

func (s SetChatPhoto) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatPhoto) WithApiBaseUrl(u string) *SetChatPhoto {
	s.baseUrl = u
	return s
}

func (s SetChatPhoto) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatPhoto) ToMultipartBody() (*bytes.Buffer, *multipart.Writer, error) {
	if _, ok := s.Photo.(objects.InputFileFromRemote); ok {
		return nil, nil, fmt.Errorf("can't use remote file when setting chat photo; only local files are supported")
	}

	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	if err := w.WriteField("chat_id", s.ChatId); err != nil {
		return nil, nil, err
	}

	part, err := w.CreateFormFile("photo", s.Photo.Name())
	if err != nil {
		return nil, nil, err
	}
	reader, err := s.Photo.Reader()
	if err != nil {
		return nil, nil, err
	}
	if _, err = io.Copy(part, reader); err != nil {
		return nil, nil, err
	}
	w.Close()
	return buf, w, nil
}

func (s SetChatPhoto) Execute(token string) (*bool, error) {
	return SendTelegramMultipartRequest[bool](token, "setChatPhoto", s)
}

// Use this method to delete a chat photo. Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type DeleteChatPhoto struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *DeleteChatPhoto) WithClient(c *http.Client) *DeleteChatPhoto {
	s.client = c
	return s
}

func (s DeleteChatPhoto) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteChatPhoto) WithApiBaseUrl(u string) *DeleteChatPhoto {
	s.baseUrl = u
	return s
}

func (s DeleteChatPhoto) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (d DeleteChatPhoto) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteChatPhoto", d)
}

// Use this method to change the title of a chat. Titles can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatTitle struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//New chat title, 1-128 characters
	Title   string `json:"title"`
	client  *http.Client
	baseUrl string
}

func (s *SetChatTitle) WithClient(c *http.Client) *SetChatTitle {
	s.client = c
	return s
}

func (s SetChatTitle) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatTitle) WithApiBaseUrl(u string) *SetChatTitle {
	s.baseUrl = u
	return s
}

func (s SetChatTitle) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatTitle) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatTitle", s)
}

// Use this method to change the description of a group, a supergroup or a channel.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Returns True on success.
type SetChatDescription struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//New chat description, 0-255 characters
	Description string `json:"description"`
	client      *http.Client
	baseUrl     string
}

func (s *SetChatDescription) WithClient(c *http.Client) *SetChatDescription {
	s.client = c
	return s
}

func (s SetChatDescription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatDescription) WithApiBaseUrl(u string) *SetChatDescription {
	s.baseUrl = u
	return s
}

func (s SetChatDescription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatDescription) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatTitle", s)
}

// Use this method to add a message to the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type PinChatMessage struct {
	//Unique identifier of the business connection on behalf of which the message will be pinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Identifier of a message to pin
	MessageId int `json:"message_id"`
	//Pass True if it is not necessary to send a notification to all chat members about the new pinned message.
	//Notifications are always disabled in channels and private chats.
	DisableNotification *bool `json:"disable_notification,omitempty"`
	client              *http.Client
	baseUrl             string
}

func (s *PinChatMessage) WithClient(c *http.Client) *PinChatMessage {
	s.client = c
	return s
}

func (s PinChatMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *PinChatMessage) WithApiBaseUrl(u string) *PinChatMessage {
	s.baseUrl = u
	return s
}

func (s PinChatMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p PinChatMessage) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "pinChatMessage", p)
}

// Use this method to remove a message from the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type UnpinChatMessage struct {
	//Unique identifier of the business connection on behalf of which the message will be unpinned
	BusinessConnectionId *string `json:"business_connection_id,omitempty"`
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Identifier of the message to unpin. Required if business_connection_id is specified. If not specified, the most recent pinned message (by sending date) will be unpinned.
	MessageId *int `json:"message_id,omitempty"`
	client    *http.Client
	baseUrl   string
}

func (s *UnpinChatMessage) WithClient(c *http.Client) *UnpinChatMessage {
	s.client = c
	return s
}

func (s UnpinChatMessage) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnpinChatMessage) WithApiBaseUrl(u string) *UnpinChatMessage {
	s.baseUrl = u
	return s
}

func (s UnpinChatMessage) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (p UnpinChatMessage) Validate() error {
	if strings.TrimSpace(p.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if p.MessageId != nil {
		if *p.MessageId < 1 {
			return objects.ErrInvalidParam("message_id parameter can't be empty")
		}
	}

	return nil
}

func (p UnpinChatMessage) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinChatMessage) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unpinChatMessage", p)
}

// Use this method to clear the list of pinned messages in a chat.
// If the chat is not a private chat, the bot must be an administrator in the chat for this to work and
// must have the 'can_pin_messages' administrator right in a supergroup or 'can_edit_messages' administrator right in a channel.
// Returns True on success.
type UnpinAllChatMessages struct {
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *UnpinAllChatMessages) WithClient(c *http.Client) *UnpinAllChatMessages {
	s.client = c
	return s
}

func (s UnpinAllChatMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnpinAllChatMessages) WithApiBaseUrl(u string) *UnpinAllChatMessages {
	s.baseUrl = u
	return s
}

func (s UnpinAllChatMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p UnpinAllChatMessages) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unpinAllChatMessages", p)
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
	//Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *LeaveChat) WithClient(c *http.Client) *LeaveChat {
	s.client = c
	return s
}

func (s LeaveChat) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *LeaveChat) WithApiBaseUrl(u string) *LeaveChat {
	s.baseUrl = u
	return s
}

func (s LeaveChat) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p LeaveChat) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "leaveChat", p)
}

// Use this method to get up-to-date information about the chat. Returns a ChatFullInfo object on success.
type GetChat struct {
	//Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *GetChat) WithClient(c *http.Client) *GetChat {
	s.client = c
	return s
}

func (s GetChat) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetChat) WithApiBaseUrl(u string) *GetChat {
	s.baseUrl = u
	return s
}

func (s GetChat) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p GetChat) Execute(token string) (*objects.ChatFullInfo, error) {
	return SendTelegramGetRequest[objects.ChatFullInfo](token, "getChat", p)
}

// Use this method to get a list of administrators in a chat, which aren't bots. Returns an Array of ChatMember objects.
type GetChatAdministrators struct {
	//Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *GetChatAdministrators) WithClient(c *http.Client) *GetChatAdministrators {
	s.client = c
	return s
}

func (s GetChatAdministrators) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetChatAdministrators) WithApiBaseUrl(u string) *GetChatAdministrators {
	s.baseUrl = u
	return s
}

func (s GetChatAdministrators) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p GetChatAdministrators) Execute(token string) (*[]objects.ChatMember, error) {
	return SendTelegramGetRequest[[]objects.ChatMember](token, "getChatAdministrators", p)
}

// Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMemberCount struct {
	//Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *GetChatMemberCount) WithClient(c *http.Client) *GetChatMemberCount {
	s.client = c
	return s
}

func (s GetChatMemberCount) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetChatMemberCount) WithApiBaseUrl(u string) *GetChatMemberCount {
	s.baseUrl = u
	return s
}

func (s GetChatMemberCount) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p GetChatMemberCount) Execute(token string) (*int, error) {
	return SendTelegramGetRequest[int](token, "getChatMemberCount", p)
}

// Use this method to get information about a member of a chat.
// The method is only guaranteed to work for other users if the bot is an administrator in the chat.
// Returns a ChatMember object on success.
type GetChatMember struct {
	//Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId  int `json:"user_id"`
	client  *http.Client
	baseUrl string
}

func (s *GetChatMember) WithClient(c *http.Client) *GetChatMember {
	s.client = c
	return s
}

func (s GetChatMember) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetChatMember) WithApiBaseUrl(u string) *GetChatMember {
	s.baseUrl = u
	return s
}

func (s GetChatMember) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p GetChatMember) Execute(token string) (*objects.ChatMember, error) {
	return SendTelegramGetRequest[objects.ChatMember](token, "getChatMember", p)
}

// Use this method to set a new group sticker set for a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True on success.
type SetChatStickerSet struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Name of the sticker set to be set as the group sticker set
	StickerSetName string `json:"sticker_set_name"`
	client         *http.Client
	baseUrl        string
}

func (s *SetChatStickerSet) WithClient(c *http.Client) *SetChatStickerSet {
	s.client = c
	return s
}

func (s SetChatStickerSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatStickerSet) WithApiBaseUrl(u string) *SetChatStickerSet {
	s.baseUrl = u
	return s
}

func (s SetChatStickerSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p SetChatStickerSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatStickerSet", p)
}

// Use this method to delete a group sticker set from a supergroup.
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.
// Returns True on success.
type DeleteChatStickerSet struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *DeleteChatStickerSet) WithClient(c *http.Client) *DeleteChatStickerSet {
	s.client = c
	return s
}

func (s DeleteChatStickerSet) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteChatStickerSet) WithApiBaseUrl(u string) *DeleteChatStickerSet {
	s.baseUrl = u
	return s
}

func (s DeleteChatStickerSet) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (p DeleteChatStickerSet) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteChatStickerSet", p)
}

// Use this method to get custom emoji stickers, which can be used as a forum topic icon by any user.
// Requires no parameters. Returns an Array of Sticker objects.
type GetForumTopicIconStickers struct {
	client  *http.Client
	baseUrl string
}

func (s *GetForumTopicIconStickers) WithClient(c *http.Client) *GetForumTopicIconStickers {
	s.client = c
	return s
}

func (s GetForumTopicIconStickers) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetForumTopicIconStickers) WithApiBaseUrl(u string) *GetForumTopicIconStickers {
	s.baseUrl = u
	return s
}

func (s GetForumTopicIconStickers) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

// always nil
func (g GetForumTopicIconStickers) Validate() error {
	return nil
}

// alwways empty json
func (g GetForumTopicIconStickers) ToRequestBody() ([]byte, error) {
	return json.Marshal(struct{}{})
}

func (g GetForumTopicIconStickers) Execute(token string) (*[]objects.Sticker, error) {
	return SendTelegramGetRequest[[]objects.Sticker](token, "getForumTopicStickers", g)
}

// Use this method to create a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns information about the created topic as a ForumTopic object.
type CreateForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Topic name, 1-128 characters
	Name string `json:"name"`
	//Color of the topic icon in RGB format.
	//Currently, must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)
	IconColor *int `json:"icon_color,omitempty"`
	//Unique identifier of the custom emoji shown as the topic icon. Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
	client            *http.Client
	baseUrl           string
}

func (s *CreateForumTopic) WithClient(c *http.Client) *CreateForumTopic {
	s.client = c
	return s
}

func (s CreateForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CreateForumTopic) WithApiBaseUrl(u string) *CreateForumTopic {
	s.baseUrl = u
	return s
}

func (s CreateForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (c CreateForumTopic) Validate() error {
	if strings.TrimSpace(c.ChatId) == "" {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if len(c.Name) < 1 || len(c.Name) > 128 {
		return objects.ErrInvalidParam("name parameter must be between 1 and 128 characters long")
	}
	if c.IconColor != nil {
		var validIconColors = map[int]struct{}{
			7322096:  {},
			16766590: {},
			13338331: {},
			9367192:  {},
			16749490: {},
			16478047: {},
		}
		if _, ok := validIconColors[*c.IconColor]; !ok {
			return objects.ErrInvalidParam("icon_color must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)")
		}
	}
	return nil
}

func (c CreateForumTopic) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateForumTopic) Execute(token string) (*objects.ForumTopic, error) {
	return SendTelegramPostRequest[objects.ForumTopic](token, "createForumTopic", c)
}

// Use this method to edit name and icon of a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type EditForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
	//New topic name, 0-128 characters. If not specified or empty, the current name of the topic will be kept
	Name *string `json:"name,omitempty"`
	//New unique identifier of the custom emoji shown as the topic icon.
	//Use getForumTopicIconStickers to get all allowed custom emoji identifiers.
	//Pass an empty string to remove the icon. If not specified, the current icon will be kept
	IconCustomEmojiId *string `json:"icon_custom_emoji_id,omitempty"`
	client            *http.Client
	baseUrl           string
}

func (s *EditForumTopic) WithClient(c *http.Client) *EditForumTopic {
	s.client = c
	return s
}

func (s EditForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditForumTopic) WithApiBaseUrl(u string) *EditForumTopic {
	s.baseUrl = u
	return s
}

func (s EditForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "editForumTopic", e)
}

// Use this method to close an open topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type CloseForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
	client          *http.Client
	baseUrl         string
}

func (s *CloseForumTopic) WithClient(c *http.Client) *CloseForumTopic {
	s.client = c
	return s
}

func (s CloseForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CloseForumTopic) WithApiBaseUrl(u string) *CloseForumTopic {
	s.baseUrl = u
	return s
}

func (s CloseForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e CloseForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "closeForumTopic", e)
}

// Use this method to reopen a closed topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights, unless it is the creator of the topic.
// Returns True on success.
type ReopenForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
	client          *http.Client
	baseUrl         string
}

func (s *ReopenForumTopic) WithClient(c *http.Client) *ReopenForumTopic {
	s.client = c
	return s
}

func (s ReopenForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ReopenForumTopic) WithApiBaseUrl(u string) *ReopenForumTopic {
	s.baseUrl = u
	return s
}

func (s ReopenForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e ReopenForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "reopenForumTopic", e)
}

// Use this method to delete a forum topic along with all its messages in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_delete_messages administrator rights.
// Returns True on success.
type DeleteForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
	client          *http.Client
	baseUrl         string
}

func (s *DeleteForumTopic) WithClient(c *http.Client) *DeleteForumTopic {
	s.client = c
	return s
}

func (s DeleteForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteForumTopic) WithApiBaseUrl(u string) *DeleteForumTopic {
	s.baseUrl = u
	return s
}

func (s DeleteForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e DeleteForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteForumTopic", e)
}

// Use this method to clear the list of pinned messages in a forum topic.
// The bot must be an administrator in the chat for this to work and
// must have the can_pin_messages administrator right in the supergroup.
// Returns True on success.
type UnpinAllForumTopicMessages struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//Unique identifier for the target message thread of the forum topic
	MessageThreadId string `json:"message_thread_id"`
	client          *http.Client
	baseUrl         string
}

func (s *UnpinAllForumTopicMessages) WithClient(c *http.Client) *UnpinAllForumTopicMessages {
	s.client = c
	return s
}

func (s UnpinAllForumTopicMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnpinAllForumTopicMessages) WithApiBaseUrl(u string) *UnpinAllForumTopicMessages {
	s.baseUrl = u
	return s
}

func (s UnpinAllForumTopicMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e UnpinAllForumTopicMessages) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unpinAllForumTopicMessages", e)
}

// Use this method to edit the name of the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns True on success.
type EditGeneralForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId string `json:"chat_id"`
	//New topic name, 1-128 characters
	Name    string `json:"name"`
	client  *http.Client
	baseUrl string
}

func (s *EditGeneralForumTopic) WithClient(c *http.Client) *EditGeneralForumTopic {
	s.client = c
	return s
}

func (s EditGeneralForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *EditGeneralForumTopic) WithApiBaseUrl(u string) *EditGeneralForumTopic {
	s.baseUrl = u
	return s
}

func (s EditGeneralForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e EditGeneralForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "editGeneralForumTopic", e)
}

// Use this method to close an open 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and
// must have the can_manage_topics administrator rights.
// Returns True on success.
type CloseGeneralForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *CloseGeneralForumTopic) WithClient(c *http.Client) *CloseGeneralForumTopic {
	s.client = c
	return s
}

func (s CloseGeneralForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *CloseGeneralForumTopic) WithApiBaseUrl(u string) *CloseGeneralForumTopic {
	s.baseUrl = u
	return s
}

func (s CloseGeneralForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e CloseGeneralForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "closeGeneralForumTopic", e)
}

// Use this method to reopen a closed 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// The topic will be automatically unhidden if it was hidden.
// Returns True on success.
type ReopenGeneralForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *ReopenGeneralForumTopic) WithClient(c *http.Client) *ReopenGeneralForumTopic {
	s.client = c
	return s
}

func (s ReopenGeneralForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *ReopenGeneralForumTopic) WithApiBaseUrl(u string) *ReopenGeneralForumTopic {
	s.baseUrl = u
	return s
}

func (s ReopenGeneralForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e ReopenGeneralForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "reopenGeneralForumTopic", e)
}

// Use this method to hide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// The topic will be automatically closed if it was open. Returns True on success.
type HideGeneralForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *HideGeneralForumTopic) WithClient(c *http.Client) *HideGeneralForumTopic {
	s.client = c
	return s
}

func (s HideGeneralForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *HideGeneralForumTopic) WithApiBaseUrl(u string) *HideGeneralForumTopic {
	s.baseUrl = u
	return s
}

func (s HideGeneralForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e HideGeneralForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "hideGeneralForumTopic", e)
}

// Use this method to unhide the 'General' topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the can_manage_topics administrator rights.
// Returns True on success.
type UnhideGeneralForumTopic struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *UnhideGeneralForumTopic) WithClient(c *http.Client) *UnhideGeneralForumTopic {
	s.client = c
	return s
}

func (s UnhideGeneralForumTopic) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnhideGeneralForumTopic) WithApiBaseUrl(u string) *UnhideGeneralForumTopic {
	s.baseUrl = u
	return s
}

func (s UnhideGeneralForumTopic) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e UnhideGeneralForumTopic) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unhideGeneralForumTopic", e)
}

// Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and
// must have the can_pin_messages administrator right in the supergroup.
// Returns True on success.
type UnpinAllGeneralForumTopicMessages struct {
	//Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatId  string `json:"chat_id"`
	client  *http.Client
	baseUrl string
}

func (s *UnpinAllGeneralForumTopicMessages) WithClient(c *http.Client) *UnpinAllGeneralForumTopicMessages {
	s.client = c
	return s
}

func (s UnpinAllGeneralForumTopicMessages) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *UnpinAllGeneralForumTopicMessages) WithApiBaseUrl(u string) *UnpinAllGeneralForumTopicMessages {
	s.baseUrl = u
	return s
}

func (s UnpinAllGeneralForumTopicMessages) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (e UnpinAllGeneralForumTopicMessages) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "unpinAllGeneralForumTopicMessages", e)
}

// Use this method to send answers to callback queries sent from inline keyboards.
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.
// On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL.
// For this option to work, you must first create a game for your bot via @BotFather and accept the terms.
// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
type AnswerCallbackQuery struct {
	//Unique identifier for the query to be answered
	CallbackQueryId string `json:"callback_query_id"`
	//Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text *string `json:"text"`
	//If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert *bool `json:"show_alert"`
	//URL that will be opened by the user's client.
	//If you have created a Game and accepted the conditions via @BotFather,
	//specify the URL that opens your game - note that this will only work if the query comes from a callback_game button.
	//
	//Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	Url *string `json:"url"`
	//The maximum amount of time in seconds that the result of the callback query may be cached client-side.
	//Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime *int `json:"cache_time"`
	client    *http.Client
	baseUrl   string
}

func (s *AnswerCallbackQuery) WithClient(c *http.Client) *AnswerCallbackQuery {
	s.client = c
	return s
}

func (s AnswerCallbackQuery) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *AnswerCallbackQuery) WithApiBaseUrl(u string) *AnswerCallbackQuery {
	s.baseUrl = u
	return s
}

func (s AnswerCallbackQuery) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (a AnswerCallbackQuery) Validate() error {
	if strings.TrimSpace(a.CallbackQueryId) == "" {
		return objects.ErrInvalidParam("callback_query_id parameter can't be empty")
	}
	if a.Text != nil {
		if len(*a.Text) > 200 {
			return objects.ErrInvalidParam("text parameter must not be longer than 200 characters if specified")
		}
	}
	return nil
}

func (a AnswerCallbackQuery) ToRequestBody() ([]byte, error) {
	return json.Marshal(a)
}

func (a AnswerCallbackQuery) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "answerCallbackQuery", a)
}

// Use this method to get the list of boosts added to a chat by a user. Requires administrator rights in the chat. Returns a UserChatBoosts object.
type GetUserChatBoosts struct {
	//Unique identifier for the chat or username of the channel (in the format @channelusername)
	ChatId string `json:"chat_id"`
	//Unique identifier of the target user
	UserId  int `json:"user_id"`
	client  *http.Client
	baseUrl string
}

func (s *GetUserChatBoosts) WithClient(c *http.Client) *GetUserChatBoosts {
	s.client = c
	return s
}

func (s GetUserChatBoosts) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetUserChatBoosts) WithApiBaseUrl(u string) *GetUserChatBoosts {
	s.baseUrl = u
	return s
}

func (s GetUserChatBoosts) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetUserChatBoosts) Execute(token string) (*objects.UserChatBoosts, error) {
	return SendTelegramGetRequest[objects.UserChatBoosts](token, "getUserChatBoosts", g)
}

// Use this method to get information about the connection of the bot with a business account.
// Returns a BusinessConnection object on success.
type GetBusinessConnection struct {
	//Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	client               *http.Client
	baseUrl              string
}

func (s *GetBusinessConnection) WithClient(c *http.Client) *GetBusinessConnection {
	s.client = c
	return s
}

func (s GetBusinessConnection) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetBusinessConnection) WithApiBaseUrl(u string) *GetBusinessConnection {
	s.baseUrl = u
	return s
}

func (s GetBusinessConnection) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (g GetBusinessConnection) Execute(token string) (*objects.BusinessConnection, error) {
	return SendTelegramGetRequest[objects.BusinessConnection](token, "getBusinessConnection", g)
}

// Use this method to change the list of the bot's commands. See this manual for more details about bot commands. Returns True on success.
type SetMyCommands struct {
	//A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []objects.BotCommand `json:"commands"`
	//A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	//A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *SetMyCommands) WithClient(c *http.Client) *SetMyCommands {
	s.client = c
	return s
}

func (s SetMyCommands) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMyCommands) WithApiBaseUrl(u string) *SetMyCommands {
	s.baseUrl = u
	return s
}

func (s SetMyCommands) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetMyCommands) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMyCommands", s)
}

// Use this method to delete the list of the bot's commands for the given scope and user language.
// After deletion, higher level commands will be shown to affected users. Returns True on success.
type DeleteMyCommands struct {
	//A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	//A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *DeleteMyCommands) WithClient(c *http.Client) *DeleteMyCommands {
	s.client = c
	return s
}

func (s DeleteMyCommands) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *DeleteMyCommands) WithApiBaseUrl(u string) *DeleteMyCommands {
	s.baseUrl = u
	return s
}

func (s DeleteMyCommands) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s DeleteMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s DeleteMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeleteMyCommands) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "deleteMyCommands", s)
}

// Use this method to get the current list of the bot's commands for the given scope and user language.
// Returns an Array of BotCommand objects. If commands aren't set, an empty list is returned.
type GetMyCommands struct {
	//A JSON-serialized object, describing scope of users. Defaults to BotCommandScopeDefault.
	Scope objects.BotCommandScope `json:"scope,omitempty"`
	//A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *GetMyCommands) WithClient(c *http.Client) *GetMyCommands {
	s.client = c
	return s
}

func (s GetMyCommands) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetMyCommands) WithApiBaseUrl(u string) *GetMyCommands {
	s.baseUrl = u
	return s
}

func (s GetMyCommands) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s GetMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return objects.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s GetMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyCommands) Execute(token string) (*[]objects.BotCommand, error) {
	return SendTelegramGetRequest[[]objects.BotCommand](token, "getMyCommands", s)
}

// Use this method to change the bot's name. Returns True on success.
type SetMyName struct {
	//New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language.
	Name *string `json:"name,omitempty"`
	//A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name.
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *SetMyName) WithClient(c *http.Client) *SetMyName {
	s.client = c
	return s
}

func (s SetMyName) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMyName) WithApiBaseUrl(u string) *SetMyName {
	s.baseUrl = u
	return s
}

func (s SetMyName) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetMyName) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMyName", s)
}

// Use this method to get the current bot name for the given user language. Returns BotName on success.
type GetMyName struct {
	//A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *GetMyName) WithClient(c *http.Client) *GetMyName {
	s.client = c
	return s
}

func (s GetMyName) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetMyName) WithApiBaseUrl(u string) *GetMyName {
	s.baseUrl = u
	return s
}

func (s GetMyName) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s GetMyName) Execute(token string) (*objects.BotName, error) {
	return SendTelegramGetRequest[objects.BotName](token, "getMyName", s)
}

// Use this method to change the bot's description, which is shown in the chat with the bot if the chat is empty. Returns True on success.
type SetMyDescription struct {
	//New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description *string `json:"description,omitempty"`
	//A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *SetMyDescription) WithClient(c *http.Client) *SetMyDescription {
	s.client = c
	return s
}

func (s SetMyDescription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMyDescription) WithApiBaseUrl(u string) *SetMyDescription {
	s.baseUrl = u
	return s
}

func (s SetMyDescription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetMyDescription) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMyDescription", s)
}

// Use this method to get the current bot description for the given user language. Returns BotDescription on success.
type GetMyDescription struct {
	//A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *GetMyDescription) WithClient(c *http.Client) *GetMyDescription {
	s.client = c
	return s
}

func (s GetMyDescription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetMyDescription) WithApiBaseUrl(u string) *GetMyDescription {
	s.baseUrl = u
	return s
}

func (s GetMyDescription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s GetMyDescription) Execute(token string) (*objects.BotDescription, error) {
	return SendTelegramGetRequest[objects.BotDescription](token, "getMyDescription", s)
}

// Use this method to change the bot's short description, which is shown on the bot's profile page and is sent together with the link when users share the bot.
// Returns True on success.
type SetMyShortDescription struct {
	//New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription *string `json:"short_description,omitempty"`
	//A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *SetMyShortDescription) WithClient(c *http.Client) *SetMyShortDescription {
	s.client = c
	return s
}

func (s SetMyShortDescription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMyShortDescription) WithApiBaseUrl(u string) *SetMyShortDescription {
	s.baseUrl = u
	return s
}

func (s SetMyShortDescription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetMyShortDescription) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMyShortDescription", s)
}

// Use this method to get the current bot short description for the given user language. Returns BotShortDescription on success.
type GetMyShortDescription struct {
	//A two-letter ISO 639-1 language code or an empty string
	LanguageCode *string `json:"language_code,omitempty"`
	client       *http.Client
	baseUrl      string
}

func (s *GetMyShortDescription) WithClient(c *http.Client) *GetMyShortDescription {
	s.client = c
	return s
}

func (s GetMyShortDescription) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetMyShortDescription) WithApiBaseUrl(u string) *GetMyShortDescription {
	s.baseUrl = u
	return s
}

func (s GetMyShortDescription) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s GetMyShortDescription) Execute(token string) (*objects.BotShortDescription, error) {
	return SendTelegramGetRequest[objects.BotShortDescription](token, "getMyShortDescription", s)
}

// Use this method to change the bot's menu button in a private chat, or the default menu button. Returns True on success.
type SetChatMenuButton struct {
	//Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	ChatId *string `json:"chat_id,omitempty"`
	//A JSON-serialized object for the bot's new menu button. Defaults to MenuButtonDefault
	MenuButton objects.MenuButton `json:"menu_button,omitempty"`
	client     *http.Client
	baseUrl    string
}

func (s *SetChatMenuButton) WithClient(c *http.Client) *SetChatMenuButton {
	s.client = c
	return s
}

func (s SetChatMenuButton) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetChatMenuButton) WithApiBaseUrl(u string) *SetChatMenuButton {
	s.baseUrl = u
	return s
}

func (s SetChatMenuButton) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s SetChatMenuButton) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setChatMenuButton", s)
}

// Use this method to get the current value of the bot's menu button in a private chat, or the default menu button.
// Returns MenuButton on success.
type GetChatMenuButton struct {
	//Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
	ChatId  *int `json:"chat_id,omitempty"`
	client  *http.Client
	baseUrl string
}

func (s *GetChatMenuButton) WithClient(c *http.Client) *GetChatMenuButton {
	s.client = c
	return s
}

func (s GetChatMenuButton) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetChatMenuButton) WithApiBaseUrl(u string) *GetChatMenuButton {
	s.baseUrl = u
	return s
}

func (s GetChatMenuButton) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
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

func (s GetChatMenuButton) Execute(token string) (*objects.MenuButton, error) {
	return SendTelegramGetRequest[objects.MenuButton](token, "setChatMenuButton", s)
}

// Use this method to change the default administrator rights requested by the bot when it's added as an administrator to groups or channels.
// These rights will be suggested to users, but they are free to modify the list before adding the bot.
// Returns True on success.
type SetMyDefaultAdministratorRights struct {
	//A JSON-serialized object describing new default administrator rights. If not specified, the default administrator rights will be cleared.
	Rights *objects.ChatAdministratorRights `json:"rights,omitempty"`
	//Pass True to change the default administrator rights of the bot in channels.
	//Otherwise, the default administrator rights of the bot for groups and supergroups will be changed.
	ForChannels *bool `json:"for_channels,omitempty"`
	client      *http.Client
	baseUrl     string
}

func (s *SetMyDefaultAdministratorRights) WithClient(c *http.Client) *SetMyDefaultAdministratorRights {
	s.client = c
	return s
}

func (s SetMyDefaultAdministratorRights) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetMyDefaultAdministratorRights) WithApiBaseUrl(u string) *SetMyDefaultAdministratorRights {
	s.baseUrl = u
	return s
}

func (s SetMyDefaultAdministratorRights) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

// always nil
func (s SetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s SetMyDefaultAdministratorRights) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyDefaultAdministratorRights) Execute(token string) (*bool, error) {
	return SendTelegramPostRequest[bool](token, "setMyDefaultAdministratorRights", s)
}

// Use this method to get the current default administrator rights of the bot. Returns ChatAdministratorRights on success.
type GetMyDefaultAdministratorRights struct {
	//Pass True to get default administrator rights of the bot in channels.
	//Otherwise, default administrator rights of the bot for groups and supergroups will be returned.
	ForChannels *bool `json:"for_channels,omitempty"`
	client      *http.Client
	baseUrl     string
}

func (s *GetMyDefaultAdministratorRights) WithClient(c *http.Client) *GetMyDefaultAdministratorRights {
	s.client = c
	return s
}

func (s GetMyDefaultAdministratorRights) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetMyDefaultAdministratorRights) WithApiBaseUrl(u string) *GetMyDefaultAdministratorRights {
	s.baseUrl = u
	return s
}

func (s GetMyDefaultAdministratorRights) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

// always nil
func (s GetMyDefaultAdministratorRights) Validate() error {
	return nil
}

func (s GetMyDefaultAdministratorRights) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s GetMyDefaultAdministratorRights) Execute(token string) (*objects.ChatAdministratorRights, error) {
	return SendTelegramPostRequest[objects.ChatAdministratorRights](token, "getMyDefaultAdministratorRights", s)
}
