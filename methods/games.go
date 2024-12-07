package methods

import (
	"encoding/json"
	"strings"

	telego "github.com/bigelle/tele.go"
	"github.com/bigelle/tele.go/errors"
	"github.com/bigelle/tele.go/internal"
	"github.com/bigelle/tele.go/types"
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
	ReplyParameters      *types.ReplyParameters
	ReplyMarkup          *types.InlineKeyboardMarkup
}

func (s SendGame) Validate() error {
	if s.ChatId == 0 {
		return errors.ErrInvalidParam("chat_id parameter can't be empty")
	}
	if strings.TrimSpace(s.GameShortName) == "" {
		return errors.ErrInvalidParam("game_short_name parameter can't be empty")
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

func (s SendGame) Execute() (*types.Message, error) {
	return internal.MakePostRequest[types.Message](telego.GetToken(), "sendGame", s)
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
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.Score < 0 {
		return errors.ErrInvalidParam("score parameter must be non-negative")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
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
		b, err := internal.MakePostRequest[bool](telego.GetToken(), "setGameScore", s)
		return MessageOrBool{
			Message: nil,
			Bool:    b,
		}, err
	} else {
		// expecting a Message
		msg, err := internal.MakePostRequest[types.Message](telego.GetToken(), "setGameScore", s)
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
		return errors.ErrInvalidParam("user_id parameter can't be empty")
	}
	if s.InlineMessageId == nil {
		if s.ChatId == nil {
			return errors.ErrInvalidParam("chat_id parameter can't be empty if inline_message_id is not specified")
		}
		if s.MessageId == nil {
			return errors.ErrInvalidParam("message_id parameter can't be empty if inline_message_id is not specified")
		}
	}
	if s.ChatId == nil && s.MessageId == nil {
		if s.InlineMessageId == nil {
			return errors.ErrInvalidParam("inline_message_id can't be empty if chat_id and message_id are not specified")
		}
	}
	return nil
}

func (s GetGameHighScores) ToRequestBody() ([]byte, error) {
	return json.Marshal(s)
}

func (g GetGameHighScores) Execute() (*[]types.GameHighScore, error) {
	return internal.MakePostRequest[[]types.GameHighScore](telego.GetToken(), "getGameHighScores", g)
}
