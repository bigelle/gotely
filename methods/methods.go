package methods

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/bigelle/tele.go/objects"
	iso6391 "github.com/emvi/iso-639-1"
)

// Use this method to send text messages. On success, the sent Message is returned.
type SendMessage[T int | string] struct {
	//Required
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId T `json:"chat_id"`
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

func (s SendMessage[T]) Validate() error {
	if strings.TrimSpace(s.Text) == "" {
		return objects.ErrInvalidParam("text parameter can't be empty")
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c == 0 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (s SendMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendMessage[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendMessage", s)
}

// Use this method to forward messages of any kind.
// Service messages and messages with protected content can't be forwarded.
// On success, the sent Message is returned.
type ForwardMessage[T int | string] struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId T
	//Required.
	//Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername
	FromChatId T
	//Required.
	//Message identifier in the chat specified in from_chat_id
	MessageId int
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int
	//Optional.
	//Sends the message silently. Users will receive a notification with no sound.
	DisableNotification *bool
	//Optional.
	//Protects the contents of the forwarded message from forwarding and saving
	ProtectContent *bool
}

func (f ForwardMessage[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if f.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (f ForwardMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(f)
}

func (f ForwardMessage[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("forwardMessage", f)
}

// Use this method to forward multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
// On success, an array of MessageId of the sent messages is returned.
type ForwardMessages[T int | string] struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId T
	//Required.
	//Unique identifier for the chat where the original messages were sent (or channel username in the format @channelusername)
	FromChatId T
	//Required.
	//A JSON-serialized list of 1-100 identifiers of messages in the chat from_chat_id to forward.
	//The identifiers must be specified in a strictly increasing order.
	MessageIds []int
	//Optional.
	//Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
	MessageThreadId *int
	//Optional.
	//Sends the messages silently. Users will receive a notification with no sound.
	DisableNotification *bool
	//Optional.
	//Protects the contents of the forwarded messages from forwarding and saving
	ProtectContent *bool
}

func (f ForwardMessages[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(f.MessageIds) < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (f ForwardMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(f)
}

func (f ForwardMessages[T]) Execute() (*[]objects.MessageId, error) {
	return MakePostRequest[[]objects.MessageId]("forwardMessages", f)
}

// Use this method to copy messages of any kind.
// Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
// A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.
// The method is analogous to the method forwardMessage, but the copied message doesn't have a link to the original message.
// Returns the MessageId of the sent message on success.
type CopyMessage[T int | string] struct {
	//Required.
	//Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatId T `json:"chat_id"`
	//Required.
	//Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	FromChatId T `json:"from_chat_id"`
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

func (c CopyMessage[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c.MessageId < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessage[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(c)
}

func (c CopyMessage[T]) Execute() (*objects.MessageId, error) {
	return MakePostRequest[objects.MessageId]("copyMessage", c)
}

type CopyMessages[T int | string] struct {
	ChatId              T
	FromChatId          T
	MessageIds          []int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
	RemoveCaption       *bool
}

func (c CopyMessages[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(c)
}

func (c CopyMessages[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return objects.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(c.MessageIds) < 1 {
		return objects.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessages[T]) Execute() (*[]objects.MessageId, error) {
	return MakePostRequest[[]objects.MessageId]("copyMessages", c)
}

type SendPhoto[T int | string, B objects.InputFile | string] struct {
	ChatId                T
	Photo                 B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]objects.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *objects.ReplyParameters
	ReplyMarkup           *objects.ReplyMarkup
}

func (s SendPhoto[T, B]) Validate() error {
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
	if p, ok := any(s.Photo).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Photo).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendPhoto[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPhoto[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPhoto", s)
}

type SendAudio[T int | string, B objects.InputFile | string] struct {
	ChatId               T
	Audio                B
	BusinessConnectionId *string
	MessageThreadId      *int
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]objects.MessageEntity
	Duration             *int
	Performer            *string
	Title                *string
	Thumbnail            *B
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendAudio[T, B]) Validate() error {
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
	if p, ok := any(s.Audio).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid audio parameter: %w", err)
		}
	}
	if p, ok := any(s.Audio).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("audio parameter can't be empty")
		}
	}
	return nil
}

func (s SendAudio[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAudio[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendAudio", s)
}

type SendDocument[T int | string, B objects.InputFile | string] struct {
	ChatId                      T
	Document                    B
	BusinessConnectionId        *string
	MessageThreadId             *int
	Caption                     *string
	ParseMode                   *string
	CaptionEntities             *[]objects.MessageEntity
	DisableContentTypeDetection *bool
	Thumbnail                   *B
	DisableNotification         *bool
	ProtectContent              *bool
	AllowPaidBroadcast          *bool
	MessageEffectId             *string
	ReplyParameters             *objects.ReplyParameters
	ReplyMarkup                 *objects.ReplyMarkup
}

func (s SendDocument[T, B]) Validate() error {
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
	if p, ok := any(s.Document).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid document parameter: %w", err)
		}
	}
	if p, ok := any(s.Document).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("document parameter can't be empty")
		}
	}
	return nil
}

func (s SendDocument[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDocument[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendDocument", s)
}

type SendVideo[T int | string, B objects.InputFile | string] struct {
	ChatId                T
	Video                 B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Duration              *int
	Width                 *int
	Height                *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]objects.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	SupportsStreaming     *bool
	Thumbnail             *B
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *objects.ReplyParameters
	ReplyMarkup           *objects.ReplyMarkup
}

func (s SendVideo[T, B]) Validate() error {
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
	if p, ok := any(s.Video).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video parameter: %w", err)
		}
	}
	if p, ok := any(s.Video).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("video parameter can't be empty")
		}
	}
	return nil
}

func (s SendVideo[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideo[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendVideo", s)
}

type SendAnimation[T int | string, B objects.InputFile | string] struct {
	ChatId                T
	Animation             B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Duration              *int
	Width                 *int
	Height                *int
	Thumbnail             *B
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]objects.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *objects.ReplyParameters
	ReplyMarkup           *objects.ReplyMarkup
}

func (s SendAnimation[T, B]) Validate() error {
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
	if p, ok := any(s.Animation).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Animation).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendAnimation[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAnimation[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendAnimation", s)
}

type SendVoice[T int | string, B objects.InputFile | string] struct {
	ChatId               T
	Voice                B
	BusinessConnectionId *string
	MessageThreadId      *int
	Duration             *int
	Thumbnail            *B
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]objects.MessageEntity
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendVoice[T, B]) Validate() error {
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
	if p, ok := any(s.Voice).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid voice parameter: %w", err)
		}
	}
	if p, ok := any(s.Voice).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("voice parameter can't be empty")
		}
	}
	return nil
}

