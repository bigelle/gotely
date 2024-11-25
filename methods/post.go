package methods

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/assertions"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type SendMessage[T int | string] struct {
	ChatId               T                         `json:"chat_id"`
	Text                 string                    `json:"text"`
	BusinessConnectionId *string                   `json:"business_connection_id,omitempty"`
	MessageThreadId      *int                      `json:"message_thread_id,omitempty"`
	ParseMode            *string                   `json:"parse_mode,omitempty"`
	Entities             *[]types.MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions   *types.LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  *bool                     `json:"disable_notification,omitempty"`
	ProtectContent       *bool                     `json:"protect_content,omitempty"`
	MessageEffectId      *string                   `json:"message_effect_id,omitempty"`
	ReplyParameters      *types.ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          *types.ReplyMarkup        `json:"reply_markup,omitempty"`
}

func (s SendMessage[T]) Validate() error {
	if err := assertions.ParamNotEmpty(s.Text, "Text"); err != nil {
		return err
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c == 0 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(string); ok {
		if err := assertions.ParamNotEmpty(c, "ChatId"); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMessage[T]) MarshalJSON() ([]byte, error) {
	type alias SendMessage[T]
	return json.Marshal(alias(s))
}

func (s SendMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendMessage", s)
}

type ForwardMessage[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageId           int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
}

func (f ForwardMessage[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if f.MessageId < 1 {
		return assertions.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (f ForwardMessage[T]) MarshalJSON() ([]byte, error) {
	type alias ForwardMessage[T]
	return json.Marshal(alias(f))
}

func (f ForwardMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessage", f)
}

type ForwardMessages[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageIds          []int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
}

func (f ForwardMessages[T]) Validate() error {
	if c, ok := any(f.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(f.MessageIds) < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (f ForwardMessages[T]) MarshalJSON() ([]byte, error) {
	type alias ForwardMessages[T]
	return json.Marshal(alias(f))
}

func (f ForwardMessages[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessages", f)
}

type CopyMessage[T string | int] struct {
	ChatId                T
	FromChatId            T
	MessageId             int
	MessageThreadId       *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	AllowPaidBroadcast    *bool
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
	DisableNotification   *bool
	ProtectContent        *bool
}

func (c CopyMessage[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c.MessageId < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessage[T]) MarshalJSON() ([]byte, error) {
	type alias CopyMessage[T]
	return json.Marshal(alias(c))
}

func (c CopyMessage[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessage", c)
}

type CopyMessages[T string | int] struct {
	ChatId              T
	FromChatId          T
	MessageIds          []int
	MessageThreadId     *int
	DisableNotification *bool
	ProtectContent      *bool
	RemoveCaption       *bool
}

func (c CopyMessages[T]) MarshalJSON() ([]byte, error) {
	type alias CopyMessages[T]
	return json.Marshal(alias(c))
}

func (c CopyMessages[T]) Validate() error {
	if i, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return assertions.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(c.MessageIds) < 1 {
		return assertions.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessages[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessages", c)
}

type SendPhoto[T string | int, B types.InputFile | string] struct {
	ChatId                T
	Photo                 B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
}

func (s SendPhoto[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Photo).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Photo).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendPhoto[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendPhoto[T, B]
	return json.Marshal(alias(s))
}

func (s SendPhoto[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPhoto", s)
}

type SendAudio[T string | int, B types.InputFile | string] struct {
	ChatId               T
	Audio                B
	BusinessConnectionId *string
	MessageThreadId      *int
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]types.MessageEntity
	Duration             *int
	Performer            *string
	Title                *string
	Thumbnail            *B
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendAudio[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Audio).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid audio parameter: %w", err)
		}
	}
	if p, ok := any(s.Audio).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("audio parameter can't be empty")
		}
	}
	return nil
}

func (s SendAudio[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendAudio[T, B]
	return json.Marshal(alias(s))
}

func (s SendAudio[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendAudio", s)
}

type SendDocument[T string | int, B types.InputFile | string] struct {
	ChatId                      T
	Document                    B
	BusinessConnectionId        *string
	MessageThreadId             *int
	Caption                     *string
	ParseMode                   *string
	CaptionEntities             *[]types.MessageEntity
	DisableContentTypeDetection *bool
	Thumbnail                   *B
	DisableNotification         *bool
	ProtectContent              *bool
	AllowPaidBroadcast          *bool
	MessageEffectId             *string
	ReplyParameters             *types.ReplyParameters
	ReplyMarkup                 *types.ReplyMarkup
}

func (s SendDocument[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Document).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid document parameter: %w", err)
		}
	}
	if p, ok := any(s.Document).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("document parameter can't be empty")
		}
	}
	return nil
}

func (s SendDocument[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendDocument[T, B]
	return json.Marshal(alias(s))
}

func (s SendDocument[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendDocument", s)
}

type SendVideo[T string | int, B types.InputFile | string] struct {
	ChatId                T
	Video                 B
	BusinessConnectionId  *string
	MessageThreadId       *int
	Duration              *int
	Width                 *int
	Height                *int
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	SupportsStreaming     *bool
	Thumbnail             *B
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
}

func (s SendVideo[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Video).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video parameter: %w", err)
		}
	}
	if p, ok := any(s.Video).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("video parameter can't be empty")
		}
	}
	return nil
}

func (s SendVideo[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendVideo[T, B]
	return json.Marshal(alias(s))
}

func (s SendVideo[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVideo", s)
}

type SendAnimation[T string | int, B types.InputFile | string] struct {
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
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	HasSpoiler            *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	MessageEffectId       *string
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
}

func (s SendAnimation[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Animation).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Animation).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("photo parameter can't be empty")
		}
	}
	return nil
}

func (s SendAnimation[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendAnimation[T, B]
	return json.Marshal(alias(s))
}

func (s SendAnimation[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendAnimation", s)
}

type SendVoice[T string | int, B types.InputFile | string] struct {
	ChatId               T
	Voice                B
	BusinessConnectionId *string
	MessageThreadId      *int
	Duration             *int
	Thumbnail            *B
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]types.MessageEntity
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendVoice[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Voice).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid voice parameter: %w", err)
		}
	}
	if p, ok := any(s.Voice).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("voice parameter can't be empty")
		}
	}
	return nil
}

