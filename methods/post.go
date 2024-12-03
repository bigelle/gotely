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
	iso6391 "github.com/emvi/iso-639-1"
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

func (s SendMessage[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(s)
}

func (s SendMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendMessage", s)
}

type ForwardMessage[T int | string] struct {
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

func (f ForwardMessage[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(f)
}

func (f ForwardMessage[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessage", f)
}

type ForwardMessages[T int | string] struct {
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

func (f ForwardMessages[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(f)
}

func (f ForwardMessages[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "forwardMessages", f)
}

type CopyMessage[T int | string] struct {
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

func (c CopyMessage[T]) ToRequestBody() ([]byte, error) {

	return json.Marshal(c)
}

func (c CopyMessage[T]) Execute() (*types.MessageId, error) {
	return internal.MakePostRequest[types.MessageId](telego.GetToken(), "copyMessage", c)
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

type SendPhoto[T int | string, B types.InputFile | string] struct {
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

func (s SendPhoto[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPhoto[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPhoto", s)
}

type SendAudio[T int | string, B types.InputFile | string] struct {
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

func (s SendAudio[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAudio[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendAudio", s)
}

type SendDocument[T int | string, B types.InputFile | string] struct {
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

func (s SendDocument[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDocument[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendDocument", s)
}

type SendVideo[T int | string, B types.InputFile | string] struct {
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

func (s SendVideo[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideo[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVideo", s)
}

type SendAnimation[T int | string, B types.InputFile | string] struct {
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

func (s SendAnimation[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendAnimation[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendAnimation", s)
}

type SendVoice[T int | string, B types.InputFile | string] struct {
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

func (s SendVoice[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVoice[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVoice", s)
}

type SendVideoNote[T int | string, B types.InputFile | string] struct {
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

func (s SendVideoNote[T, B]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVideoNote[T, B]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVideoNote", s)
}

type SendPaidMedia[T int | string] struct {
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

func (s SendPaidMedia[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPaidMedia[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPaidMedia", s)
}

type SendMediaGroup[T int | string] struct {
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

func (s SendMediaGroup[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendMediaGroup[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendMediaGroup", s)
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

func (s SendLocation[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendLocation[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendLocation", s)
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

func (s SendVenue[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendVenue[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendVenue", s)
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
func (s SendContact[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendContact[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendContact", s)
}

type SendPoll[T int | string] struct {
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

func (s SendPoll[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendPoll[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPoll", s)
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

func (s SendDice[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendDice[T]) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendPoll", s)
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

type SetMessageReaction[T int | string] struct {
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

func (s SetMessageReaction[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
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

func (s SetUserEmojiStatus) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
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

func (b BanChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatMember[T]) Execute() (*bool, error) {
	return internal.MakeGetRequest[bool](telego.GetToken(), "banChatMember", b)
}

type UnbanChatMember[T int | string] struct {
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

func (b UnbanChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatMember[T]) Execute() (*bool, error) {
	return internal.MakeGetRequest[bool](telego.GetToken(), "unbanChatMember", b)
}

type RestrictChatMember[T int | string] struct {
	ChatId                         T
	UserId                         int
	Permissions                    types.ChatPermissions
	UserIndependentChatPermissions *bool
	UntilDate                      *int
}

func (r RestrictChatMember[T]) Validate() error {
	if c, ok := any(r.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(r.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if r.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (r RestrictChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(r)
}

func (r RestrictChatMember[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "restrictChatMember", r)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (p PromoteChatMember[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PromoteChatMember[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "promoteChatMember", p)
}

type SetChatAdministratorCustomTitle[T int | string] struct {
	ChatId      T
	UserId      int
	CustomTitle string
}

func (s SetChatAdministratorCustomTitle[T]) Validate() error {
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
	if s.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if len(s.CustomTitle) > 16 {
		return errors.ErrInvalidParam("custom_title parameter must be not longer than 16 characters")
	}
	for _, r := range s.CustomTitle {
		if (r >= 0x1F600 && r <= 0x1F64F) || // Emoticons
			(r >= 0x1F300 && r <= 0x1F5FF) || // Miscellaneous Symbols and Pictographs
			(r >= 0x1F680 && r <= 0x1F6FF) || // Transport and Map Symbols
			(r >= 0x1F700 && r <= 0x1F77F) { // Alchemical Symbols
			return errors.ErrInvalidParam("invalid custom_title parameter: emojis are not allowed")
		}
	}
	return nil
}

func (s SetChatAdministratorCustomTitle[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatAdministratorCustomTitle[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatAdministratorCustomTitle", s)
}

type BanChatSenderChat[T int | string] struct {
	ChatId       T
	SenderChatId int
}

func (b BanChatSenderChat[T]) Validate() error {
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
	if b.SenderChatId < 1 {
		return errors.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b BanChatSenderChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b BanChatSenderChat[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "banChatSenderChat", b)
}

type UnbanChatSenderChat[T int | string] struct {
	ChatId       T
	SenderChatId int
}

func (b UnbanChatSenderChat[T]) Validate() error {
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
	if b.SenderChatId < 1 {
		return errors.ErrInvalidParam("sender_chat_id parameter can't be empty")
	}
	return nil
}

func (b UnbanChatSenderChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(b)
}

func (b UnbanChatSenderChat[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unbanChatSenderChat", b)
}

type SetChatPermissions[T int | string] struct {
	ChatId                         T
	Permissions                    types.ChatPermissions
	UserIndependentChatPermissions *bool
}

func (s SetChatPermissions[T]) Validate() error {
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
	return nil
}

func (s SetChatPermissions[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPermissions[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatPermissions", s)
}

type ExportChatInviteLink[T int | string] struct {
	ChatId T
}

func (e ExportChatInviteLink[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e ExportChatInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ExportChatInviteLink[T]) Execute() (*string, error) {
	return internal.MakePostRequest[string](telego.GetToken(), "exportChatInviteLink", e)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return errors.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return errors.ErrInvalidParam("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (c CreateInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateInviteLink[T]) Execute() (*types.ChatInviteLink, error) {
	return internal.MakePostRequest[types.ChatInviteLink](telego.GetToken(), "createInviteLink", c)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return errors.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	if c.MemberLimit != nil {
		if *c.MemberLimit < 1 || *c.MemberLimit > 99999 {
			return errors.ErrInvalidParam("member limit parameter must be between 1 and 99999")
		}
	}
	return nil
}

func (c EditInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditInviteLink[T]) Execute() (*types.ChatInviteLink, error) {
	return internal.MakePostRequest[types.ChatInviteLink](telego.GetToken(), "editInviteLink", c)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c.SubscriptionPeriod != 2592000 {
		return errors.ErrInvalidParam("subscription_period currently must always be 2592000 seconds (30 days)")
	}
	if c.SubscriptionPrice < 1 || c.SubscriptionPrice > 2500 {
		return errors.ErrInvalidParam("subscription_price must be between 1 and 2500")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return errors.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c CreateChatSubscriptionInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateChatSubscriptionInviteLink[T]) Execute() (*types.ChatInviteLink, error) {
	return internal.MakePostRequest[types.ChatInviteLink](telego.GetToken(), "createChatSubscriptionInviteLink", c)
}

type EditChatSubscriptionInviteLink[T int | string] struct {
	ChatId     T
	InviteLink string
	Name       *string
}

func (c EditChatSubscriptionInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(c.InviteLink) == "" {
		return errors.ErrInvalidParam("invite_link parameter can't be empty")
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return errors.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c EditChatSubscriptionInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c EditChatSubscriptionInviteLink[T]) Execute() (*types.ChatInviteLink, error) {
	return internal.MakePostRequest[types.ChatInviteLink](telego.GetToken(), "editChatSubscriptionInviteLink", c)
}

type RevokeInviteLink[T int | string] struct {
	ChatId     T
	InviteLink string
	Name       *string
}

func (c RevokeInviteLink[T]) Validate() error {
	if c, ok := any(c.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c.Name != nil {
		if len(*c.Name) > 32 {
			return errors.ErrInvalidParam("name parameter must not be longer than 32 characters")
		}
	}
	return nil
}

func (c RevokeInviteLink[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c RevokeInviteLink[T]) Execute() (*types.ChatInviteLink, error) {
	return internal.MakePostRequest[types.ChatInviteLink](telego.GetToken(), "revokeInviteLink", c)
}

type ApproveChatJoinRequest[T int | string] struct {
	ChatId T
	UserId int
}

func (s ApproveChatJoinRequest[T]) Validate() error {
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
	if s.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s ApproveChatJoinRequest[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s ApproveChatJoinRequest[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "approveChatJoinRequest", s)
}

type DeclineChatJoinRequest[T int | string] struct {
	ChatId T
	UserId int
}

func (s DeclineChatJoinRequest[T]) Validate() error {
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
	if s.UserId < 1 {
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	return nil
}

func (s DeclineChatJoinRequest[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeclineChatJoinRequest[T]) Execute() (*bool, error) {
	// NOTE: maybe there's a better way to get token?
	return internal.MakePostRequest[bool](telego.GetToken(), "declineChatJoinRequest", s)
}

type SetChatPhoto[T int | string] struct {
	ChatId T
	Photo  types.InputFile
}

func (s SetChatPhoto[T]) Validate() error {
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
	if err := s.Photo.Validate(); err != nil {
		return err
	}
	return nil
}

func (s SetChatPhoto[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatPhoto[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatPhoto", s)
}

type DeleteChatPhoto[T int | string] struct {
	ChatId T
}

func (d DeleteChatPhoto[T]) Validate() error {
	if c, ok := any(d.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(d.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (d DeleteChatPhoto[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteChatPhoto[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteChatPhoto", d)
}

type SetChatTitle[T int | string] struct {
	ChatId T
	Title  string
}

func (s SetChatTitle[T]) Validate() error {
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
	if len(s.Title) < 1 || len(s.Title) > 128 {
		return errors.ErrInvalidParam("title parameter must be between 1 and 128 characters long")
	}
	return nil
}

func (s SetChatTitle[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatTitle[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatTitle", s)
}

type SetChatDescription[T int | string] struct {
	ChatId      T
	Description string
}

func (s SetChatDescription[T]) Validate() error {
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
	if len(s.Description) > 255 {
		return errors.ErrInvalidParam("description parameter must not be longer than 255 characters")
	}
	return nil
}

func (s SetChatDescription[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetChatDescription[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatTitle", s)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.MessageId < 1 {
		return errors.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p PinChatMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p PinChatMessage[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "pinChatMessage", p)
}

type UnpinChatMessage[T int | string] struct {
	ChatId               T
	MessageId            int
	BusinessConnectionId *string
}

func (p UnpinChatMessage[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if p.MessageId < 1 {
		return errors.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (p UnpinChatMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinChatMessage[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unpinChatMessage", p)
}

type UnpinAllChatMessages[T int | string] struct {
	ChatId T
}

func (p UnpinAllChatMessages[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p UnpinAllChatMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p UnpinAllChatMessages[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unpinAllChatMessages", p)
}

type LeaveChat[T int | string] struct {
	ChatId T
}

func (p LeaveChat[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p LeaveChat[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p LeaveChat[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "leaveChat", p)
}

type SetChatStickerSet[T int | string] struct {
	ChatId         T
	StickerSetName string
}

func (p SetChatStickerSet[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(p.StickerSetName) == "" {
		return errors.ErrInvalidParam("sticker_set_name parameter can't be empty")
	}
	return nil
}

func (p SetChatStickerSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p SetChatStickerSet[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatStickerSet", p)
}

type DeleteChatStickerSet[T int | string] struct {
	ChatId T
}

func (p DeleteChatStickerSet[T]) Validate() error {
	if c, ok := any(p.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(p.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (p DeleteChatStickerSet[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(p)
}

func (p DeleteChatStickerSet[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteChatStickerSet", p)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(c.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if len(c.Name) < 1 || len(c.Name) > 128 {
		return errors.ErrInvalidParam("name parameter must be between 1 and 128 characters long")
	}
	if c.IconColor != nil {
		if _, ok := validIconColors[*c.IconColor]; !ok {
			return errors.ErrInvalidParam("icon_color must be one of 7322096 (0x6FB9F0), 16766590 (0xFFD67E), 13338331 (0xCB86DB), 9367192 (0x8EEE98), 16749490 (0xFF93B2), or 16478047 (0xFB6F5F)")
		}
	}
	return nil
}

func (c CreateForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(c)
}

func (c CreateForumTopic[T]) Execute() (*types.ForumTopic, error) {
	return internal.MakePostRequest[types.ForumTopic](telego.GetToken(), "createForumTopic", c)
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
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return errors.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	if e.Name != nil {
		if len(*e.Name) > 128 {
			return errors.ErrInvalidParam("name parameter must not be longer than 128 characters")
		}
	}
	return nil
}

func (e EditForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "editForumTopic", e)
}

type CloseForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e CloseForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return errors.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e CloseForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "closeForumTopic", e)
}

type ReopenForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e ReopenForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return errors.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e ReopenForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "reopenForumTopic", e)
}

type DeleteForumTopic[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e DeleteForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return errors.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e DeleteForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e DeleteForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteForumTopic", e)
}

type UnpinAllForumTopicMessages[T int | string] struct {
	ChatId          T
	MessageThreadId string
}

func (e UnpinAllForumTopicMessages[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.MessageThreadId) == "" {
		return errors.ErrInvalidParam("message_thread_id parameter can't be empty")
	}
	return nil
}

func (e UnpinAllForumTopicMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllForumTopicMessages[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unpinAllForumTopicMessages", e)
}

type EditGeneralForumTopic[T int | string] struct {
	ChatId T
	Name   string
}

func (e EditGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if strings.TrimSpace(e.Name) == "" {
		return errors.ErrInvalidParam("name parameter can't be empty")
	}
	return nil
}

func (e EditGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e EditGeneralForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "editGeneralForumTopic", e)
}

type CloseGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e CloseGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e CloseGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e CloseGeneralForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "closeGeneralForumTopic", e)
}

type ReopenGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e ReopenGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e ReopenGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e ReopenGeneralForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "reopenGeneralForumTopic", e)
}

type HideGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e HideGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e HideGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e HideGeneralForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "hideGeneralForumTopic", e)
}

type UnhideGeneralForumTopic[T int | string] struct {
	ChatId T
}

func (e UnhideGeneralForumTopic[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e UnhideGeneralForumTopic[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnhideGeneralForumTopic[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unhideGeneralForumTopic", e)
}

type UnpinAllGeneralForumTopicMessages[T int | string] struct {
	ChatId T
}

func (e UnpinAllGeneralForumTopicMessages[T]) Validate() error {
	if c, ok := any(e.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(e.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	return nil
}

func (e UnpinAllGeneralForumTopicMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(e)
}

func (e UnpinAllGeneralForumTopicMessages[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "unpinAllGeneralForumTopicMessages", e)
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
		return errors.ErrInvalidParam("callback_query_id parameter can't be empty")
	}
	if a.Text != nil {
		if len(*a.Text) > 200 {
			return errors.ErrInvalidParam("text parameter must not be longer than 200 characters ")
		}
	}
	return nil
}

type SetMyCommands struct {
	Commands     []types.BotCommand
	Scope        *types.BotCommandScope
	LanguageCode *string
}

func (s SetMyCommands) Validate() error {
	for _, command := range s.Commands {
		if err := command.Validate(); err != nil {
			return err
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		// FIXME: maybe should replace it with my own list of lang codes
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyCommands) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setMyCommands", s)
}

type DeleteMyCommands struct {
	Scope        *types.BotCommandScope
	LanguageCode *string
}

func (s DeleteMyCommands) Validate() error {
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s DeleteMyCommands) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s DeleteMyCommands) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteMyCommands", s)
}

type SetMyName struct {
	Name         *string
	LanguageCode *string
}

func (s SetMyName) Validate() error {
	if s.Name != nil {
		if len(*s.Name) > 64 {
			return errors.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyName) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyName) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setMyName", s)
}

type SetMyDescription struct {
	Description  *string
	LanguageCode *string
}

func (s SetMyDescription) Validate() error {
	if s.Description != nil {
		if len(*s.Description) > 64 {
			return errors.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyDescription) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setMyDescription", s)
}

type SetMyShortDescription struct {
	ShortDescription *string
	LanguageCode     *string
}

func (s SetMyShortDescription) Validate() error {
	if s.ShortDescription != nil {
		if len(*s.ShortDescription) > 64 {
			return errors.ErrInvalidParam("name parameter must not be longer than 64 characters")
		}
	}
	if s.LanguageCode != nil && *s.LanguageCode != "" {
		if !iso6391.ValidCode(*s.LanguageCode) {
			return errors.ErrInvalidParam(fmt.Sprintf("invalid language code: %s", *s.LanguageCode))
		}
	}
	return nil
}

func (s SetMyShortDescription) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetMyShortDescription) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "setMyShortDescription", s)
}

type SetChatMenuButton[T int | string] struct {
	ChatId     *T
	MenuButton *types.MenuButton
}

func (s SetChatMenuButton[T]) Validate() error {
	if s.ChatId != nil {
		if c, ok := any(*s.ChatId).(string); ok {
			if strings.TrimSpace(c) == "" {
				return errors.ErrInvalidParam("chat_id parameter can't be empty")
			}
		}
		if c, ok := any(*s.ChatId).(int); ok {
			if c < 1 {
				return errors.ErrInvalidParam("chat_id parameter can't be empty")
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
	return internal.MakePostRequest[bool](telego.GetToken(), "setChatMenuButton", s)
}

type SetMyDefaultAdministratorRights struct {
	Rights      *types.ChatAdministratorRights
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
	return internal.MakePostRequest[bool](telego.GetToken(), "setMyDefaultAdministratorRights", s)
}

type StopPoll[T int | string] struct {
	ChatId               T
	MessageId            int
	BusinessConnectionId *string
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (s StopPoll[T]) Validate() error {
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
	if s.ReplyMarkup != nil {
		if err := s.ReplyMarkup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (s StopPoll[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s StopPoll[T]) Execute() (*types.Poll, error) {
	return internal.MakePostRequest[types.Poll](telego.GetToken(), "stopPoll", s)
}

type DeleteMessage[T int | string] struct {
	ChatId    T
	MessageId int
}

func (d DeleteMessage[T]) Validate() error {
	if c, ok := any(d.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(d.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if d.MessageId < 1 {
		return errors.ErrInvalidParam("message_id parameter can't be empty")
	}
	return nil
}

func (d DeleteMessage[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessage[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteMessage", d)
}

type DeleteMessages[T int | string] struct {
	ChatId     T
	MessageIds []int
}

func (d DeleteMessages[T]) Validate() error {
	if c, ok := any(d.ChatId).(string); ok {
		if strings.TrimSpace(c) == "" {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if c, ok := any(d.ChatId).(int); ok {
		if c < 1 {
			return errors.ErrInvalidParam("chat_id parameter can't be empty")
		}
	}
	if len(d.MessageIds) < 1 || len(d.MessageIds) > 100 {
		return errors.ErrInvalidParam("message_ids parameter must be between 1 and 100")
	}

	return nil
}

func (d DeleteMessages[T]) ToRequestBody() ([]byte, error) {
	return json.Marshal(d)
}

func (d DeleteMessages[T]) Execute() (*bool, error) {
	return internal.MakePostRequest[bool](telego.GetToken(), "deleteMessages", d)
}
