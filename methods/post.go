package methods

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
)

type ChatId interface {
	int | string
}
type SendMessage[T ChatId] struct {
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
	if strings.TrimSpace(s.Text) == "" {
		return errors.ErrInvalidParam("text parameter can't be empty")
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c == 0 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
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

type ForwardMessage[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if f.MessageId < 1 {
		return errors.ErrInvalidParam("message_id parameter can't be empty")
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

type ForwardMessages[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c, ok := any(f.FromChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(f.MessageIds) < 1 {
		return errors.ErrInvalidParam("message_ids parameter can't be empty")
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

type CopyMessage[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if c.MessageId < 1 {
		return errors.ErrInvalidParam("message_ids parameter can't be empty")
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

type CopyMessages[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.ChatId).(int); ok {
		if i < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(string); ok {
		if strings.TrimSpace(i) == "" {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if i, ok := any(c.FromChatId).(int); ok {
		if i < 1 {
			return errors.ErrInvalidParam("from_chat_id parameter can't be empty")
		}
	}
	if len(c.MessageIds) < 1 {
		return errors.ErrInvalidParam("message_ids parameter can't be empty")
	}
	return nil
}

func (c CopyMessages[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessages", c)
}

type SendPhoto[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Photo).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Photo).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("photo parameter can't be empty")
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

type SendAudio[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Audio).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid audio parameter: %w", err)
		}
	}
	if p, ok := any(s.Audio).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("audio parameter can't be empty")
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

type SendDocument[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Document).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid document parameter: %w", err)
		}
	}
	if p, ok := any(s.Document).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("document parameter can't be empty")
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

type SendVideo[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Video).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video parameter: %w", err)
		}
	}
	if p, ok := any(s.Video).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("video parameter can't be empty")
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

type SendAnimation[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Animation).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid photo parameter: %w", err)
		}
	}
	if p, ok := any(s.Animation).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("photo parameter can't be empty")
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

type SendVoice[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.Voice).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid voice parameter: %w", err)
		}
	}
	if p, ok := any(s.Voice).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("voice parameter can't be empty")
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

type SendVideoNote[T ChatId, B types.InputFile | string] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p, ok := any(s.VideoNote).(types.InputFile); ok {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("invalid video_note parameter: %w", err)
		}
	}
	if p, ok := any(s.VideoNote).(string); ok {
		if strings.TrimSpace(p) == "" {
			return errors.ErrInvalidParam("video_note parameter can't be empty")
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

type SendPaidMedia[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.StarCount < 1 || s.StarCount > 2500 {
		return errors.ErrInvalidParam("star_count parameter must be between 1 and 2500")
	}
	if len(s.Media) < 1 {
		return errors.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return errors.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
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

type SendMediaGroup[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if len(s.Media) < 1 {
		return errors.ErrInvalidParam("media parameter can't be empty")
	}
	if len(s.Media) > 10 {
		return errors.ErrInvalidParam("can't accept more than 10 InputPaidMedia in media parameter")
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

type SendLocation[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.Latitude == nil {
		return errors.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return errors.ErrInvalidParam("longtitude parameter can't be empty")
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

type SendVenue[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.Latitude == nil {
		return errors.ErrInvalidParam("latitude parameter can't be empty")
	}
	if s.Longtitude == nil {
		return errors.ErrInvalidParam("longtitude parameter can't be empty")
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

type SendContact[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.PhoneNumber) == "" {
		return errors.ErrInvalidParam("phone_number parameter can't be empty")
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return errors.ErrInvalidParam("first_name parameter can't be empty")
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

type SendPoll[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Question) == "" {
		return errors.ErrInvalidParam("question parameter can't be empty")
	}
	if len(s.Options) < 2 || len(s.Options) > 10 {
		return errors.ErrInvalidParam("options parameter must be between 2 and 10")
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

type SendDice[T ChatId] struct {
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Emoji) == "" {
		return errors.ErrInvalidParam("emoji parameter can't be empty")
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

type SendChatAction[T ChatId] struct {
	ChatId               T
	Action               string
	BusinessConnectionId *string
	MessageThreadId      *string
}

func (s SendChatAction[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(s.Action) == "" {
		return errors.ErrInvalidParam("action parameter can't be empty")
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
		return errors.ErrInvalidParam(fmt.Sprintf("action must be %s or upload_video_note", strings.Join(allowed[:len(allowed)-1], ", ")))
	}
	return nil
}

type SetMessageReaction[T ChatId] struct {
	ChatId    T
	MessageId int
	Reaction  *[]types.ReactionType
	IsBig     *bool
}

func (s SetMessageReaction[T]) Validate() error {
	if c, ok := any(s.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if s.MessageId < 1 {
		return errors.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (s SetMessageReaction[T]) MarshalJSON() ([]byte, error) {
	type alias SetMessageReaction[T]
	return json.Marshal(alias(s))
}

func (s SetMessageReaction[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setMessageReaction", s)
}

type SetUserEmojiStatus struct {
	UserId                    int
	EmojiStatusCustomEmojiId  *string
	EmojiStatusExpirationDate *int
}

func (s SetUserEmojiStatus) Validate() error {
	if s.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s SetUserEmojiStatus) MarshalJSON() ([]byte, error) {
	type alias SetUserEmojiStatus
	return json.Marshal(alias(s))
}

type BanChatMember[T ChatId] struct {
	ChatId         T
	UserId         int
	UntilDate      *int
	RevokeMessages *bool
}

func (b BanChatMember[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b BanChatMember[T]) MarshalJSON() ([]byte, error) {
	type alias BanChatMember[T]
	return json.Marshal(alias(b))
}

func (b BanChatMember[T]) Execute() (*bool, error) {
	return internal.MakeGetRequest[bool](telego.GetToken(), "banChatMember", b)
}

type UnbanChatMember[T ChatId] struct {
	ChatId       T
	UserId       int
	OnlyIfBanned *bool
}

func (b UnbanChatMember[T]) Validate() error {
	if c, ok := any(b.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(b.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if b.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatMember[T]) MarshalJSON() ([]byte, error) {
	type alias UnbanChatMember[T]
	return json.Marshal(alias(b))
}

func (b UnbanChatMember[T]) Execute() (*bool, error) {
	return internal.MakeGetRequest[bool](telego.GetToken(), "unbanChatMember", b)
}
