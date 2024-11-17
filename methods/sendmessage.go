package methods

import (
	"encoding/json"
	"fmt"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/internal/assertions"
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
	ReplyMarkup          *types.ReplyKeyboard      `json:"reply_markup,omitempty"`
}

func (s SendMessage[T]) New(chatId T, text string) *SendMessage[T] {
	return &SendMessage[T]{
		ChatId: chatId,
		Text:   text,
	}
}

func (s *SendMessage[T]) SetBusinessConnectionId(id string) *SendMessage[T] {
	s.BusinessConnectionId = &id
	return s
}

func (s *SendMessage[T]) SetMessageThreadId(id int) *SendMessage[T] {
	s.MessageThreadId = &id
	return s
}

func (s *SendMessage[T]) SetParseMode(m string) *SendMessage[T] {
	s.ParseMode = &m
	return s
}

func (s *SendMessage[T]) SetEntities(en []types.MessageEntity) *SendMessage[T] {
	s.Entities = &en
	return s
}

func (s *SendMessage[T]) SetLinkPreviewOptions(opt types.LinkPreviewOptions) *SendMessage[T] {
	s.LinkPreviewOptions = &opt
	return s
}

func (s *SendMessage[T]) SetDisableNotifications(b bool) *SendMessage[T] {
	s.DisableNotification = &b
	return s
}

func (s *SendMessage[T]) SetProtectContent(b bool) *SendMessage[T] {
	s.ProtectContent = &b
	return s
}

func (s *SendMessage[T]) SetMessageEffectId(id string) *SendMessage[T] {
	s.MessageEffectId = &id
	return s
}

func (s *SendMessage[T]) SetReplyParameters(rp types.ReplyParameters) *SendMessage[T] {
	s.ReplyParameters = &rp
	return s
}

func (s *SendMessage[T]) SetReplyMarkup(rm types.ReplyKeyboard) *SendMessage[T] {
	s.ReplyMarkup = &rm
	return s
}

func (s SendMessage[T]) Validate() error {
	if err := assertions.ParamNotEmpty(s.Text, "Text"); err != nil {
		return err
	}
	if c, ok := any(s.ChatId).(int); ok {
		if c == 0 {
			return assertions.ErrorEmptyParam{Param: "Id"}
		}
	}
	if c, ok := any(s.ChatId).(string); ok {
		if err := assertions.ParamNotEmpty(c, "ChatId"); err != nil {
			return err
		}
	}
	return nil
}

func (s SendMessage[T]) Execute() (*types.Message, error) {
	//TODO: proper error handling (logging, custom error types)

	// validating before preparing request payload
	if err := s.Validate(); err != nil {
		return nil, err
	}

	// request payload
	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	// sending request and getting response bytes
	// FIXME: should have better way to get token because
	// its not always doing requests using longpolling bot
	// maybe should create a chan HERE, pass it to MakePostRequest and
	// get a response from this channel
	b, err := internal.MakePostRequest(telego.GetToken(), "sendMessage", data)
	if err != nil {
		return nil, err
	}

	// parsing response
	var apiresp types.ApiResponse[types.Message]
	if err := json.Unmarshal(b, &apiresp); err != nil {
		return nil, err
	}
	if !apiresp.Ok {
		return nil, fmt.Errorf("%d: %s", apiresp.ErrorCode, *apiresp.Description)
	}
	return &apiresp.Result, nil
}
