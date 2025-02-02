package methods

import (
	"encoding/json"
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

func (s SendGame) Execute() (*objects.Message, error) {
	return MakePostRequest[objects.Message]("sendGame", s)
}

type SetGameHighScore struct {
	UserId             int
	Score              int
	Force              *bool
	DisableEditMessage *bool
	ChatId             *int
	MessageId          *int
	InlineMessageId    *string
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

func (s SetGameHighScore) Execute() (MessageOrBool, error) {
	if s.InlineMessageId != nil {
		// expecting a boolean
		b, err := MakePostRequest[bool]("setGameScore", s)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := MakePostRequest[objects.Message]("setGameScore", s)
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

func (g GetGameHighScores) Execute() (*[]objects.GameHighScore, error) {
	return MakePostRequest[[]objects.GameHighScore]("getGameHighScores", g)
}