func (s SendVoice[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVoice[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendVoice", s)
}

type SendVideoNote[T int | string, B objects.InputFile | string] struct {
	ChatId               T
	VideoNote            B
	BusinessConnectionId *string
	MessageThreadId      *int
	Duration             *int
	Length               *int
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]objects.MessageEntity
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendVideoNote[T, B]) Validate() error {
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
	if p, ok := any(s.VideoNote).(objects.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video_note parameter: %w", err)
		}
	}
	if p, ok := any(s.VideoNote).(string); ok {
		if strings.TrimSpace(p) == "" {
			return objects.ErrInvalidParam("video_note parameter can't be empty")
		}
	}
	// TODO: validate non-nill thumbnails
	return nil
}

func (s SendVideoNote[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideoNote[T, B]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendVideoNote", s)
}

type SendPaidMedia[T int | string] struct {
	ChatId                T
	Media                 []objects.InputPaidMedia
	StarCount             int
	BusinessConnectionId  *string
	Payload               *string
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]objects.MessageEntity
	ShowCaptionAboveMedia *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	ReplyParameters       *objects.ReplyParameters
	ReplyMarkup           *objects.ReplyMarkup
}

func (s SendPaidMedia[T]) Validate() error {
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

func (s SendPaidMedia[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPaidMedia[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPaidMedia", s)
}

type SendMediaGroup[T int | string] struct {
	ChatId               T
	Media                []objects.InputMedia
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
}

func (s SendMediaGroup[T]) Validate() error {
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

func (s SendMediaGroup[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendMediaGroup[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendMediaGroup", s)
}

type SendLocation[T int | string] struct {
	ChatId               T
	Latitude             *float64
	Longtitude           *float64
	HorizontalAccuracy   *float64
	BusinessConnectionId *string
	MessageThreadId      *string
	LivePeriod           *int
	Heading              *int
	ProximityAlertRadius *int
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendLocation[T]) Validate() error {
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
	if s.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return objects.ErrInvalidParam("longtitude parameter can't be empty")
	}
	// TODO: validate replyparameters everywhere
	return nil
}

func (s SendLocation[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendLocation[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendLocation", s)
}

type SendVenue[T int | string] struct {
	ChatId               T
	Latitude             *float64
	Longtitude           *float64
	Title                string
	Address              string
	FoursquareId         *string
	FoursquareType       *string
	GooglePlaceId        *string
	GooglePlaceType      *string
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendVenue[T]) Validate() error {
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
	if s.Latitude == nil {
		return objects.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return objects.ErrInvalidParam("longtitude parameter can't be empty")
	}
	return nil
}

func (s SendVenue[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVenue[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendVenue", s)
}

type SendContact[T int | string] struct {
	ChatId               T
	PhoneNumber          string
	FirstName            string
	LastName             *string
	Vcard                *string
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.ReplyMarkup
}

func (s SendContact[T]) Validate() error {
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
	if strings.TrimSpace(s.PhoneNumber) == "" {
		return objects.ErrInvalidParam("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return objects.ErrInvalidParam("first_name parameter can't be empty")
	}
	return nil
}

// NOTE: do i need it?
func (s SendContact[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendContact[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendContact", s)
}

type SendPoll[T int | string] struct {
	ChatId               T
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

func (s SendPoll[T]) Validate() error {
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
	if strings.TrimSpace(s.Question) == "" {
		return objects.ErrInvalidParam("question parameter can't be empty")
	}
	if len(s.Options) < 2 || len(s.Options) > 10 {
		return objects.ErrInvalidParam("options parameter must be between 2 and 10")
	}
	return nil
}

func (s SendPoll[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPoll[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPoll", s)
}

type SendDice[T int | string] struct {
	ChatId               T
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

func (s SendDice[T]) Validate() error {
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
	if strings.TrimSpace(s.Emoji) == "" {
		return objects.ErrInvalidParam("emoji parameter can't be empty")
	}
	return nil
}

func (s SendDice[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDice[T]) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendPoll", s)
}

type SendChatAction[T int | string] struct {
	ChatId               T
	Action               string
	BusinessConnectionId *string
	MessageThreadId      *string
}

func (s SendChatAction[T]) Validate() error {
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

type SetMessageReaction[T int | string] struct {
	ChatId    T
	MessageId int
	Reaction  *[]objects.ReactionType
	IsBig     *bool
}

func (s SetMessageReaction[T]) Validate() error {
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
	return nil
}

func (s SetMessageReaction[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMessageReaction[T]) Execute() (*bool, error) {
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

type BanChatMember[T int | string] struct {
	ChatId         T
	UserId         int
	UntilDate      *int
	RevokeMessages *bool
}

func (b BanChatMember[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b BanChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatMember[T]) Execute() (*bool, error) {
	return MakeGetRequest[bool]("banChatMember", b)
}

type UnbanChatMember[T int | string] struct {
	ChatId       T
	UserId       int
	OnlyIfBanned *bool
}

func (b UnbanChatMember[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatMember[T]) Execute() (*bool, error) {
	return MakeGetRequest[bool]("unbanChatMember", b)
}

type RestrictChatMember[T int | string] struct {
	ChatId                         T
	UserId                         int
	Permissions                    objects.ChatPermissions
	UserIndependentChatPermissions *bool
	UntilDate                      *int
}

func (r RestrictChatMember[T]) Validate() error {
	if c, ok := any(r.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(r.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if r.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (r RestrictChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r RestrictChatMember[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("restrictChatMember", r)
}

type PromoteChatMember[T int | string] struct {
	ChatId              T
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

func (p PromoteChatMember[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p PromoteChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PromoteChatMember[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("promoteChatMember", p)
}

type SetChatAdministratorCustomTitle[T int | string] struct {
	ChatId      T
	UserId      int
	CustomTitle string
}

func (s SetChatAdministratorCustomTitle[T]) Validate() error {
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

func (s SetChatAdministratorCustomTitle[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatAdministratorCustomTitle[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatAdministratorCustomTitle", s)
}

type BanChatSenderChat[T int | string] struct {
	ChatId       T
	SenderChatId int
}

func (b BanChatSenderChat[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.SenderChatId < 1 {
		return objects.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b BanChatSenderChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatSenderChat[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("banChatSenderChat", b)
}

type UnbanChatSenderChat[T int | string] struct {
	ChatId       T
	SenderChatId int
}

func (b UnbanChatSenderChat[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.SenderChatId < 1 {
		return objects.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatSenderChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatSenderChat[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("unbanChatSenderChat", b)
}

type SetChatPermissions[T int | string] struct {
	ChatId                         T
	Permissions                    objects.ChatPermissions
	UserIndependentChatPermissions *bool
}

func (s SetChatPermissions[T]) Validate() error {
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
	return nil
}

func (s SetChatPermissions[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPermissions[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatPermissions", s)
}

type ExportChatInviteLink[T int | string] struct {
	ChatId T
}

func (e ExportChatInviteLink[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e ExportChatInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ExportChatInviteLink[T]) Execute() (*string, error) {
	return MakePostRequest[string]("exportChatInviteLink", e)
}

type CreateInviteLink[T int | string] struct {
	ChatId             T
	Name               *string
	ExpireDate         *int
	MemberLimit        *int
	CreatesJoinRequest *bool
}

func (c CreateInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (c CreateInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateInviteLink[T]) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("createInviteLink", c)
}

type EditInviteLink[T int | string] struct {
	ChatId             T
	InviteLink         string
	Name               *string
	ExpireDate         *int
	MemberLimit        *int
	CreatesJoinRequest *bool
}

func (c EditInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (c EditInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditInviteLink[T]) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("editInviteLink", c)
}

type CreateChatSubscriptionInviteLink[T int | string] struct {
	ChatId             T
	SubscriptionPeriod int
	SubscriptionPrice  int
	Name               *string
}

func (c CreateChatSubscriptionInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (c CreateChatSubscriptionInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateChatSubscriptionInviteLink[T]) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("createChatSubscriptionInviteLink", c)
}

type EditChatSubscriptionInviteLink[T int | string] struct {
	ChatId     T
	InviteLink string
	Name       *string
}

func (c EditChatSubscriptionInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (c EditChatSubscriptionInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditChatSubscriptionInviteLink[T]) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("editChatSubscriptionInviteLink", c)
}

type RevokeInviteLink[T int | string] struct {
	ChatId     T
	InviteLink string
	Name       *string
}

func (c RevokeInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return objects.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c RevokeInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c RevokeInviteLink[T]) Execute() (*objects.ChatInviteLink, error) {
	return MakePostRequest[objects.ChatInviteLink]("revokeInviteLink", c)
}

type ApproveChatJoinRequest[T int | string] struct {
	ChatId T
	UserId int
}

func (s ApproveChatJoinRequest[T]) Validate() error {
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
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s ApproveChatJoinRequest[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s ApproveChatJoinRequest[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("approveChatJoinRequest", s)
}

type DeclineChatJoinRequest[T int | string] struct {
	ChatId T
	UserId int
}

func (s DeclineChatJoinRequest[T]) Validate() error {
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
	if s.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s DeclineChatJoinRequest[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeclineChatJoinRequest[T]) Execute() (*bool, error) {
	// NOTE: maybe there's a better way to get token?
	return MakePostRequest[bool]("declineChatJoinRequest", s)
}

type SetChatPhoto[T int | string] struct {
	ChatId T
	Photo  objects.InputFile
}

func (s SetChatPhoto[T]) Validate() error {
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
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SetChatPhoto[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPhoto[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatPhoto", s)
}

type DeleteChatPhoto[T int | string] struct {
	ChatId T
}

func (d DeleteChatPhoto[T]) Validate() error {
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
	return nil
}

func (d DeleteChatPhoto[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteChatPhoto[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteChatPhoto", d)
}

type SetChatTitle[T int | string] struct {
	ChatId T
	Title  string
}

func (s SetChatTitle[T]) Validate() error {
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
	if len(s.Title) < 1 || len(s.Title) > 128 {
		return objects.ErrInvalidParam("title parameter must be between 1 and 128 characters long")
	}
	return nil
}

func (s SetChatTitle[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatTitle[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatTitle", s)
}

type SetChatDescription[T int | string] struct {
	ChatId      T
	Description string
}

func (s SetChatDescription[T]) Validate() error {
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
	if len(s.Description) > 255 {
		return objects.ErrInvalidParam("description parameter must not be longer than 255 characters")
	}
	return nil
}

func (s SetChatDescription[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatDescription[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatTitle", s)
}

type PinChatMessage[T int | string] struct {
	ChatId               T
	MessageId            int
	BusinessConnectionId *string
	DisableNotification  *bool
}

func (p PinChatMessage[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p PinChatMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PinChatMessage[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("pinChatMessage", p)
}

type UnpinChatMessage[T int | string] struct {
	ChatId               T
	MessageId            int
	BusinessConnectionId *string
}

func (p UnpinChatMessage[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.MessageId < 1 {
		return objects.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p UnpinChatMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinChatMessage[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinChatMessage", p)
}

type UnpinAllChatMessages[T int | string] struct {
	ChatId T
}

func (p UnpinAllChatMessages[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p UnpinAllChatMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinAllChatMessages[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinAllChatMessages", p)
}

type LeaveChat[T int | string] struct {
	ChatId T
}

func (p LeaveChat[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p LeaveChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p LeaveChat[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("leaveChat", p)
}

type GetChat[T int | string] struct {
	ChatId T
}

func (p GetChat[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChat[T]) Execute() (*objects.ChatFullInfo, error) {
	return MakeGetRequest[objects.ChatFullInfo]("getChat", p)
}

type GetChatAdministrators[T int | string] struct {
	ChatId T
}

func (p GetChatAdministrators[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChatAdministrators[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatAdministrators[T]) Execute() (*[]objects.ChatMember, error) {
	return MakeGetRequest[[]objects.ChatMember]("getChatAdministrators", p)
}

type GetChatMemberCount[T int | string] struct {
	ChatId T
}

func (p GetChatMemberCount[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p GetChatMemberCount[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatMemberCount[T]) Execute() (*int, error) {
	return MakeGetRequest[int]("getChatMemberCount", p)
}

type GetChatMember[T int | string] struct {
	ChatId T
	UserId int
}

func (p GetChatMember[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p GetChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p GetChatMember[T]) Execute() (*objects.ChatMember, error) {
	return MakeGetRequest[objects.ChatMember]("getChatMember", p)
}

type SetChatStickerSet[T int | string] struct {
	ChatId         T
	StickerSetName string
}

func (p SetChatStickerSet[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(p.StickerSetName) == "" {
		return objects.ErrInvalidParam("sticker_set_name parameter can't be empty")
	}
	return nil
}

func (p SetChatStickerSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p SetChatStickerSet[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatStickerSet", p)
}

type DeleteChatStickerSet[T int | string] struct {
	ChatId T
}

func (p DeleteChatStickerSet[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p DeleteChatStickerSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p DeleteChatStickerSet[T]) Execute() (*bool, error) {
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

type CreateForumTopic[T int | string] struct {
	ChatId            T
	Name              string
	IconColor         *int
	IconCustomEmojiId *string
}

func (c CreateForumTopic[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (c CreateForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateForumTopic[T]) Execute() (*objects.ForumTopic, error) {
	return MakePostRequest[objects.ForumTopic]("createForumTopic", c)
}

type EditForumTopic[T int | string] struct {
	ChatId            T
	MessageThreadId   string
	Name              *string
	IconCustomEmojiId *string
}

func (e EditForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
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

func (e EditForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("editForumTopic", e)
}

type CloseForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e CloseForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e CloseForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("closeForumTopic", e)
}

type ReopenForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e ReopenForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e ReopenForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("reopenForumTopic", e)
}

type DeleteForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e DeleteForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e DeleteForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e DeleteForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("deleteForumTopic", e)
}

type UnpinAllForumTopicMessages[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e UnpinAllForumTopicMessages[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return objects.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e UnpinAllForumTopicMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllForumTopicMessages[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("unpinAllForumTopicMessages", e)
}

type EditGeneralForumTopic[T int | string] struct {
	ChatId T
	Name   string
}

func (e EditGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.Name) == "" {
		return objects.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (e EditGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditGeneralForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("editGeneralForumTopic", e)
}

type CloseGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e CloseGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e CloseGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseGeneralForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("closeGeneralForumTopic", e)
}

type ReopenGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e ReopenGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e ReopenGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenGeneralForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("reopenGeneralForumTopic", e)
}

type HideGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e HideGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e HideGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e HideGeneralForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("hideGeneralForumTopic", e)
}

type UnhideGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e UnhideGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e UnhideGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnhideGeneralForumTopic[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("unhideGeneralForumTopic", e)
}

type UnpinAllGeneralForumTopicMessages[T int | string] struct {
	ChatId T
}

func (e UnpinAllGeneralForumTopicMessages[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e UnpinAllGeneralForumTopicMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllGeneralForumTopicMessages[T]) Execute() (*bool, error) {
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

type GetUserChatBoosts[T int | string] struct {
	ChatId T
	UserId int
}

func (g GetUserChatBoosts[T]) Validate() error {
	if c, ok := any(g.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(g.ChatId).(int); ok {
		if c < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if g.UserId < 1 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (g GetUserChatBoosts[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(g)
}

func (g GetUserChatBoosts[T]) Execute() (*objects.UserChatBoosts, error) {
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

type SetChatMenuButton[T int | string] struct {
	ChatId     *T
	MenuButton objects.MenuButton
}

func (s SetChatMenuButton[T]) Validate() error {
	if s.ChatId != nil {
		if c, ok := any(*s.ChatId).(string); ok {
			if strings.TrimSpace(c) == "" {
				return objects.ErrInvalidParam("chat_id parameter can't be empty")
			}
		}
		if c, ok := any(*s.ChatId).(int); ok {
			if c < 1 {
				return objects.ErrInvalidParam("chat_id parameter can't be empty")
			}
		}
	}
	if s.MenuButton != nil {
		if err := s.MenuButton.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SetChatMenuButton[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatMenuButton[T]) Execute() (*bool, error) {
	return MakePostRequest[bool]("setChatMenuButton", s)
}

type GetChatMenuButton struct {
	ChatId *int
}

func (s GetChatMenuButton) Validate() error {
	if s.ChatId != nil {
		if *s.ChatId < 1 {
			return objects.ErrInvalidParam("chat_id parameter can't be empty")
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