func (s SendVoice[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendVoice[T, B]
	return json.Marshal(alias(s))
}

func (s SendVoice[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVoice", s)
}

type SendVideoNote[T string | int, B types.InputFile | string] struct {
	ChatId               T
	VideoNote            B
	BusinessConnectionId *string
	MessageThreadId      *int
	Duration             *int
	Length               *int
	Caption              *string
	ParseMode            *string
	CaptionEntities      *[]types.MessageEntity
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendVideoNote[T, B]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.VideoNote).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video_note parameter: %w", err)
		}
	}
	if p, ok := any(s.VideoNote).(string); ok {
		if strings.TrimSpace(p) == "" {
			return assertions.ErrInvalidParam("video_note parameter can't be empty")
		}
	}
	// TODO: validate non-nill thumbnails
	return nil
}

func (s SendVideoNote[T, B]) MarshalJSON() ([]byte, error) {
	type alias SendVideoNote[T, B]
	return json.Marshal(alias(s))
}

func (s SendVideoNote[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVideoNote", s)
}

type SendPaidMedia[T string | int] struct {
	ChatId                T
	Media                 []types.InputPaidMedia
	StarCount             int
	BusinessConnectionId  *string
	Payload               *string
	Caption               *string
	ParseMode             *string
	CaptionEntities       *[]types.MessageEntity
	ShowCaptionAboveMedia *bool
	DisableNotification   *bool
	ProtectContent        *bool
	AllowPaidBroadcast    *bool
	ReplyParameters       *types.ReplyParameters
	ReplyMarkup           *types.ReplyMarkup
}

func (s SendPaidMedia[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.StarCount < 1 || s.StarCount > 2500 {
		return assertions.ErrInvalidParam("star_count parameter must be between 1 and 2500")
	}
	if len(s.Media) < 1 {
		return assertions.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return assertions.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendPaidMedia[T]) MarshalJSON() ([]byte, error) {
	type alias SendPaidMedia[T]
	return json.Marshal(alias(s))
}

func (s SendPaidMedia[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPaidMedia", s)
}

type SendMediaGroup[T string | int] struct {
	ChatId               T
	Media                []types.InputMedia
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
}

func (s SendMediaGroup[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if len(s.Media) < 1 {
		return assertions.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return assertions.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
	}
	for _, m := range s.Media {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMediaGroup[T]) MarshalJSON() ([]byte, error) {
	type alias SendMediaGroup[T]
	return json.Marshal(alias(s))
}

func (s SendMediaGroup[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendMediaGroup", s)
}

type SendLocation[T string | int] struct {
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
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendLocation[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.Latitude == nil {
		return assertions.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return assertions.ErrInvalidParam("longtitude parameter can't be empty")
	}
	// TODO: validate replyparameters everywhere
	return nil
}

func (s SendLocation[T]) MarshalJSON() ([]byte, error) {
	type alias SendLocation[T]
	return json.Marshal(alias(s))
}

func (s SendLocation[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendLocation", s)
}

type SendVenue[T string | int] struct {
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
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendVenue[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.Latitude == nil {
		return assertions.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return assertions.ErrInvalidParam("longtitude parameter can't be empty")
	}
	return nil
}

func (s SendVenue[T]) MarshalJSON() ([]byte, error) {
	type alias SendVenue[T]
	return json.Marshal(alias(s))
}

func (s SendVenue[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVenue", s)
}

type SendContact[T string | int] struct {
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
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendContact[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.PhoneNumber) == "" {
		return assertions.ErrInvalidParam("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return assertions.ErrInvalidParam("first_name parameter can't be empty")
	}
	return nil
}

// NOTE: do i need it?
func (s SendContact[T]) MarshalJSON() ([]byte, error) {
	type alias SendContact[T]
	return json.Marshal(alias(s))
}

func (s SendContact[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendContact", s)
}

type SendPoll[T string | int] struct {
	ChatId               T
	Question             string
	Options              []types.InputPollOption
	QuestionParseMode    *string
	QuestionEntities     *[]types.MessageEntity
	IsAnonymous          *bool
	Type                 *string
	AllowMultipleAnswers *bool
	CorrectOptionId      *int
	Explanation          *string
	ExplanationParseMode *string
	ExplanationEntities  *[]types.MessageEntity
	OpenPeriod           *int
	CloseDate            *int
	IsClosed             *bool
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendPoll[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Question) == "" {
		return assertions.ErrInvalidParam("question parameter can't be empty")
	}
	if len(s.Options) < 2 || len(s.Options) > 10 {
		return assertions.ErrInvalidParam("options parameter must be between 2 and 10")
	}
	return nil
}

func (s SendPoll[T]) MarshalJSON() ([]byte, error) {
	type alias SendPoll[T]
	return json.Marshal(alias(s))
}

func (s SendPoll[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPoll", s)
}

type SendDice[T string | int] struct {
	ChatId               T
	Emoji                string
	BusinessConnectionId *string
	MessageThreadId      *string
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyMarkup
}

func (s SendDice[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Emoji) == "" {
		return assertions.ErrInvalidParam("emoji parameter can't be empty")
	}
	return nil
}

func (s SendDice[T]) MarshalJSON() ([]byte, error) {
	type alias SendDice[T]
	return json.Marshal(alias(s))
}

func (s SendDice[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPoll", s)
}

type SendChatAction[T string | int] struct {
	ChatId               T
	Action               string
	BusinessConnectionId *string
	MessageThreadId      *string
}

func (s SendChatAction[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return assertions.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Action) == "" {
		return assertions.ErrInvalidParam("action parameter can't be empty")
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
		return assertions.ErrInvalidParam(fmt.Sprintf("action must be %s or upload_video_note", strings.Join(allowed[:len(allowed)-1], ", ")))
	}
	return nil
}
