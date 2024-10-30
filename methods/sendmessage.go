package methods

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/internal/assertions"
	"github.com/bigelle/tele.go/types"
)

type SendMessage[T int | string] struct {
	ChatId               T
	Text                 string
	BusinessConnectionId *string
	MessageThreadId      *int
	ParseMode            *string
	Entities             *[]types.MessageEntity
	LinkPreviewOptions   *types.LinkPreviewOptions
	DisableNotification  *bool
	ProtectContent       *bool
	MessageEffectId      *string
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.ReplyKeyboard
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
	if err := s.Validate(); err != nil {
		return nil, err
	}

	data, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	bot := telego.GetBot()
	if bot.Token == "" {
		return nil, errors.New("API token can't be empty")
	}

	reqUrl := fmt.Sprintf("%s%s%s", bot.ApiUrl, bot.Token, "/sendMessage")
	resp, err := http.Post(reqUrl, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiresp types.ApiResponse[types.Message]
	if err := json.Unmarshal(b, &apiresp); err != nil {
		return nil, err
	}
	if !apiresp.Ok{
		return nil, fmt.Errorf("failed request: %s", *apiresp.Description)
	}
	return &apiresp.Result, nil
}
