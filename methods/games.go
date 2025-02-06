package methods

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bigelle/gotely/objects"
)

type SendGame struct {
	ChatId               int
	GameShortName        string
	BusinessConnectionId *string
	MessageThreadId      *int
	DisableNotification  *bool
	ProtectContent       *bool
	AllowPaidBroadcast   *bool
	MessageEffectId      *string
	ReplyParameters      *objects.ReplyParameters
	ReplyMarkup          *objects.InlineKeyboardMarkup
	client               *http.Client
	baseUrl              string
}

func (s *SendGame) WithClient(c *http.Client) *SendGame {
	s.client = c
	return s
}

func (s SendGame) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SendGame) WithApiBaseUrl(u string) *SendGame {
	s.baseUrl = u
	return s
}

func (s SendGame) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s SendGame) Validate() error {
	if s.ChatId == 0 {
		return objects.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.GameShortName) == "" {
		return objects.ErrInvalidParam("game_short_name parameter can't be empty")
	}
	if s.ReplyMarkup != nil {
		if err := s.ReplyMarkup.Validate(); err != nil {
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

func (s SendGame) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SendGame) Execute(token string) (*objects.Message, error) {
	return SendTelegramPostRequest[objects.Message](token, "sendGame", s)
}

type SetGameHighScore struct {
	UserId             int
	Score              int
	Force              *bool
	DisableEditMessage *bool
	ChatId             *int
	MessageId          *int
	InlineMessageId    *string
	client             *http.Client
	baseUrl            string
}

func (s *SetGameHighScore) WithClient(c *http.Client) *SetGameHighScore {
	s.client = c
	return s
}

func (s SetGameHighScore) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *SetGameHighScore) WithApiBaseUrl(u string) *SetGameHighScore {
	s.baseUrl = u
	return s
}

func (s SetGameHighScore) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s SetGameHighScore) Validate() error {
	if s.UserId == 0 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.Score < 0 {
		return objects.ErrInvalidParam("score parameter must be non-negative")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
		}
	}
	return nil
}

func (s SetGameHighScore) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (s SetGameHighScore) Execute(token string) (MessageOrBool, error) {
	if s.InlineMessageId != nil {
		// expecting a boolean
		b, err := SendTelegramPostRequest[bool](token, "setGameScore", s)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := SendTelegramPostRequest[objects.Message](token, "setGameScore", s)
		return MessageOrBool{
			Message: msg,
			Bool:    nil,
		}, err
	}
}

type GetGameHighScores struct {
	UserId          int
	ChatId          *int
	MessageId       *int
	InlineMessageId *string
	client          *http.Client
	baseUrl         string
}

func (s *GetGameHighScores) WithClient(c *http.Client) *GetGameHighScores {
	s.client = c
	return s
}

func (s GetGameHighScores) Client() *http.Client {
	if s.client == nil {
		return &http.Client{}
	}
	return s.client
}

func (s *GetGameHighScores) WithApiBaseUrl(u string) *GetGameHighScores {
	s.baseUrl = u
	return s
}

func (s GetGameHighScores) ApiBaseUrl() string {
	if s.baseUrl == "" {
		return "https://api.telegram.org/bot%s/%s"
	}
	return s.baseUrl
}

func (s GetGameHighScores) Validate() error {
	if s.UserId == 0 {
		return objects.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return objects.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return objects.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return objects.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
		}
	}
	return nil
}

func (s GetGameHighScores) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (g GetGameHighScores) Execute(token string) (*[]objects.GameHighScore, error) {
	return SendTelegramPostRequest[[]objects.GameHighScore](token, "getGameHighScores", g)
}
